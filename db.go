package hydrocarbon

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var (
	ErrInvalidSession = errors.New("invalid session token")
)

// A DB is responsible for all interactions with postgres
type DB struct {
	sql *sql.DB
}

// NewDB returns a new database
func NewDB(dsn string) (*DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = runMigrations(db)
	if err != nil {
		return nil, err
	}

	return &DB{
		sql: db,
	}, nil
}

// CreateOrGetUser creates a new user and returns the users ID
func (db *DB) CreateOrGetUser(ctx context.Context, email string) (string, error) {
	row := db.sql.QueryRowContext(ctx, `INSERT INTO users 
										(email) 
										VALUES ($1)
										ON CONFLICT (email)
										DO UPDATE SET email = EXCLUDED.email
										RETURNING id;`, email)

	var userID string
	err := row.Scan(&userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}

// CreateLoginToken creates a new one-time-use login token
func (db *DB) CreateLoginToken(ctx context.Context, userID, userAgent, ip string) (string, error) {
	row := db.sql.QueryRowContext(ctx, `INSERT INTO login_tokens
										(user_id, user_agent, ip)
										VALUES ($1, $2, $3::cidr)
										RETURNING token;`, userID, userAgent, ip)

	var token string
	err := row.Scan(&token)
	if err != nil {
		return "", err
	}

	return token, nil
}

// ActivateLoginToken activates the given LoginToken and returns the user
// the token was for
func (db *DB) ActivateLoginToken(ctx context.Context, token string) (string, error) {
	row := db.sql.QueryRowContext(ctx, `UPDATE login_tokens
										SET (used) = (true)
										WHERE token = $1
										AND expires_at > now()
										AND used = false
										RETURNING user_id;`, token)

	var userID string
	err := row.Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("token invalid")
		}
		return "", err
	}

	return userID, nil
}

// CreateSession creates a new session for the user ID and returns the
// session key
func (db *DB) CreateSession(ctx context.Context, userID, userAgent, ip string) (email string, key string, err error) {
	row := db.sql.QueryRowContext(ctx, `INSERT INTO sessions 
										(user_id, user_agent, ip)
										VALUES ($1, $2, $3::cidr)
										RETURNING key;`, userID, userAgent, ip)
	err = row.Scan(&key)
	if err != nil {
		return "", "", err
	}

	row = db.sql.QueryRowContext(ctx, `SELECT email
										FROM users
										WHERE id = $1`, userID)
	err = row.Scan(&email)
	if err != nil {
		return "", "", err
	}

	return email, key, nil
}

// A Session is a session
type Session struct {
	CreatedAt time.Time `json:"created_at"`
	UserAgent string    `json:"user_agent"`
	IP        string    `json:"ip"`
	Active    bool      `json:"active"`
}

// ListSessions lists all sessions a user has
func (db *DB) ListSessions(ctx context.Context, key string, page int) ([]*Session, error) {
	rows, err := db.sql.QueryContext(ctx, `SELECT created_at, user_agent, ip, active
							FROM sessions
							WHERE user_id = (SELECT user_id FROM sessions WHERE key = $1)
							LIMIT 25
							OFFSET $2`, key, page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*Session
	for rows.Next() {
		var s Session
		err = rows.Scan(&s.CreatedAt, &s.UserAgent, &s.IP, &s.Active)
		if err != nil {
			return nil, err
		}
		out = append(out, &s)
	}

	return out, nil
}

// DeactivateSession invalidates the current session
func (db *DB) DeactivateSession(ctx context.Context, key string) error {
	_, err := db.sql.QueryContext(ctx, `UPDATE 
										sessions
										SET (active) = (false)
										WHERE key = $1;`, key)

	return err
}

// ActiveUserFromKey returns the active user ID from the session key
func (db *DB) ActiveUserFromKey(ctx context.Context, key string) (string, error) {
	row := db.sql.QueryRowContext(ctx, `SELECT user_id 
										FROM sessions 
										WHERE key = $1
										AND active = true;`, key)

	var userID string
	err := row.Scan(&userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}

// AddFeed adds the given URL to the users default folder
// and links it across feed_folder
func (db *DB) AddFeed(ctx context.Context, sessionKey, folderID, title, plugin, feedURL string) (err error) {
	if folderID == "" {
		// ensure we don't shadow folderID
		var err error
		folderID, err = db.getDefaultFolderID(ctx, sessionKey)
		if err != nil {
			return err
		}
	}

	tx, err := db.sql.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	row := tx.QueryRowContext(ctx, `INSERT INTO feeds
							 (title, plugin, url)
							 VALUES ($1, $2, $3)
							 RETURNING id;`, title, plugin, feedURL)

	var feedID string
	err = row.Scan(&feedID)
	if err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			return fmt.Errorf("%s - %s", err, txErr)
		}
		return err
	}

	_, err = tx.ExecContext(ctx, `INSERT INTO feed_folders
							 (folder_id, feed_id)
							 VALUES ($1, $2);`, folderID, feedID)
	if err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			return fmt.Errorf("%s - %s", err, txErr)
		}
		return err
	}

	return tx.Commit()
}

// getDefaultFolderID returns a users default folder ID
func (db *DB) getDefaultFolderID(ctx context.Context, sessionKey string) (string, error) {
	row := db.sql.QueryRowContext(ctx, `SELECT id FROM folders 
							 			WHERE name = 'default' 
							 			AND user_id = (SELECT user_id FROM sessions WHERE key = $1);`, sessionKey)

	var fid string
	err := row.Scan(&fid)
	if err != nil {
		// if there is no default folder, go create one
		if err == sql.ErrNoRows {
			row := db.sql.QueryRowContext(ctx, `INSERT INTO folders
												(user_id)
												VALUES ((SELECT user_id FROM sessions WHERE key = $1 LIMIT 1))
												RETURNING id;`, sessionKey)

			err := row.Scan(&fid)
			if err != nil {
				return "", fmt.Errorf("could not create default folder: %s", err)
			}

		} else {
			return "", fmt.Errorf("could not find default folder: %s", err)
		}

	}

	return fid, nil
}

// RemoveFeed removes the given feed ID from the user
func (db *DB) RemoveFeed(ctx context.Context, sessionKey, folderID, feedID string) error {
	return nil
}

// GetFolders returns all of the folders for a user - if there are none it creates a
// default folder
func (db *DB) GetFolders(ctx context.Context, sessionKey string) ([]*Folder, error) {
	return nil, nil
}

// GetFeed returns a single feed
func (db *DB) GetFeed(ctx context.Context, sessionKey, feedID string) (*Feed, error) {
	return nil, nil
}

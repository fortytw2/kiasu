package pg

import "github.com/fortytw2/hydrocarbon"

// GetUser saves a user by ID
func (s *Store) GetUser(id string) (*hydrocarbon.User, error) {
	row := s.db.QueryRowx("SELECT * FROM users WHERE id = $1", id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var u hydrocarbon.User
	err := row.StructScan(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// GetUserByEmail gets a yser by email
func (s *Store) GetUserByEmail(email string) (*hydrocarbon.User, error) {
	row := s.db.QueryRowx("SELECT * FROM users WHERE email = $1", email)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var u hydrocarbon.User
	err := row.StructScan(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// CreateUser saves a user and returns it, with it's new ID
func (s *Store) CreateUser(u *hydrocarbon.User) (*hydrocarbon.User, error) {
	row := s.db.QueryRowx(`
		INSERT INTO users (email, encrypted_password, failed_login_count, active, confirmed, confirmation_token, token_created_at)
	    VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING *
	`, u.Email, u.EncryptedPassword, u.FailedLoginCount, u.Active, u.Confirmed, u.ConfirmationToken, u.TokenCreatedAt)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var usr hydrocarbon.User
	err := row.StructScan(&usr)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}

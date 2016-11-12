package kiasu

import "time"

// SessionStore stores sessions for users
type SessionStore interface {
	GetSession(id string) (*Session, error)
	GetSessionsByUserID(userID string, pg *Pagination) ([]Session, error)
	GetSessionByAccessToken(token string) (*Session, error)
	SaveSession(*Session) (*User, error)
}

// A Session is a single user session
type Session struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
	Token     string    `json:"token"`
}

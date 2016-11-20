package kiasu

import (
	"errors"
	"time"

	"github.com/fortytw2/abdi"
)

// Errors
var (
	ErrUserExists = errors.New("user already exists")
)

//go:generate moq -out store_moq_test.go . PrimitiveStore

// Store is responsible for persistent (or not) data storage and retrieval
// and abstracting that into business-logic level functions
type Store struct {
	Users        UserStore
	Sessions     SessionStore
	Feeds        FeedStore
	Posts        PostStore
	ReadStatuses ReadStatusStore

	EncryptionKey []byte
}

// PrimitiveStore encapsulates all primitive store types
type PrimitiveStore interface {
	UserStore
	SessionStore
	FeedStore
	PostStore
	ReadStatusStore
}

// NewStore builds a data storage layer out of the persistence primitives
func NewStore(ps PrimitiveStore, encryptionKey []byte) (*Store, error) {
	return &Store{
		Users:         ps,
		Sessions:      ps,
		Feeds:         ps,
		Posts:         ps,
		ReadStatuses:  ps,
		EncryptionKey: encryptionKey,
	}, nil
}

// CreateUser creates a new user from an email and password
func (s *Store) CreateUser(email, password string) (*User, error) {
	encPass, err := abdi.Hash(password, s.EncryptionKey)
	if err != nil {
		return nil, err
	}

	if _, err = s.Users.GetUserByEmail(email); err != nil {
		return nil, ErrUserExists
	}

	now := time.Now()
	u, err := s.Users.SaveUser(&User{
		CreatedAt:         &now,
		Email:             email,
		EncryptedPassword: *encPass,
		Confirmed:         false,
		TokenCreatedAt:    &now,
	})

	return u, err
}

// GetUserByToken returns the user with the given access token
func (s *Store) GetUserByToken(token string) (*User, error) {
	return nil, nil
}

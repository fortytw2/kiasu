package kiasu

//go:generate moq -out store_moq_test.go . PrimitiveStore

// Store is responsible for persistent (or not) data storage and retrieval
// and abstracting that into business-logic level functions
type Store struct {
	Users        UserStore
	Sessions     SessionStore
	Feeds        FeedStore
	Posts        PostStore
	ReadStatuses ReadStatusStore
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
func NewStore(ps PrimitiveStore) (*Store, error) {
	return &Store{
		Users:        ps,
		Sessions:     ps,
		Feeds:        ps,
		Posts:        ps,
		ReadStatuses: ps,
	}, nil
}

// CreateUser creates a new user from an email and password
func (s *Store) CreateUser(email, password string) error {
	_, err := s.Users.SaveUser(&User{
		Email: email,
	})

	return err
}

package pg

import "github.com/fortytw2/hydrocarbon"

// GetUser saves a user by ID
func (s *Store) GetUser(id string) (*hydrocarbon.User, error) {
	row := s.db.QueryRowx("SELECT * FROM users WHERE id = $1", id)

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

	var u hydrocarbon.User
	err := row.StructScan(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// SaveUser saves a user and returns it, with it's new ID
func (s *Store) SaveUser(u *hydrocarbon.User) (*hydrocarbon.User, error) {

	return nil, nil
}

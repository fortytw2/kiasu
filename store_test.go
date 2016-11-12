package kiasu

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	var users = []struct {
		Valid    bool
		Email    string
		Password string
	}{
		{true, "ian@fortytw2.com", "sa8dwu9djio23jl"},
		{false, "joe@barbados.com", "password"}, // invalid password
	}

	ps := &PrimitiveStoreMock{
		SaveUserFunc: func(u *User) (*User, error) {
			for _, us := range users {
				if us.Email == u.Email && us.Valid {
					u.ID = "random"
					return u, nil
				}
			}
			return nil, errors.New("invalid")
		},
		GetUserFunc: func(id string) (*User, error) {
			panic("TODO: mock out the SaveUser function")
		},
	}

	s, err := NewStore(ps)
	assert.Nil(t, err)

	for _, u := range users {
		err = s.CreateUser(u.Email, u.Password)
		if u.Valid {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}

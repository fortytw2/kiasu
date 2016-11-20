package kiasu_test

import (
	"testing"

	"github.com/fortytw2/kiasu"
	"github.com/fortytw2/kiasu/stores/bunt"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	var users = []struct {
		Valid    bool
		Dupe     bool
		Email    string
		Password string
	}{
		{true, false, "ian@fortytw2.com", "sa8dwu9djio23jl"},
		{false, false, "joe@barbados.com", "no"}, // invalid password
		{true, true, "ian@fortytw2.com", "sa8dwu9djio23jl"},
	}

	ps, err := bunt.NewMemStore()
	assert.Nil(t, err)

	s, err := kiasu.NewStore(ps, []byte{1, 2, 3, 4, 2})
	assert.Nil(t, err)

	for _, u := range users {
		_, err = s.CreateUser(u.Email, u.Password)
		if err != nil {
			if u.Dupe {
				assert.Equal(t, kiasu.ErrUserExists, err)
			}
			if !u.Valid {
				assert.NotNil(t, err)
			}
		}

		if u.Valid {
			assert.Nil(t, err)
		}
	}
}

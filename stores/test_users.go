package stores

import (
	"testing"

	"github.com/fortytw2/kiasu"
	"github.com/stretchr/testify/assert"
)

// TestUserStore ensures a given userStore does what it should do
func TestUserStore(t *testing.T, us kiasu.UserStore) {
	u, err := us.SaveUser(&kiasu.User{
		Email:             "ian@ian.com",
		EncryptedPassword: "we12312312",
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, u.ID)

	u2, err := us.GetUser(u.ID)

	assert.Nil(t, err)
	assert.Equal(t, u.ID, u2.ID)
}

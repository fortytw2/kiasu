package bunt

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/fortytw2/kiasu/stores"
	"github.com/stretchr/testify/assert"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestMemStore(t *testing.T) {
	s, err := NewMemStore()
	assert.Nil(t, err)
	assert.NotNil(t, s)

	stores.TestUserStore(t, s)
	stores.TestReadStatusStore(t, s)
	stores.TestFeedStore(t, s)
	stores.TestPostStore(t, s)
	stores.TestSessionStore(t, s)

	stores.FuzzUserStore(t, s, 256)
}

func TestDiskStore(t *testing.T) {
	path := fmt.Sprintf("/tmp/bunt-%d", rand.Int63())

	s, err := NewStore(path)
	assert.Nil(t, err)
	assert.NotNil(t, s)

	stores.TestUserStore(t, s)
	stores.TestReadStatusStore(t, s)
	stores.TestFeedStore(t, s)
	stores.TestPostStore(t, s)
	stores.TestSessionStore(t, s)

	stores.FuzzUserStore(t, s, 256)

	err = os.RemoveAll(path)
	if err != nil {
		fmt.Println("could not clean up after test")
	}
}

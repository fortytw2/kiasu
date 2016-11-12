package bunt

import (
	"fmt"
	"math/rand"
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
}

func TestDiskStore(t *testing.T) {
	s, err := NewStore(fmt.Sprintf("/tmp/bunt-%d", rand.Int63()))
	assert.Nil(t, err)
	assert.NotNil(t, s)

	stores.TestUserStore(t, s)
	stores.TestReadStatusStore(t, s)
	stores.TestFeedStore(t, s)
	stores.TestPostStore(t, s)
	stores.TestSessionStore(t, s)
}

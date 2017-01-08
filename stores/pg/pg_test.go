package pg

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	dockertest "gopkg.in/ory-am/dockertest.v2"

	"github.com/fortytw2/hydrocarbon/stores"
	"github.com/stretchr/testify/assert"
)

var dsn string

func TestMain(m *testing.M) {
	if testing.Short() {
		fmt.Println("skipping pgmigrate test")
		return
	}

	c, err := dockertest.ConnectToPostgreSQL(1, 5*time.Second, func(u string) bool {
		dsn = u
		return true
	})
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	// Run tests
	result := m.Run()
	c.KillRemove()
	os.Exit(result)
}

func TestStore(t *testing.T) {
	s, err := NewStore(dsn)
	assert.Nil(t, err)
	assert.NotNil(t, s)

	// stores.TestUserStore(t, s)
	// stores.TestReadStatusStore(t, s)
	stores.TestFeedStore(t, s)
	// stores.TestPostStore(t, s)
	stores.TestSessionStore(t, s)
}

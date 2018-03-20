package pg

import (
	"context"
	"errors"
	"testing"
)

func TestPG(t *testing.T) {
	db, shutdown := setupTestDB(t)
	defer shutdown()

	t.Run("users", userTests(db))
}

func userTests(db *DB) func(t *testing.T) {
	var createUserHelper = func(t *testing.T) string {
		id, _, err := db.CreateOrGetUser(context.Background(), "ian@hydrocarbon.io")
		if err != nil {
			t.Fatal(err)
		}
		return id
	}

	var cases = []testCase{
		{
			"create-user",
			func(t *testing.T) error {
				createUserHelper(t)
				return nil
			},
			nil,
		},
		{
			"create-token",
			func(t *testing.T) error {
				id := createUserHelper(t)
				token, err := db.CreateLoginToken(context.Background(), id, "Firefox", "192.168.1.21")
				if err != nil {
					return err
				}
				if token == "" {
					return errors.New("no token made")
				}

				return nil
			},
			nil,
		},
		{
			"create-session",
			func(t *testing.T) error {
				id := createUserHelper(t)
				_, _, err := db.CreateSession(context.Background(), id, "Firefox", "192.168.1.21")
				return err
			},
			nil,
		},
	}

	return func(t *testing.T) {
		runCases(t, db, cases)
	}
}

package bunt

import (
	"encoding/json"

	"github.com/fortytw2/kiasu"
	"github.com/fortytw2/kiasu/internal/uuid"
	"github.com/tidwall/buntdb"
)

const (
	feedPrefix = "feed:"
)

// GetFeed returns a feed by its ID
func (s *Store) GetFeed(id string) (*kiasu.Feed, error) {
	var f kiasu.Feed

	err := s.db.View(func(tx *buntdb.Tx) error {
		js, err := tx.Get(feedPrefix + id)
		if err != nil {
			return err
		}

		return json.Unmarshal([]byte(js), &f)
	})
	if err != nil {
		return nil, err
	}

	return &f, nil
}

// SaveFeed saves a feed and returns it with it's new ID
func (s *Store) SaveFeed(f *kiasu.Feed) (*kiasu.Feed, error) {
	id := uuid.NewV4()
	f.ID = id.String()

	buf, err := json.Marshal(f)
	if err != nil {
		return nil, err
	}

	err = s.db.Update(func(tx *buntdb.Tx) error {
		_, _, err = tx.Set(feedPrefix+id.String(), string(buf), &buntdb.SetOptions{Expires: false})
		return err
	})
	if err != nil {
		return nil, err
	}

	return f, nil
}

// GetFeeds returns and filters on feeds
func (s *Store) GetFeeds(pg *kiasu.Pagination) ([]kiasu.Feed, error) {
	return nil, nil
}

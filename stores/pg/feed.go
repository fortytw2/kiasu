package pg

import "github.com/fortytw2/hydrocarbon"

// GetFeed returns a feed by its ID
func (s *Store) GetFeed(id string) (*hydrocarbon.Feed, error) {
	// SELECT * FROM feeds WHERE id = $1

	return nil, nil
}

// SaveFeed saves a feed and returns it with it's new ID
func (s *Store) SaveFeed(*hydrocarbon.Feed) (*hydrocarbon.Feed, error) {
	// INSERT

	return nil, nil
}

// GetFeeds returns and filters on feeds
func (s *Store) GetFeeds(pg *hydrocarbon.Pagination) ([]hydrocarbon.Feed, error) {
	// SELECT * FROM feeds LIMIT $1 OFFSET $2

	return nil, nil
}

package hydrocarbon

import (
	"time"

	"github.com/lib/pq"
)

// FeedStore saves and persists feeds and posts
type FeedStore interface {
	GetFeed(id string) (*Feed, error)
	CreateFeed(*Feed) (*Feed, error)
	GetFeeds(pg *Pagination) ([]Feed, error)
}

// A Feed is a single encapsulating unit around a source of news / posts / whatever
type Feed struct {
	ID         string `json:"id"`
	Plugin     string `json:"plugin,omitempty"`
	InitialURL string `json:"initial_url,omitempty"`

	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	RefreshedAt pq.NullTime `json:"refreshed_at"`

	URL         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`

	HexColor string `json:"hex_color"`
	IconURL  string `json:"icon_url"`

	UnreadCount int `json:"unread_count"`
}

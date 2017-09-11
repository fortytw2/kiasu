package hydrocarbon

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// A FeedStore is an interface used to seperate the FeedAPI from knowledge of the
// actual underlying database
type FeedStore interface {
	AddFeed(ctx context.Context, sessionKey, folderID, title, plugin, feedURL string) error
	RemoveFeed(ctx context.Context, sessionKey, folderID, feedID string) error

	// GetFolders should not return any Posts in the nested Feeds
	GetFolders(ctx context.Context, sessionKey string) ([]*Folder, error)
	GetFeed(ctx context.Context, feedID string, limit, offset int) (*Feed, error)
}

// A Folder holds a collection of feeds
type Folder struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Feeds []*Feed `json:"feeds"`
}

// A Feed is a collection of posts
type Feed struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Plugin    string    `json:"plugin"`
	BaseURL   string    `json:"base_url"`

	Posts []*Post `json:"posts,omitempty"`
}

// A Post is a single post on a feed
type Post struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	OriginalURL string    `json:"original_url"`

	Title  string `json:"title"`
	Author string `json:"author"`
	Body   string `json:"body"`

	Extra map[string]interface{} `json:"extra"`
}

func (p *Post) ContentHash() string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s:%s:%s", p.Title, p.Author, p.Body)))

	var b []byte
	return hex.EncodeToString(h.Sum(b))
}

// FeedAPI encapsulates everything related to user management
type FeedAPI struct {
	s FeedStore
	p *PluginList
}

// NewFeedAPI returns a new Feed API
func NewFeedAPI(s FeedStore, p *PluginList) *FeedAPI {
	return &FeedAPI{
		s: s,
		p: p,
	}
}

// AddFeed adds the specified feed to the given user
// if folder_id is left out, the feed is added to the users "default" folder
func (fa *FeedAPI) AddFeed(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("X-Hydrocarbon-Key")
	if key == "" {
		writeErr(w, errors.New("no api key present"))
		return
	}

	var feed struct {
		FolderID string `json:"folder_id,omitempty"`
		Plugin   string `json:"plugin"`
		URL      string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&feed)
	if err != nil {
		writeErr(w, err)
		return
	}

	if feed.URL == "" || feed.Plugin == "" {
		writeErr(w, errors.New("one of url or plugin is empty"))
		return
	}

	plug, err := fa.p.ByName(feed.Plugin)
	if err != nil {
		writeErr(w, err)
		return
	}

	title, baseURL, err := plug.Info(r.Context(), feed.URL)
	if err != nil {
		writeErr(w, err)
		return
	}

	err = fa.s.AddFeed(r.Context(), key, feed.FolderID, title, feed.Plugin, baseURL)
	if err != nil {
		writeErr(w, err)
		return
	}
}

// RemoveFeed removes the given feed from the users list
func (fa *FeedAPI) RemoveFeed(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("X-Hydrocarbon-Key")
	if key == "" {
		writeErr(w, errors.New("no api key present"))
		return
	}

	var feed struct {
		FolderID string `json:"folder_id"`
		FeedID   string `json:"feed_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&feed)
	if err != nil {
		writeErr(w, err)
		return
	}

	if feed.FeedID == "" || feed.FolderID == "" {
		writeErr(w, errors.New("no feed or folder ID sent"))
		return
	}

	err = fa.s.RemoveFeed(r.Context(), key, feed.FolderID, feed.FeedID)
	if err != nil {
		writeErr(w, err)
		return
	}
}

// GetFolders writes all of a users folders out
func (fa *FeedAPI) GetFolders(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("X-Hydrocarbon-Key")
	if key == "" {
		writeErr(w, errors.New("no api key present"))
		return
	}

	folders, err := fa.s.GetFolders(r.Context(), key)
	if err != nil {
		writeErr(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(folders)
	if err != nil {
		writeErr(w, err)
		return
	}
}

// GetFeed writes a specific feed
func (fa *FeedAPI) GetFeed(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("X-Hydrocarbon-Key")
	if key == "" {
		writeErr(w, errors.New("no api key present"))
		return
	}

	var id struct {
		ID string `json:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		writeErr(w, err)
		return
	}

	feed, err := fa.s.GetFeed(r.Context(), id.ID, 50, 0)
	if err != nil {
		writeErr(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(feed)
	if err != nil {
		writeErr(w, err)
		return
	}
}

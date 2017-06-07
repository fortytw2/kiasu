package hydrocarbon

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// A FeedStore is an interface used to seperate the FeedAPI from knowledge of the
// actual underlying database
type FeedStore interface {
	AddFeed(ctx context.Context, sessionKey, folderID, plugin, feedURL string) error
	RemoveFeed(ctx context.Context, sessionKey, folderID, feedID string) error
}

// FeedAPI encapsulates everything related to user management
type FeedAPI struct {
	s FeedStore
	p PluginList
}

// NewFeedAPI returns a new Feed API
func NewFeedAPI(s FeedStore, p PluginList) *FeedAPI {
	return &FeedAPI{
		s: s,
		p: p,
	}
}

// AddFeed adds the specified feed to the given user
func (fa *FeedAPI) AddFeed(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("X-Hydrocarbon-Key")
	if key == "" {
		writeErr(w, errors.New("no api key present"))
		return
	}

	var feed struct {
		FolderID string `json:"folder_id"`
		Plugin   string `json:"plugin"`
		URL      string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&feed)
	if err != nil {
		writeErr(w, err)
		return
	}

	if feed.URL == "" || feed.FolderID == "" || feed.Plugin == "" {
		writeErr(w, errors.New("one of url, plugin, folder_id is empty"))
		return
	}

	err = fa.s.AddFeed(r.Context(), key, feed.FolderID, feed.Plugin, feed.URL)
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
		FeedID   string `json:"id"`
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

package rss

import (
	"context"
	"net/http"

	"github.com/fortytw2/hydrocarbon"
)

// Reader implements hydrocarbon.Plugin for RSS feeds
type Reader struct {
	Client *http.Client
}

// Name returns the name of the plugin
func (r *Reader) Name() string {
	return "rss"
}

// Info sanitizes a URL send to this plugin, maybe by making a test request
func (r *Reader) Info(ctx context.Context, inputURL string) (title, baseURL string, err error) {
	return "", "", nil
}

// Fetch performs the actual scraping shenanigans
func (r *Reader) Fetch(ctx context.Context, baseURL string) ([]*hydrocarbon.Post, error) {
	return nil, nil
}

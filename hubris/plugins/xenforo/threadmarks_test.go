package xenforo

import (
	"net/http"
	"testing"
	"time"

	"github.com/fortytw2/kiasu"
	"github.com/fortytw2/watney"
)

func TestExtractor(t *testing.T) {
	tr := watney.Configure(http.DefaultTransport, t)
	c := &http.Client{
		Transport: tr,
	}

	defer watney.Save(c)

	e := NewExtractor(c)
	_, err := e.FindSince(&kiasu.Feed{
		URL: "https://forums.spacebattles.com/threads/twinning-s-worm-altpower.408788/threadmarks",
	}, time.Time{})

	if err != nil {
		t.Fatal(err)
	}
}

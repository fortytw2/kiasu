package hubris

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
)

// ErrCountryNotFound is returned when a request can't be made in that country
type ErrCountryNotFound struct {
	Country string
}

func (e ErrCountryNotFound) Error() string {
	return fmt.Sprintf("country %s not found in proxy configuration", e.Country)
}

// A Client is used to make all HTTP requests to the outside world
type Client interface {
	Do(*http.Request, *ReqOpts) (*http.Response, error)
}

// ReqOpts are used to process the given HTTP request
type ReqOpts struct {
	Country   string
	CookieJar *cookiejar.Jar
}

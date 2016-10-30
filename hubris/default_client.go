package hubris

import (
	"net/http"
	"sync"

	"github.com/fortytw2/kiasu/internal/cleanhttp"
)

type client struct {
	cMu sync.Mutex
	c   *http.Client
}

// StandardClient is a
func StandardClient() Client {
	return &client{
		c: cleanhttp.DefaultPooledClient(),
	}
}

func (c *client) Do(r *http.Request, ro *ReqOpts) (*http.Response, error) {
	if ro.Country != AnyCountry {
		return nil, ErrCountryNotFound{
			Country: ro.Country,
		}
	}

	c.cMu.Lock()
	defer c.cMu.Unlock()
	c.c.Jar = ro.CookieJar

	return c.Do(r, ro)
}

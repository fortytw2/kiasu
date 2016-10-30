package hubris

import (
	"github.com/fortytw2/kiasu"
	"github.com/go-kit/kit/log"
)

// A Run is a single, complete run of a plugin + config
// typically, these should complete in a matter of seconds
type Run struct {
	c   Client
	cfg Config
	p   *Plugin

	l log.Logger
	f kiasu.FeedStore

	stats *Stats
}

func (r *Run) Start() {

}

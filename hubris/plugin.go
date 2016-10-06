package flechette

import (
	"errors"

	"github.com/moovweb/rubex"
)

// Plugin Errors
var (
	ErrPluginUninitialized = errors.New("plugin not initialized")
	ErrNoHandlerForRoute   = errors.New("no handler for route")
)

// A Plugin scrapes something and returns something
type Plugin struct {
	Name string

	Configs func(Client) ([]Config, error)
	// Entrypoint starts the scrape with the given config
	Entrypoint func(Client, *Config) ([]Info, error)
	// User input routes function
	Routes map[string]Handler

	// parsed routes into regexps
	realRoutes map[*rubex.Regexp]Handler
}

// A Handler processes a given task URL
type Handler func(Client, string) ([]Info, error)

// NewPlugin constructs a plugin from the given functions
func NewPlugin(name string,
	configs func(Client) ([]Config, error),
	entrypoint func(Client, *Config) ([]Info, error),
	Routes map[string]Handler) (*Plugin, error) {
	return nil, nil
}

// GetHandler returns the Handler for a given URL, if one is present.
// URLs are always full qualified `https://www.google.com/potatoes?xd=34`
func (p *Plugin) GetHandler(url string) (Handler, error) {
	if p.realRoutes == nil {
		return nil, ErrPluginUninitialized
	}

	for re, h := range p.realRoutes {
		if re.MatchString(url) {
			return h, nil
		}
	}

	return nil, ErrNoHandlerForRoute
}

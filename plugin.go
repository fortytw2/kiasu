package hydrocarbon

import (
	"context"
	"errors"
	"sync"
)

// A Plugin is responsible for fetching and maintaining synchronized feeds
type Plugin interface {
	Name() string
	Info(ctx context.Context, inputURL string) (title, baseURL string, err error)
	Fetch(ctx context.Context, baseURL string) ([]*Post, error)
}

// A PluginList is a collection of plugins
type PluginList struct {
	pMu     sync.RWMutex
	plugins map[string]Plugin
}

// NewPluginList builds a new plugin list
func NewPluginList(plugins ...Plugin) *PluginList {
	m := make(map[string]Plugin)
	for _, p := range plugins {
		m[p.Name()] = p
	}

	return &PluginList{
		plugins: m,
	}
}

// ByName returns the plugin that is named such
func (pl *PluginList) ByName(n string) (Plugin, error) {
	pl.pMu.RLock()
	defer pl.pMu.RUnlock()

	plug, ok := pl.plugins[n]
	if !ok {
		return nil, errors.New("plugin not found")
	}

	return plug, nil
}

package hydrocarbon

import "errors"

// A Plugin is responsible for fetching and maintaining synchronized feeds
type Plugin interface {
	Name() string
	Info(inputURL string) (title, baseURL string, err error)
	Fetch(baseURL string)
}

// A PluginList is a collection of plugins
type PluginList map[string]Plugin

// ByName returns the plugin that is named such
func (pl PluginList) ByName(n string) (Plugin, error) {
	return nil, errors.New("not yet implemented")
}

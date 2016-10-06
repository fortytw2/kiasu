package hubris

// Info is returned by a handler, can add more URLs, or be processed by a
// backend
type Info interface {
	Type() string
	URLs() []string
}

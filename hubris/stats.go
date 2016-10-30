package hubris

import (
	"sync/atomic"
	"time"
)

// Stats is a fast metrics container
type Stats struct {
	started time.Time
	ended   time.Time

	urls  uint64
	posts uint64
}

func NewStats() *Stats {
	return &Stats{
		started: time.Now(),
	}
}

func (s *Stats) AddURL() {
	atomic.AddUint64(&s.urls, 1)
}

func (s *Stats) AddPost() {
	atomic.AddUint64(&s.posts, 1)
}

func (s *Stats) End() {

}

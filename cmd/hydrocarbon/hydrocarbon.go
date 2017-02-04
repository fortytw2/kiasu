package main

import (
	"net/http"
	"os"

	"github.com/fortytw2/hydrocarbon"
	"github.com/fortytw2/hydrocarbon/internal/log"
	"github.com/fortytw2/hydrocarbon/plugins/fanfictionnet"
	"github.com/fortytw2/hydrocarbon/plugins/xenforo"
	"github.com/fortytw2/hydrocarbon/stores/pg"
	"github.com/fortytw2/hydrocarbon/web"
	raven "github.com/getsentry/raven-go"
)

func main() {
	l := log.NewContext(log.NewLogfmtLogger(os.Stdout)).With("ts", log.DefaultTimestampUTC)
	l.Log("msg", "launching hydrocarbon", "port", getPort())

	if os.Getenv("POSTGRES_DSN") == "" {
		l.Log("msg", "env POSTGRES_DSN must be set")
		return
	}

	if sentryDSN := os.Getenv("SENTRY_DSN"); sentryDSN != "" {
		raven.SetDSN(sentryDSN)
	} else {
		l.Log("msg", "sentry dsn not set, not reporting errors")
	}

	store, err := pg.NewStore(l, os.Getenv("POSTGRES_DSN"))
	if err != nil {
		l.Log("msg", "cannot start", "error", err)
		return
	}

	s, err := hydrocarbon.NewStore(store, []byte{1, 2, 3, 4})
	if err != nil {
		l.Log("msg", "cannot start", "error", err)
		return
	}

	go launchScraper(l, s)

	r := web.Routes(s, l)
	err = http.ListenAndServe(getPort(), r)
	if err != nil {
		l.Log("msg", "cannot start", "error", err)
		return
	}
}

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func launchScraper(l log.Logger, s *hydrocarbon.Store) {
	plugins := map[string]hydrocarbon.Instantiator{
		"xenforo":       xenforo.NewPlugin,
		"fanfictionnet": fanfictionnet.NewPlugin,
	}

	hydrocarbon.ScrapeLoop(l, s, plugins)
}

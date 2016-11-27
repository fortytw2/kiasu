package main

import (
	"net/http"
	"os"

	"github.com/felixge/httpsnoop"
	"github.com/fortytw2/hydrocarbon"
	"github.com/fortytw2/hydrocarbon/internal/log"
	"github.com/fortytw2/hydrocarbon/stores/bunt"
	"github.com/fortytw2/hydrocarbon/web"
)

func main() {
	l := log.NewContext(log.NewLogfmtLogger(os.Stdout)).With("ts", log.DefaultTimestampUTC)
	l.Log("msg", "launching hydrocarbon", "port", getPort())

	memStore, err := bunt.NewMemStore()
	if err != nil {
		l.Log("msg", "cannot start", "error", err)
		return
	}

	s, err := hydrocarbon.NewStore(memStore, []byte{1, 2, 3, 4})
	if err != nil {
		l.Log("msg", "cannot start", "error", err)
		return
	}

	r := web.Routes(s)
	err = http.ListenAndServe(getPort(), httpLogger(r, l))
	if err != nil {
		l.Log("msg", "could not start hydrocarbon", "error", err)
	}
}

func httpLogger(router http.Handler, l log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		m := httpsnoop.CaptureMetrics(router, w, req)
		l.Log("msg", "request", "method", req.Method, "url", req.URL.String(), "code", m.Code, "time", m.Duration, "bytes", m.Written)
	})
}

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

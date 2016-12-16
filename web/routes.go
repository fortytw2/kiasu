package web

import (
	"net/http"
	"time"

	"github.com/felixge/httpsnoop"
	"github.com/fortytw2/httpkit"
	"github.com/fortytw2/hydrocarbon"
	"github.com/fortytw2/hydrocarbon/internal/log"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

const (
	homeURL           = "/"
	loginURL          = "/login"
	registerURL       = "/register"
	confirmTokenURL   = "/confirm"
	forgotPasswordURL = "/password_reset"

	feedsURL        = "/feeds"
	reorderFeedsURL = "/reorder_feeds"
	oneFeedURL      = "/feeds/:feedID"

	onePostURL    = "/posts/:postID"
	readStatusURL = "/mark_read"
)

//go:generate ftmpl -targetgo ./templates_generated.go templates/

// Routes returns all routes for this application
func Routes(s *hydrocarbon.Store, l log.Logger) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(httplog(l))
	r.Use(middleware.Timeout(5 * time.Second))
	r.Use(middleware.DefaultCompress)

	r.Handle("/hydrocarbon.min.css", httpkit.ErrorHandler(Stylesheet))

	r.Handle(homeURL, httpkit.ErrorHandler(renderHome))

	r.Get(registerURL, renderRegister)
	r.Post(registerURL, newUser)
	r.Get(confirmTokenURL, confirmUser)

	r.Post(forgotPasswordURL, forgotPassword)

	r.Get(loginURL, renderLogin)
	r.Post(loginURL, newSession)
	r.Delete(loginURL, deleteSession)

	r.Get(feedsURL, renderFeed)
	r.Get(oneFeedURL, renderFeed)

	r.Post(feedsURL, addFeed)
	r.Post(reorderFeedsURL, reorderFeeds)
	r.Delete(feedsURL, deleteFeed)

	r.Get(onePostURL, renderPost)
	r.Post(readStatusURL, markRead)

	return r
}

func httplog(l log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, req *http.Request) {
			m := httpsnoop.CaptureMetrics(next, w, req)
			l.Log("msg", "request", "method", req.Method, "url", req.URL.String(), "code", m.Code, "time", m.Duration, "bytes", m.Written)
		}

		return http.HandlerFunc(fn)
	}
}

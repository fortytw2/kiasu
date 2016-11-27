package web

import (
	"github.com/fortytw2/hydrocarbon"
	"github.com/neustar/httprouter"
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
func Routes(s *hydrocarbon.Store) *httprouter.Router {
	r := httprouter.New()

	r.GET(homeURL, renderHome)

	r.GET(registerURL, renderRegister)
	r.POST(registerURL, newUser)
	r.GET(confirmTokenURL, confirmUser)

	r.POST(forgotPasswordURL, forgotPassword)

	r.GET(loginURL, renderLogin)
	r.POST(loginURL, newSession)
	r.DELETE(loginURL, deleteSession)

	r.GET(feedsURL, renderFeed)
	r.GET(oneFeedURL, renderFeed)

	r.POST(feedsURL, addFeed)
	r.POST(reorderFeedsURL, reorderFeeds)
	r.DELETE(feedsURL, deleteFeed)

	r.GET(onePostURL, renderPost)
	r.POST(readStatusURL, markRead)

	return r
}

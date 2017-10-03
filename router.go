package hydrocarbon

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/bouk/httprouter"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/fortytw2/hydrocarbon/public"
)

//go:generate bash -c "pushd ui && yarn run build-js && popd"
//go:generate bash -c "pushd ui && yarn run build-css && popd"
//go:generate bash -c "go-bindata -pkg public -mode 0644 -modtime 499137600 -o public/assets_generated.go ui/build/..."

// NewRouter configures a new http.Handler that serves hydrocarbon
func NewRouter(ua *UserAPI, fa *FeedAPI, domain string) http.Handler {
	m := httprouter.New()

	fs := http.FileServer(
		&assetfs.AssetFS{
			Asset:     rewriteAsset(domain, public.Asset),
			AssetDir:  public.AssetDir,
			AssetInfo: public.AssetInfo,
			Prefix:    "ui/build/",
		})

	m.Handle("GET", "/static/*file", http.StripPrefix("/static/", fs).ServeHTTP)
	m.Handle("GET", "/favicon.ico", fs.ServeHTTP)

	// serve the single page app for every other route, it has a 404 page builtin
	m.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		buf := public.MustAsset("ui/build/index.html")
		w.Header().Set("Content-Type", "text/html")
		w.Write(buf)
	})

	// login tokens
	m.POST("/api/token/create", ua.RequestToken)

	// payment managemnet
	m.POST("/api/payment/create", ua.CreatePayment)

	// api keys
	m.POST("/api/key/create", ua.Activate)
	m.POST("/api/key/delete", ua.Deactivate)
	m.POST("/api/key/list", ua.ListSessions)

	// feed management
	m.POST("/api/feed/create", fa.AddFeed)
	m.POST("/api/feed/delete", fa.RemoveFeed)
	m.POST("/api/feed/list", fa.GetFeed)

	// folder management
	m.POST("/api/folder/list", fa.GetFolders)

	if httpsOnly(domain) {
		return redirectHTTPS(m)
	}

	return m
}

func rewriteAsset(domain string, fileFunc func(name string) ([]byte, error)) func(name string) ([]byte, error) {
	var cacheMu sync.RWMutex
	cache := make(map[string][]byte)

	return func(name string) ([]byte, error) {
		// return rewritten assets from cache if possible
		cacheMu.RLock()
		if body, ok := cache[name]; ok {
			cacheMu.RUnlock()
			return body, nil
		}
		cacheMu.RUnlock()

		if strings.Contains(name, ".min.js") {
			buf, err := fileFunc(name)
			if err != nil {
				return nil, err
			}
			buf = bytes.Replace(buf, []byte("URL_ENDPOINT_CHANGE_ME"), []byte(domain+"/api"), -1)

			// add rewritten assets to cache
			cacheMu.Lock()
			cache[name] = buf
			cacheMu.Unlock()

			return buf, nil
		}
		return fileFunc(name)
	}
}

func httpsOnly(domain string) bool {
	u, err := url.Parse(domain)
	if err != nil {
		panic(err)
	}
	return u.Scheme == "https"
}

func redirectHTTPS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Forwarded-Proto") == "http" {
			r.URL.Scheme = "https"
			http.Redirect(w, r, r.URL.String(), http.StatusTemporaryRedirect)
			return
		}

		next.ServeHTTP(w, r)
	})
}

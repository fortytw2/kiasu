// Package httputil is borrowed from github.com/shurcool/home/httputil
package httputil

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

// ErrorHandler factors error handling out of the HTTP handler.
type ErrorHandler func(w http.ResponseWriter, req *http.Request) error

func (h ErrorHandler) Func() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

func (h ErrorHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	rw := &responseWriter{ResponseWriter: w}
	err := h(rw, req)
	switch {
	case err != nil && rw.WroteHeader:
		// The header has already been written, so it's too late to send
		// a different status code. Just log the error and move on.
		log.Println(err)
	case IsMethodError(err):
		w.Header().Set("Allow", strings.Join(err.(MethodError).Allowed, ", "))
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
	case IsRedirect(err):
		http.Redirect(w, req, err.(Redirect).URL, http.StatusSeeOther)
	case IsHTTPError(err):
		http.Error(w, err.Error(), err.(HTTPError).Code)
	case IsJSONResponse(err):
		w.Header().Set("Content-Type", "application/json")
		jw := json.NewEncoder(w)
		err := jw.Encode(err.(JSONResponse).V)
		if err != nil {
			log.Println("error encoding JSONResponse:", err)
		}
	case os.IsNotExist(err):
		http.Error(w, err.Error(), http.StatusNotFound)
	case os.IsPermission(err):
		http.Error(w, err.Error(), http.StatusUnauthorized)
	default:
		// Do nothing.
	case err != nil:
		log.Println(err)
		// TODO: Only display error details to SiteAdmin users?
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// responseWriter wraps a real http.ResponseWriter and captures
// whether or not the header has been written.
type responseWriter struct {
	http.ResponseWriter

	WroteHeader bool // Write or WriteHeader was called.
}

func (rw *responseWriter) Write(p []byte) (n int, err error) {
	rw.WroteHeader = true
	return rw.ResponseWriter.Write(p)
}
func (rw *responseWriter) WriteHeader(code int) {
	rw.WroteHeader = true
	rw.ResponseWriter.WriteHeader(code)
}

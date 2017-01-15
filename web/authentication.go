package web

import (
	"context"
	"net/http"

	"github.com/fortytw2/hydrocarbon"
	"github.com/fortytw2/hydrocarbon/internal/log"
)

type userKey struct{}

var userCookieToken = "hydrocarbontok"

func loggedIn(r *http.Request) *hydrocarbon.User {
	u := r.Context().Value(userKey{})
	if u == nil {
		return nil
	}
	return u.(*hydrocarbon.User)
}

func authenticate(s *hydrocarbon.Store, l log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, req *http.Request) {
			token, err := req.Cookie(userCookieToken)
			if err != nil {
				next.ServeHTTP(w, req)
				return
			}

			user, err := s.GetUserByToken(token.Value)
			if err != nil {
				next.ServeHTTP(w, req)
				return
			}

			newCtx := context.WithValue(req.Context(), userKey{}, user)
			next.ServeHTTP(w, req.WithContext(newCtx))
		}

		return http.HandlerFunc(fn)
	}
}

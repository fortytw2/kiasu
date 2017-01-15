package web

import (
	"errors"
	"net/http"
	"time"

	"github.com/fortytw2/hydrocarbon"
	"github.com/fortytw2/hydrocarbon/internal/httputil"
	"github.com/fortytw2/hydrocarbon/internal/token"
)

// login renders a dummy page for logging in
func renderLogin(w http.ResponseWriter, r *http.Request) error {
	out, err := TMPLERRlogin("Hydrocarbon", loggedIn(r))
	if err != nil {
		return httputil.Wrap(err, http.StatusInternalServerError)
	}

	_, err = w.Write([]byte(out))
	if err != nil {
		return httputil.Wrap(err, http.StatusInternalServerError)
	}

	return nil
}

// register displays a sign up page
func renderRegister(w http.ResponseWriter, r *http.Request) error {
	out := TMPLregister("Hydrocarbon", loggedIn(r))
	_, err := w.Write([]byte(out))
	if err != nil {
		return httputil.Wrap(err, http.StatusInternalServerError)
	}

	return nil
}

type registration struct {
	Email    string
	Password string
}

func (r *registration) Valid() error {
	if r.Email == "" {
		return errors.New("email is required")
	}
	if r.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

// newUser processes a post request
func newUser(s *hydrocarbon.Store) httputil.ErrorHandler {
	return func(w http.ResponseWriter, req *http.Request) error {
		r := registration{
			Email:    req.FormValue("email"),
			Password: req.FormValue("password"),
		}

		err := r.Valid()
		if err != nil {
			return err
		}

		user, err := s.CreateUser(r.Email, r.Password)
		if err != nil {
			return err
		}

		newToken, err := token.GenerateRandomString(32)
		if err != nil {
			return err
		}
		// right here should send a confirmation email
		// user.ConfirmationToken
		sesh, err := s.Sessions.CreateSession(&hydrocarbon.Session{
			UserID:    user.ID,
			ExpiresAt: time.Now().Add(14 * 24 * time.Hour),
			Token:     newToken,
		})
		if err != nil {
			return err
		}

		http.SetCookie(w, &http.Cookie{
			Name:  userCookieToken,
			Value: sesh.Token,
		})

		http.Redirect(w, req, "/feeds", http.StatusSeeOther)

		return nil
	}
}

// newSession creates a new session
func newSession(w http.ResponseWriter, r *http.Request) {

}

// confirmUser asserts that the user has a valid email
func confirmUser(w http.ResponseWriter, r *http.Request) {

}

// forgotPassword sends a reset email
func forgotPassword(w http.ResponseWriter, r *http.Request) {

}

// deleteSession invalidates an existing session
func deleteSession(w http.ResponseWriter, r *http.Request) {

}

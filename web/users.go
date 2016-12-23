package web

import (
	"fmt"
	"net/http"

	"github.com/fortytw2/hydrocarbon"
	"github.com/fortytw2/hydrocarbon/internal/httputil"
	"github.com/mholt/binding"
)

// login renders a dummy page for logging in
func renderLogin(w http.ResponseWriter, r *http.Request) error {
	out, err := TMPLERRlogin("Hydrocarbon", false, 0)
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
	out := TMPLregister("Hydrocarbon", false, 0)
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

// Then provide a field mapping (pointer receiver is vital)
func (r *registration) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&r.Email:    "email",
		&r.Password: "password",
	}
}

// newUser processes a post request
func newUser(s *hydrocarbon.Store) httputil.ErrorHandler {
	return func(w http.ResponseWriter, req *http.Request) error {
		r := new(registration)
		errs := binding.Bind(req, r)
		if errs.Handle(w) {
			return nil
		}
		fmt.Fprintf(w, "From:    %s\n", r.Email)
		fmt.Fprintf(w, "Message: %s\n", r.Password)
		return nil
	}
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

// newSession creates a new session
func newSession(w http.ResponseWriter, r *http.Request) {

}

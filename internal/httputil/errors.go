package httputil

import (
	"fmt"
	"strings"
)

// MethodError is an error type used for methods that aren't allowed.
type MethodError struct {
	Allowed []string // Allowed methods.
}

func (m MethodError) Error() string {
	return fmt.Sprintf("method should be %v", strings.Join(m.Allowed, " or "))
}

func IsMethodError(err error) bool {
	_, ok := err.(MethodError)
	return ok
}

// Redirect is an error type used for representing a simple HTTP redirection.
type Redirect struct {
	URL string
}

func (r Redirect) Error() string { return fmt.Sprintf("redirecting to %s", r.URL) }

func IsRedirect(err error) bool {
	_, ok := err.(Redirect)
	return ok
}

func Wrap(err error, code int) error {
	return HTTPError{
		Err:  err,
		Code: code,
	}
}

// HTTPError is an error type used for representing a non-nil error with a status code.
type HTTPError struct {
	Code int
	Err  error // Not nil.
}

// Error returns HTTPError.err.Error().
func (h HTTPError) Error() string { return h.Err.Error() }

func IsHTTPError(err error) bool {
	_, ok := err.(HTTPError)
	return ok
}

// JSONResponse is an error type used for representing a JSON response.
type JSONResponse struct {
	V interface{}
}

func (JSONResponse) Error() string { return "JSONResponse" }

func IsJSONResponse(err error) bool {
	_, ok := err.(JSONResponse)
	return ok
}

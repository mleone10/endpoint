package middleware

import (
	"net/http"
)

type errorWrapper struct {
	http.ResponseWriter
	status int
}

func (e *errorWrapper) WriteHeader(code int) {
	e.status = code
	e.ResponseWriter.WriteHeader(code)
}

// ErrorReporter is a middleware which logs errors and, in the case of HTTP 500 errors, converts the error message to a common user-facing message.
func ErrorReporter() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ew := &errorWrapper{w, http.StatusOK}
			next.ServeHTTP(ew, r)
			if ew.status == http.StatusInternalServerError {
				// TODO: Overwrite with generic error text
			}
		})
	}
}

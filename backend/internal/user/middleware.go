package user

import "net/http"

// OrMiddleware is a middleware which makes a runtime decision between two provided middleware.  If cond evaluates to true, middleware a is used, else middleware b is used.
func OrMiddleware(cond bool, a, b func(next http.Handler) http.Handler) func(next http.Handler) http.Handler {
	if cond {
		return a
	}
	return b
}

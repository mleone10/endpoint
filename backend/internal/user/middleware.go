package user

import (
	"fmt"
	"net/http"
)

// OrMiddleware is a middleware which makes a runtime decision between two provided middleware.  If cond evaluates to true, middleware a is used, else middleware b is used.
func OrMiddleware(cond bool, a, b func(next http.Handler) http.Handler) func(next http.Handler) http.Handler {
	if cond {
		return a
	}
	return b
}

// AuthTokenVerifier is a middleware which verifies an Authorization header JWT using the Firebase Admin SDK.
func AuthTokenVerifier() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		authenticator, err := NewAuthenticator()
		if err != nil {
			panic(fmt.Sprintf("error initializing auth middleware: %v", err))
		}
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Update Authorization header to include "Bearer" identifier
			userID, err := authenticator.VerifyJWT(r.Context(), r.Header.Get("Authorization"))
			if err != nil {
				http.Error(w, "failed to verify authentication token", http.StatusForbidden)
				return
			}

			ctx := NewContextWithID(r.Context(), NewID(userID))
			next.ServeHTTP(w, r.Clone(ctx))
		})
	}
}

// AuthStubber is a middleware which ignores the provided Authorization header and instead injects a static user ID into the request.
func AuthStubber() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := NewContextWithID(r.Context(), NewID("testUserID"))
			next.ServeHTTP(w, r.Clone(ctx))
		})
	}
}

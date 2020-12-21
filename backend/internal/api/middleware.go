package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/mleone10/endpoint/internal/user"
)

type userIDKeyType string

const userIDKey userIDKeyType = "userID"

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
			userID, err := authenticator.VerifyJWT(r.Context(), strings.Split("Bearer ", r.Header.Get("Authorization"))[1])
			if err != nil {
				http.Error(w, "failed to verify authentication token", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r.Clone(context.WithValue(r.Context(), userIDKey, user.ID(userID))))
		})
	}
}

// AuthStubber is a middleware which ignores the provided Authorization header and instead injects a static user ID into the request.
func AuthStubber() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.Clone(context.WithValue(r.Context(), userIDKey, user.ID("testUserID"))))
		})
	}
}

// IDFromContext returns the ID stored in ctx, if any.
func IDFromContext(ctx context.Context) (user.ID, error) {
	uid, ok := ctx.Value(userIDKey).(user.ID)
	if !ok {
		return "", fmt.Errorf("could not retrieve user ID from context")
	}
	return uid, nil
}

// KeyTokenVerifier is a middleware which confirms that the given API key has sufficient permissions to perform the target operation.
func KeyTokenVerifier() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract API key from Authorization header (separate method, so stations/ can use it)
			// Use key and station ID to check database for permissions (read/write)
			next.ServeHTTP(w, r)
		})
	}
}

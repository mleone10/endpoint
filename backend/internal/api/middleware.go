package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/mleone10/endpoint/internal/account"
)

type userIDKeyType string

const userIDKey userIDKeyType = "userID"

// Authenticator describes a client which can validate an identity token JWT.
type Authenticator interface {
	VerifyJWT(context.Context, string) (string, error)
}

// authTokenVerifier is a middleware which verifies an Authorization header JWT using the Firebase Admin SDK.
func authTokenVerifier(auth Authenticator) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, err := auth.VerifyJWT(r.Context(), strings.Split("Bearer ", r.Header.Get("Authorization"))[1])
			if err != nil {
				http.Error(w, "failed to verify authentication token", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r.Clone(context.WithValue(r.Context(), userIDKey, account.ID(userID))))
		})
	}
}

// idFromContext returns the ID stored in ctx, if any.
func idFromContext(ctx context.Context) (account.ID, error) {
	uid, ok := ctx.Value(userIDKey).(account.ID)
	if !ok {
		return "", fmt.Errorf("could not retrieve user ID from context")
	}
	return uid, nil
}

// keyTokenVerifier is a middleware which confirms that the given API key has sufficient permissions to perform the target operation.
func keyTokenVerifier() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract API key from Authorization header (separate method, so stations/ can use it)
			// Use key and station ID to check database for permissions (read/write)
			next.ServeHTTP(w, r)
		})
	}
}

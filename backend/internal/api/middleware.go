package api

import (
	"context"
	"net/http"
	"strings"
)

// Authenticator describes a client which can validate an identity token JWT.
type Authenticator interface {
	VerifyJWT(context.Context, string) (string, error)
}

// authTokenVerifier is a middleware which verifies an Authorization header JWT using the Firebase Admin SDK.
func (s *Server) authTokenVerifier(auth Authenticator) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				s.logger.Print("no authorization header found")
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}

			splitAuthHeader := strings.Split(authHeader, " ")
			if len(splitAuthHeader) != 2 {
				s.logger.Printf("authorzation header value invalid: %s", splitAuthHeader)
				http.Error(w, "improperly formatted authorization header", http.StatusUnauthorized)
				return
			}

			userID, err := auth.VerifyJWT(r.Context(), splitAuthHeader[1])
			if err != nil {
				s.logger.Print(err)
				http.Error(w, "failed to verify authentication token", http.StatusForbidden)
				return
			}

			if userID != getAccountID(r).String() {
				http.Error(w, "not authorized for given account", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// keyTokenVerifier is a middleware which confirms that the given API key has sufficient permissions to perform the target operation.
func (s *Server) keyTokenVerifier() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract API key from Authorization header (separate method, so stations/ can use it)
			// Use key and station ID to check database for permissions (read/write)
			next.ServeHTTP(w, r)
		})
	}
}

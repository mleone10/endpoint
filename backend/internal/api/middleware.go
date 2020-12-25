package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/mleone10/endpoint/internal/account"
	"github.com/mleone10/endpoint/internal/dynamo"
)

// JWTVerifier describes a client which can validate an identity token JWT.
type JWTVerifier interface {
	VerifyJWT(context.Context, string) (string, error)
}

// keyAccountMapper describes a client which can look up the account ID corresponding to a given API Key.
type keyPermissionMapper interface {
	GetKeyPermission(a string) (account.Permission, error)
}

// authTokenVerifier is a middleware which verifies an Authorization header JWT using the Firebase Admin SDK.
func (s *Server) authTokenVerifier(auth JWTVerifier) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := getHeader(r, headerAuthorization)

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

// keyPermissionMapper is a middleware which confirms that the given API key has sufficient permissions to perform the target operation.
func (s *Server) keyPermissionMapper(m keyPermissionMapper) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			a := getHeader(r, headerAPIKey)
			if a == "" {
				http.Error(w, "api key not found", http.StatusUnauthorized)
				return
			}

			p, err := m.GetKeyPermission(a)
			if err != nil && err == dynamo.ErrorItemNotFound {
				http.Error(w, fmt.Sprintf("no account id found for given api key [%s]", a), http.StatusUnauthorized)
				return
			} else if err != nil {
				s.logger.Printf("api key lookup failed: %v", err)
				http.Error(w, "failed to map api key to account id", http.StatusInternalServerError)
				return
			}

			next.ServeHTTP(w, reqWithCtxValue(r, ctxKeyPermission, p))
		})
	}
}

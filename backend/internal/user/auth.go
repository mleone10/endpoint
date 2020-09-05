package user

import (
	"context"
	"fmt"
	"net/http"

	firebase "firebase.google.com/go"
)

// AuthTokenVerifier is a middleware which verifies an Authorization header JWT using the Firebase Admin SDK.
func AuthTokenVerifier() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		app, err := firebase.NewApp(context.Background(), nil)
		if err != nil {
			panic(fmt.Sprintf("error initializing auth middleware: %v", err))
		}
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create a request-scoped authentication client.
			client, err := app.Auth(r.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Verify the token, returning an HTTP 403 if it cannot be.
			verifiedToken, err := client.VerifyIDToken(r.Context(), r.Header.Get("Authorization"))
			if err != nil {
				http.Error(w, "failed to verify authentication token", http.StatusForbidden)
				return
			}

			// Create a new context off of the original request context, but add the user ID from the verified token.
			ctx := NewContextWithID(r.Context(), NewID(verifiedToken.UID))
			next.ServeHTTP(w, r.Clone(ctx))
		})
	}
}

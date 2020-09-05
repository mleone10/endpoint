package middleware

import (
	"context"
	"fmt"
	"net/http"

	firebase "firebase.google.com/go"
)

// AuthTokenVerifier is a middleware which verifies an Authorization header JWT using the Firebase Admin SDK
func AuthTokenVerifier() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		// TODO: Key will have to be pulled from environment variable in lambda
		app, err := firebase.NewApp(context.Background(), nil)
		if err != nil {
			panic(fmt.Sprintf("error initializing auth middleware: %v", err))
		}
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			client, err := app.Auth(r.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			_, err = client.VerifyIDToken(r.Context(), r.Header.Get("Authorization"))
			if err != nil {
				http.Error(w, "failed to verify authentication token", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

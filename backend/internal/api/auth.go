package api

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
)

// Authenticator is a client object which authenticates credentials against a backend service.
type Authenticator struct {
	fbApp *firebase.App
}

// NewAuthenticator returns a fully initialized Authenticator, or an error if it cannot be created.
func NewAuthenticator() (*Authenticator, error) {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error initializing authentication client: %w", err)
	}

	return &Authenticator{
		fbApp: app,
	}, nil
}

// VerifyJWT verifies a give JWT against a backend service and returns a user ID string if successful.  If the operation fails, or if the token cannot be verified, an error is returned.
func (a *Authenticator) VerifyJWT(ctx context.Context, jwt string) (string, error) {
	// Create a request-scoped authentication client.
	client, err := a.fbApp.Auth(ctx)
	if err != nil {
		// TODO: Find a way to differentiate between this error (HTTP 500) and the next one (HTTP 403)
		return "", fmt.Errorf("error verifying authentication token: %w", err)
	}

	// Verify the token, returning an HTTP 403 if it cannot be.
	verifiedToken, err := client.VerifyIDToken(ctx, jwt)
	if err != nil {
		return "", fmt.Errorf("failed to verify authentication token")
	}

	return verifiedToken.UID, nil
}

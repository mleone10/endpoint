package mock

import "context"

// Authenticator is a mock authenticator
type Authenticator struct {
}

// NewAuthenticator returns a new, empty, mock Authentictor.  The second return value, the error, is always nil.
func NewAuthenticator() (*Authenticator, error) {
	return &Authenticator{}, nil
}

// VerifyJWT always returns the same user ID ("testUserID") and a nil error.
func (a *Authenticator) VerifyJWT(ctx context.Context, jwt string) (string, error) {
	return "testUserID", nil
}

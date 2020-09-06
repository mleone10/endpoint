package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/segmentio/ksuid"
)

var errorReadingUserID = errors.New("could not retrieve user ID from context")

// APIKey represents an internal and client-facing API key
type APIKey struct {
	Key      string `json:"key"`
	Nickname string `json:"nickname"`
	ReadOnly bool   `json:"readOnly"`
}

type apiKeyDatastore interface {
	GetAPIKeys(uid *ID) ([]APIKey, error)
	PutAPIKey(uid *ID, apiKey *APIKey) error
}

func getAPIKeys(ctx context.Context, db apiKeyDatastore) ([]APIKey, error) {
	if uid, ok := IDFromContext(ctx); ok {
		return db.GetAPIKeys(uid)
	}
	return nil, errorReadingUserID
}

func newAPIKey(ctx context.Context, db apiKeyDatastore, nickname string, readOnly bool) (*APIKey, error) {
	apiKey := &APIKey{
		Key:      ksuid.New().String(),
		Nickname: nickname,
		ReadOnly: readOnly,
	}

	uid, ok := IDFromContext(ctx)
	if !ok {

		return nil, errorReadingUserID
	}

	err := db.PutAPIKey(uid, apiKey)
	if err != nil {
		return nil, fmt.Errorf("could not save key to database: %w", err)
	}

	return apiKey, nil
}

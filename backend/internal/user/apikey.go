package user

import (
	"context"
	"fmt"

	"github.com/segmentio/ksuid"
)

// APIKey represents an internal and client-facing API key
type APIKey struct {
	Key      string `json:"key"`
	Nickname string `json:"nickname"`
	ReadOnly bool   `json:"readOnly"`
}

type apiKeyDatastore interface {
	GetAPIKeys(uid *ID) ([]APIKey, error)
	PutAPIKey(apiKey *APIKey) error
}

func getAPIKeys(ctx context.Context, db apiKeyDatastore) ([]APIKey, error) {
	if uid, ok := IDFromContext(ctx); ok {
		return db.GetAPIKeys(uid)
	}
	return nil, fmt.Errorf("could not retrieve user ID from context")
}

func newAPIKey(ctx context.Context, db apiKeyDatastore, nickname string, readOnly bool) (*APIKey, error) {
	apiKey := &APIKey{
		Key:      ksuid.New().String(),
		Nickname: nickname,
		ReadOnly: readOnly,
	}

	err := db.PutAPIKey(apiKey)
	if err != nil {
		return nil, fmt.Errorf("could not save key to database: %w", err)
	}

	return apiKey, nil
}

package user

import (
	"context"
	"fmt"
)

// APIKey represents an internal and client-facing API key
type APIKey struct {
	Key      string `json:"key"`
	Nickname string `json:"nickname"`
	ReadOnly bool   `json:"readOnly"`
}

type apiKeyDatastore interface {
	GetAPIKeys(uid *ID) ([]APIKey, error)
}

func getAPIKeys(ctx context.Context, db apiKeyDatastore) ([]APIKey, error) {
	if uid, ok := IDFromContext(ctx); ok {
		return db.GetAPIKeys(uid)
	}
	return nil, fmt.Errorf("could not retrieve user ID from context")
}

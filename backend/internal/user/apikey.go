package user

import (
	"context"
	"fmt"
)

type APIKey string

type apiKeyDatastore interface {
	GetAPIKeys(uid *ID) ([]APIKey, error)
}

func getAPIKeys(ctx context.Context, db apiKeyDatastore) ([]APIKey, error) {
	if uid, ok := IDFromContext(ctx); ok {
		return db.GetAPIKeys(uid)
	}
	return nil, fmt.Errorf("could not retrieve user ID from context")
}

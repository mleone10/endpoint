package user

import "context"

type apiKey string

func getAPIKeys(ctx context.Context) ([]apiKey, error) {
	return []apiKey{apiKey("key1"), apiKey("key2")}, nil
}

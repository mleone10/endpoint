package account

import (
	"github.com/segmentio/ksuid"
)

// Datastore is a persistence layer for Users
type Datastore interface {
	SaveAPIKey(ID, *APIKey) error
	ListAPIKeys(ID) ([]APIKey, error)
}

// APIKey represents an internal and client-facing API key
type APIKey struct {
	Key      string `json:"key"`
	ReadOnly bool   `json:"readOnly"`
}

// NewAPIKey returns an initialized APIKey
func NewAPIKey(readOnly bool) *APIKey {
	return &APIKey{
		Key:      ksuid.New().String(),
		ReadOnly: readOnly,
	}
}

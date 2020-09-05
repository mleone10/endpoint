package user

import "context"

type userIDKeyType string

const userIDKey userIDKeyType = "userID"

// NewContext returns a Context that contains the given ID.
func NewContext(ctx context.Context, id *ID) context.Context {
	return context.WithValue(ctx, userIDKey, id)
}

// FromContext returns the ID stored in ctx, if any.
func FromContext(ctx context.Context) (string, bool) {
	uid, ok := ctx.Value(userIDKey).(string)
	return uid, ok
}

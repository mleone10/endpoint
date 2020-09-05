package user

import "context"

type userIDKeyType string

const userIDKey userIDKeyType = "userID"

// NewContextWithID returns a Context that contains the given ID.
func NewContextWithID(ctx context.Context, id *ID) context.Context {
	return context.WithValue(ctx, userIDKey, id)
}

// IDFromContext returns the ID stored in ctx, if any.
func IDFromContext(ctx context.Context) (string, bool) {
	uid, ok := ctx.Value(userIDKey).(string)
	return uid, ok
}

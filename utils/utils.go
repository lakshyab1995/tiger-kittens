package utils

import (
	"context"

	"github.com/lakshyab1995/tiger-kittens/db"
)

var UserCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *db.User {
	raw, _ := ctx.Value(UserCtxKey).(*db.User)
	return raw
}

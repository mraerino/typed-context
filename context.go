package typedcontext

import (
	"context"
)

type ctxKey interface{}

func Set[T any](ctx context.Context, value T) context.Context {
	key := ctxKey((*T)(nil))
	return context.WithValue(ctx, key, value)
}

func Get[T any](ctx context.Context) (val T, ok bool) {
	key := ctxKey((*T)(nil))
	if ctxVal := ctx.Value(key); ctxVal != nil {
		return ctxVal.(T), true
	}
	return // uses zero value for T
}

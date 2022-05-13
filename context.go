package typedcontext

import (
	"context"
)

type ctxKey interface{}

func Set[T any](ctx context.Context, value T) context.Context {
	var zero T
	key := ctxKey(zero)
	return context.WithValue(ctx, key, value)
}

func Get[T any](ctx context.Context) (val T, ok bool) {
	var zero T
	if ctxVal := ctx.Value(zero); ctxVal != nil {
		return ctxVal.(T), true
	}
	return // uses zero value for T
}

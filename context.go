package typedcontext

import (
	"context"
	"reflect"
)

type ctxKey reflect.Type

func Set[T any](ctx context.Context, value T) context.Context {
	t := reflect.TypeOf(value)
	key := ctxKey(t)
	return context.WithValue(ctx, key, value)
}

func Get[T any](ctx context.Context) (val T, ok bool) {
	t := reflect.TypeOf(val)
	key := ctxKey(t)
	if ctxVal := ctx.Value(key); ctxVal != nil {
		return ctxVal.(T), true
	}
	return // uses zero value for T
}

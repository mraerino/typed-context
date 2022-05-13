package typedcontext_test

import (
	"context"
	"testing"

	typedcontext "github.com/mraerino/typed-context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Welcome string
type RequestID string

func TestGetSet(t *testing.T) {
	ctx := context.Background()

	value := Welcome("Hello World")
	ctx = typedcontext.Set(ctx, value)

	actual, ok := typedcontext.Get[Welcome](ctx)
	require.True(t, ok)
	assert.Equal(t, "Hello World", string(actual))

	// override
	value2 := Welcome("Goodbye")
	ctx = typedcontext.Set(ctx, value2)

	actual2, ok := typedcontext.Get[Welcome](ctx)
	require.True(t, ok)
	assert.Equal(t, "Goodbye", string(actual2))

	// different type
	value3 := RequestID("0c4f7d51-af18-4475-9fdc-5f022fb8079c")
	ctx = typedcontext.Set(ctx, value3)

	actual3, ok := typedcontext.Get[Welcome](ctx)
	require.True(t, ok)
	assert.Equal(t, "Goodbye", string(actual3))

	actual4, ok := typedcontext.Get[RequestID](ctx)
	require.True(t, ok)
	assert.Equal(t, "0c4f7d51-af18-4475-9fdc-5f022fb8079c", string(actual4))
}

type ctxKey uint64

const (
	requestIDKey ctxKey = iota
)

var valStdlib any

func BenchmarkStdlib(b *testing.B) {
	ctx := context.Background()

	ctx = context.WithValue(ctx, requestIDKey, "0c4f7d51-af18-4475-9fdc-5f022fb8079c")

	var val any
	for n := 0; n < b.N; n++ {
		val = ctx.Value(requestIDKey)
		if val == nil {
			b.Fatal("not found")
		}
		if val.(string) != "0c4f7d51-af18-4475-9fdc-5f022fb8079c" {
			b.Fatal("wrong value")
		}
	}
	valStdlib = val
}

var valTyped RequestID

func BenchmarkTyped(b *testing.B) {
	ctx := context.Background()

	ctx = typedcontext.Set(ctx, RequestID("0c4f7d51-af18-4475-9fdc-5f022fb8079c"))

	var val RequestID
	var ok bool
	for n := 0; n < b.N; n++ {
		val, ok = typedcontext.Get[RequestID](ctx)
		if !ok {
			b.Fatal("not found")
		}
		if val != "0c4f7d51-af18-4475-9fdc-5f022fb8079c" {
			b.Fatal("wrong value")
		}
	}
	valTyped = val
}

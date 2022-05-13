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

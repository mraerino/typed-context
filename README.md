# Typed Context Access for Go

Using generics to access values in a `context.Context`.

This makes sure that you never need to cast the values you get from `ctx.Value` at runtime.

## Example

```go
package main

import (
	"context"
	"fmt"

	typedcontext "github.com/mraerino/typed-context"
)

type RequestID string

func main() {
	ctx := context.Background()

	// set a value
	ctx = typedcontext.Set(ctx, RequestID("12345"))

	// get a value
	requestID, ok := typedcontext.Get[RequestID](ctx)
	if !ok {
		panic("Not found!")
	}
	fmt.Printf("Request ID: %s\n", requestID)
}
```

[Run in the Playground](https://go.dev/play/p/yIb_mnvQQjS)
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

## Performance

In a micro-benchmark, this library performs about 50% slower as manually using the stdlib.

```
goos: darwin
goarch: amd64
pkg: github.com/mraerino/typed-context
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkStdlib-16      205549032                5.782 ns/op           0 B/op          0 allocs/op
BenchmarkTyped-16       123351498                8.951 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/mraerino/typed-context       3.977s
```
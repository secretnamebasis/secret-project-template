package hello_world

import (
	"context"

	"github.com/deroproject/derohe/cmd/derod/rpc"
)

func Hello() string {
	args := []string{"Hello, World!"}
	ctx := context.Background()
	deroHello := rpc.Echo(ctx, args)
	return deroHello[5:] // trim `DERO ` from the echo
}

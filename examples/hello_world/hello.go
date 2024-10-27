package hello_world

import (
	"context"

	"github.com/deroproject/derohe/cmd/derod/rpc"
)

func Hello() string {
	return rpc.Echo(
		context.Background(),
		[]string{"Hello, World!"},
	)[5:] // trim the `DERO ` from the echo
}

package hello_world_test

import (
	"os"
	"project-template/examples/hello_world"
	"testing"

	"github.com/deroproject/derohe/blockchain"
	"github.com/deroproject/derohe/cmd/derod/rpc"
	"github.com/deroproject/derohe/globals"
	"github.com/deroproject/derohe/p2p"
)

type args map[string]any

var arg_flags args = args{
	"--rpc-bind":  "http://127.0.0.1:42069",
	"--simulator": true,
	"--testnet":   true,
	"--p2p-bind":  ":0",
}

func TestHello(t *testing.T) {
	var err error
	func() {

		// implement simulated DERO serve
		globals.Arguments = arg_flags

		if err := os.RemoveAll(globals.GetDataDirectory()); err != nil {
			t.Error(err)
		} // remove oldirectory

		globals.Initialize() // setup network and proxy

		params := args{
			"--simulator": true,
		}

		params["chain"], err = blockchain.Blockchain_Start(params) //start chain in simulator mode

		if err != nil {
			t.Error(err)
		}

		if err := p2p.P2P_Init(params); err != nil {
			t.Error(err)
		}

		helloworldServer, err := rpc.RPCServer_Start(params)
		if err != nil {
			t.Error(err)
		}
		defer helloworldServer.RPCServer_Stop()
	}()
	expect := "Hello, World!"
	observed := hello_world.Hello()
	if observed != expect {
		t.Errorf("expected: %s, observed: %s,", expect, observed)
	}
}

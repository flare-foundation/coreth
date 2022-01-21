// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"fmt"
	"github.com/flare-foundation/flare/ids"
	"os"

	"github.com/hashicorp/go-plugin"

	"github.com/flare-foundation/coreth/plugin/evm"
	"github.com/flare-foundation/flare/vms/rpcchainvm"
)

func main() {
	version, err := PrintVersion()
	if err != nil {
		fmt.Printf("couldn't get config: %s", err)
		os.Exit(1)
	}
	if version {
		fmt.Println(evm.Version)
		os.Exit(0)
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: rpcchainvm.Handshake,
		Plugins: map[string]plugin.Plugin{
			"vm":         rpcchainvm.New(&evm.VM{}),
			"validators": rpcchainvm.NewPluginValidator(&KV2{}), // TODO add a separate component here which only does what we need
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}

// Here is a real implementation of KV that writes to a local file with
// the key name and the contents are the value of the key.
type KV2 struct{}

func (KV2) GetValidators(id ids.ID) (map[ids.ShortID]float64, error) {
	fmt.Println("Real implementation of GetValidators called")
	m := make(map[ids.ShortID]float64)

	shortID := [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	m[shortID] = 2.3
	return m, nil
}

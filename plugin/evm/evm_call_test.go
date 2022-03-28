//go:build integration
// +build integration

package evm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/coreth/plugin/evm/testcontracts"
)

func TestEVMCall_Execute(t *testing.T) {

	auth, be := simulatedBlockchain(t)
	defer be.Close()

	contractAddr := testcontracts.DeployStore(auth, be)

	be.Commit(true)

	contract := EVMContract{
		address: contractAddr,
		abi:     testcontracts.StoreABI(),
	}

	blockChain := be.Blockchain()

	evmBind := BindEVM(blockChain)

	t.Run("nominal case", func(t *testing.T) {
		header := blockChain.LastAcceptedBlock().Header()

		ret := evmBind.AtBlock(header.Hash()).OnContract(contract).Execute("setItem", [32]byte{1}, [32]byte{1})
		require.NoError(t, ret.err)

		ret = evmBind.AtBlock(header.Hash()).OnContract(contract).Execute("items", [32]byte{1})
		require.NoError(t, ret.err)
		require.Len(t, ret.values, 1)
	})

	t.Run("handles invalid method", func(t *testing.T) {
		header := blockChain.LastAcceptedBlock().Header()
		ret := evmBind.AtBlock(header.Hash()).OnContract(contract).Execute("setItem2", [32]byte{1}, [32]byte{1})
		assert.EqualError(t, ret.err, "could not pack parameters: method 'setItem2' not found")
	})

	t.Run("handles invalid parameters", func(t *testing.T) {
		header := blockChain.LastAcceptedBlock().Header()
		ret := evmBind.AtBlock(header.Hash()).OnContract(contract).Execute("setItem", [32]byte{1}, [32]byte{1}, [32]byte{1})
		assert.EqualError(t, ret.err, "could not pack parameters: argument count mismatch: got 3 for 2")
	})

	t.Run("handles invalid header number", func(t *testing.T) {
		header := blockChain.GetHeaderByNumber(2000)
		hash := header.Hash()
		ret := evmBind.AtBlock(hash).OnContract(contract).Execute("setItem", [32]byte{1}, [32]byte{1}, [32]byte{1})
		assert.EqualError(t, ret.err, fmt.Sprintf("unknown block (hash: %x)", hash))
	})
}

//go:build integration
// +build integration

package ftso

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/coreth/plugin/evm/ftso/testcontracts"
)

func Test_contractCall_execute(t *testing.T) {
	auth, be := simulatedBackend(t)
	defer be.Close()

	contractAddr := testcontracts.DeployStore(auth, be)

	contract := evmContract{
		address: contractAddr,
		abi:     testcontracts.StoreABI(),
	}

	be.Commit(true)

	blockchain := be.Blockchain()

	evm := testEVM(t, blockchain)

	t.Run("nominal case", func(t *testing.T) {
		res := newContractCall(evm, contract).execute("setItem", [32]byte{10}, [32]byte{2})
		require.NoError(t, res.err)

		resInt := big.NewInt(0)
		err := res.decode(&resInt)
		require.NoError(t, err)
		assert.Equal(t, uint64(1), resInt.Uint64())

		res = newContractCall(evm, contract).execute("items", [32]byte{10})
		require.NoError(t, res.err)
		require.Len(t, res.values, 1)
		var resItem [32]byte
		err = res.decode(&resItem)
		require.NoError(t, err)
		assert.Equal(t, [32]byte{2}, resItem)
	})

	t.Run("handles invalid method", func(t *testing.T) {
		res := newContractCall(evm, contract).execute("setItem2", [32]byte{1}, [32]byte{1})
		assert.EqualError(t, res.err, "could not pack parameters: method 'setItem2' not found")
	})

	t.Run("handles invalid parameters", func(t *testing.T) {
		res := newContractCall(evm, contract).execute("setItem", [32]byte{1}, [32]byte{1}, [32]byte{1})
		assert.EqualError(t, res.err, "could not pack parameters: argument count mismatch: got 3 for 2")
	})

	t.Run("handles invalid len of array error", func(t *testing.T) {
		res := newContractCall(evm, contract).execute("items", [32]byte{10})
		require.NoError(t, res.err)
		require.Len(t, res.values, 1)
		var resItem [40]byte
		err := res.decode(&resItem)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid len for return array")
	})

	t.Run("handles invalid type for return value error", func(t *testing.T) {
		res := newContractCall(evm, contract).execute("items", [32]byte{10})
		require.NoError(t, res.err)
		require.Len(t, res.values, 1)
		var resItem int
		err := res.decode(&resItem)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid type for return value")
	})

	t.Run("handles invalid number of decode values error", func(t *testing.T) {
		res := newContractCall(evm, contract).execute("items", [32]byte{10})
		require.NoError(t, res.err)
		require.Len(t, res.values, 1)
		var resItem int
		err := res.decode(&resItem, &resItem)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid number of decode values")
	})
}

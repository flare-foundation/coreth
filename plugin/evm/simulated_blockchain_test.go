//go:build integration
// +build integration

package evm

import (
	"math"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/coreth/accounts/abi/bind"
	"github.com/flare-foundation/coreth/accounts/abi/bind/backends"
	"github.com/flare-foundation/coreth/core"
)

func simulatedBlockchain(t *testing.T) (*bind.TransactOpts, *backends.SimulatedBackend) {
	t.Helper()

	key, err := crypto.GenerateKey()
	require.NoError(t, err)

	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	require.NoError(t, err)

	balance := big.NewInt(0).SetUint64(math.MaxUint64)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{
		Balance: balance,
	}
	gasLimit := uint64(math.MaxUint64)

	be := backends.NewSimulatedBackend(alloc, gasLimit)

	return auth, be
}

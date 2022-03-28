//go:build integration
// +build integration

package evm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFTSOSnapshot_Cap(t *testing.T) {
	auth, be := simulatedBlockchain(t)
	defer be.Close()

	cfg := defaultTestContractsConfig
	withRewardEpochPowerHeights([]*big.Int{big.NewInt(2), big.NewInt(1)})(&cfg)
	withRewardEpochStartHeights([]*big.Int{big.NewInt(3), big.NewInt(4)})(&cfg)

	t.Run("nominal case", func(t *testing.T) {
		validationAddr, submitterAddr := deployAllContracts(t, auth, be, cfg)

		// to get current header to be '4'
		for i := 0; i < 4; i++ {
			be.Commit(true)
		}

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		snapshot, err := ftsoSystem.Snapshot(cfg.epochs[0].Uint64())
		require.NoError(t, err)

		actualCap, err := snapshot.Cap()
		require.NoError(t, err)

		capInt := big.NewInt(0).Div(cfg.totalSupply, cfg.fraction)
		capFloat := big.NewFloat(0).SetInt(capInt)
		expCap, _ := capFloat.Float64()

		assert.Equal(t, expCap, actualCap)
	})

	t.Run("handles could not get total supply error", func(t *testing.T) {
		validationAddr, submitterAddr := deployAllContracts(t, auth, be, cfg)

		// to get current header to be '4'
		for i := 0; i < 4; i++ {
			be.Commit(true)
		}

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		snapshot, err := ftsoSystem.Snapshot(cfg.epochs[0].Uint64())
		require.NoError(t, err)

		someHash := common.BigToHash(big.NewInt(1))
		snapshot.(*FTSOSnapshot).start = someHash

		_, err = snapshot.Cap()
		assert.EqualError(t, err, fmt.Sprintf("could not get total supply: unknown block (hash: %x)", someHash))
	})
}

func TestFTSOSnapshot_Providers(t *testing.T) {

}

func TestFTSOSnapshot_Rewards(t *testing.T) {

}

func TestFTSOSnapshot_Validator(t *testing.T) {

}

func TestFTSOSnapshot_Votepower(t *testing.T) {

}

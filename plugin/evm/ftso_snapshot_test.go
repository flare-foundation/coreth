//go:build integration
// +build integration

package evm

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/coreth/accounts/abi/bind"
	"github.com/flare-foundation/coreth/accounts/abi/bind/backends"
	"github.com/flare-foundation/coreth/plugin/evm/testcontracts"
	"github.com/flare-foundation/flare/ids"
)

func TestFTSOSnapshot_Cap(t *testing.T) {
	cfg := defaultTestContractsConfig
	withRewardEpochPowerHeights([]*big.Int{big.NewInt(2), big.NewInt(1)})(&cfg)
	withRewardEpochStartHeights([]*big.Int{big.NewInt(3), big.NewInt(4)})(&cfg)

	t.Run("nominal case", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		snapshot := testFTSOSnapshot(t, auth, be, cfg)

		actualCap, err := snapshot.Cap()
		require.NoError(t, err)

		capInt := big.NewInt(0).Div(cfg.totalSupply, cfg.fraction)
		capFloat := big.NewFloat(0).SetInt(capInt)
		expCap, _ := capFloat.Float64()

		assert.Equal(t, expCap, actualCap)
	})

	t.Run("handles could not get total supply error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		snapshot := testFTSOSnapshot(t, auth, be, cfg)

		someHash := common.BigToHash(big.NewInt(1))
		snapshot.(*FTSOSnapshot).start = someHash

		_, err := snapshot.Cap()
		assert.EqualError(t, err, fmt.Sprintf("could not get total supply: unknown block (hash: %x)", someHash))
	})
}

func TestFTSOSnapshot_Providers(t *testing.T) {
	cfg := defaultTestContractsConfig
	withRewardEpochPowerHeights([]*big.Int{big.NewInt(2), big.NewInt(1)})(&cfg)
	withRewardEpochStartHeights([]*big.Int{big.NewInt(3), big.NewInt(4)})(&cfg)

	t.Run("nominal case", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		snapshot := testFTSOSnapshot(t, auth, be, cfg)

		addresses, err := snapshot.Providers()
		require.NoError(t, err)
		require.Len(t, addresses, len(cfg.providers))
		assert.Equal(t, addresses[0], cfg.providers[0])
		assert.Equal(t, addresses[1], cfg.providers[1])
	})

	t.Run("handles no providers in the result set", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		votePowerAddr := testcontracts.DeployVotepower(auth, be, cfg.providers, cfg.votePowers)

		wnatAddr := testcontracts.DeployWNAT(auth, be, votePowerAddr, cfg.totalSupply)

		rewardAddr := testcontracts.DeployReward(auth, be, wnatAddr, cfg.epochs, cfg.providers, cfg.unclaimedRewards)

		ftsoManagerAddr := testcontracts.DeployManager(auth, be, rewardAddr, cfg.rewardEpochDurationSeconds, cfg.rewardEpochsStartTs,
			cfg.currentRewardEpoch, cfg.fraction, cfg.epochs, cfg.rewardEpochPowerHeights, cfg.rewardEpochStartHeights, cfg.rewardEpochStartTimes)

		ftsoRegistryAddr := testcontracts.DeployFTSORegistry(auth, be, cfg.supportedIndices)

		indices := []*big.Int{big.NewInt(88888), big.NewInt(99999)}

		voterWhitelisterAddr := testcontracts.DeployWhitelist(auth, be, indices, cfg.providers)

		validationAddr := testcontracts.DeployValidator(auth, be, cfg.providers, cfg.nodes)

		submitterAddr := testcontracts.DeploySubmitter(auth, be, voterWhitelisterAddr, ftsoRegistryAddr, ftsoManagerAddr)

		// to get current header to be '4'
		for i := 0; i < 4; i++ {
			be.Commit(true)
		}

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		snapshot, err := ftsoSystem.Snapshot(cfg.epochs[0].Uint64())
		require.NoError(t, err)

		addresses, err := snapshot.Providers()
		require.NoError(t, err)
		require.Len(t, addresses, 0)
	})

	t.Run("handles could not get series indices error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		snapshot := testFTSOSnapshot(t, auth, be, cfg)

		snapshot.(*FTSOSnapshot).contracts.Registry = EVMContract{}

		_, err := snapshot.Providers()
		assert.EqualError(t, err, fmt.Sprintf("could not get series indices: could not pack parameters: method '%s' not found", SeriesIndices))
	})

	t.Run("handles could not get provider addresses error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		snapshot := testFTSOSnapshot(t, auth, be, cfg)

		snapshot.(*FTSOSnapshot).contracts.Whitelist = EVMContract{}

		_, err := snapshot.Providers()
		assert.EqualError(t, err, fmt.Sprintf("could not get provider addresses (index: %d): could not pack parameters: method '%s' not found", cfg.supportedIndices[0].Uint64(), DataProviders))
	})
}

func TestFTSOSnapshot_Rewards(t *testing.T) {
	cfg := defaultTestContractsConfig
	withEpochs([]*big.Int{big.NewInt(1), big.NewInt(2)})
	withRewardEpochPowerHeights([]*big.Int{big.NewInt(2), big.NewInt(1)})(&cfg)
	withRewardEpochStartHeights([]*big.Int{big.NewInt(3), big.NewInt(4)})(&cfg)

	t.Run("nominal case", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		snapshot := testFTSOSnapshot(t, auth, be, cfg)

		got, err := snapshot.Rewards(cfg.providers[0])
		require.NoError(t, err)

		rwFloat := big.NewFloat(0).SetInt(cfg.unclaimedRewards[0])
		expected, _ := rwFloat.Float64()

		assert.Equal(t, expected, got)
	})

	t.Run("handles could not get provider rewards error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		snapshot := testFTSOSnapshot(t, auth, be, cfg)

		snapshot.(*FTSOSnapshot).contracts.Rewards = EVMContract{}

		_, err := snapshot.Rewards(cfg.providers[0])
		assert.EqualError(t, err, fmt.Sprintf("could not get provider rewards: could not pack parameters: method '%s' not found", ProviderRewards))
	})
}

func TestFTSOSnapshot_Validator(t *testing.T) {
	cfg := defaultTestContractsConfig
	withRewardEpochPowerHeights([]*big.Int{big.NewInt(2), big.NewInt(1)})(&cfg)
	withRewardEpochStartHeights([]*big.Int{big.NewInt(3), big.NewInt(4)})(&cfg)

	t.Run("nominal case", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		snapshot := testFTSOSnapshot(t, auth, be, cfg)

		got, err := snapshot.Validator(cfg.providers[0])
		require.NoError(t, err)
		assert.Equal(t, ids.ShortID(cfg.nodes[0]), got)
	})

	t.Run("handles could not get provider node error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		snapshot := testFTSOSnapshot(t, auth, be, cfg)

		snapshot.(*FTSOSnapshot).contracts.Validation = EVMContract{}

		_, err := snapshot.Validator(cfg.providers[0])
		assert.EqualError(t, err, fmt.Sprintf("could not get provider node: could not pack parameters: method '%s' not found", ProviderNode))
	})
}

func TestFTSOSnapshot_Votepower(t *testing.T) {
	cfg := defaultTestContractsConfig
	withRewardEpochPowerHeights([]*big.Int{big.NewInt(2), big.NewInt(1)})(&cfg)
	withRewardEpochStartHeights([]*big.Int{big.NewInt(3), big.NewInt(4)})(&cfg)

	t.Run("nominal case", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		snapshot := testFTSOSnapshot(t, auth, be, cfg)

		got, err := snapshot.Votepower(cfg.providers[0])
		require.NoError(t, err)

		vpFloat := big.NewFloat(0).SetInt(cfg.votePowers[0])
		votepower, _ := vpFloat.Float64()

		assert.Equal(t, votepower, got)
	})

	t.Run("handles could not get provider node error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		snapshot := testFTSOSnapshot(t, auth, be, cfg)

		snapshot.(*FTSOSnapshot).contracts.Votepower = EVMContract{}

		_, err := snapshot.Votepower(cfg.providers[0])
		assert.EqualError(t, err, fmt.Sprintf("could not get provider votepower: could not pack parameters: method '%s' not found", ProviderVotepower))
	})
}

func testFTSOSnapshot(t *testing.T, auth *bind.TransactOpts, be *backends.SimulatedBackend, cfg testContractsConfig) Snapshot {
	t.Helper()

	validationAddr, submitterAddr := deployAllContracts(t, auth, be, cfg)

	// to get current header to be '4'
	for i := 0; i < 4; i++ {
		be.Commit(true)
	}

	ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

	snapshot, err := ftsoSystem.Snapshot(cfg.epochs[0].Uint64())
	require.NoError(t, err)

	return snapshot
}

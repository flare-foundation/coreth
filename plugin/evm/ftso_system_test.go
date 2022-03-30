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
)

func TestNewFTSOSystem(t *testing.T) {
	_, be := simulatedBlockchain(t)
	defer be.Close()

	addressSubmitter := common.BigToAddress(big.NewInt(1))
	addressValidation := common.BigToAddress(big.NewInt(2))

	ftsoSystem, err := NewFTSOSystem(be.Blockchain(), addressSubmitter, addressValidation)
	require.NoError(t, err)
	assert.Equal(t, addressSubmitter, ftsoSystem.submitter.address)
	assert.Equal(t, addressValidation, ftsoSystem.validation.address)
}

func TestFTSOSystem_Contracts(t *testing.T) {

	t.Run("nominal case", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		cfg := defaultTestContractsConfig

		votePowerAddr := testcontracts.DeployVotepower(auth, be, cfg.providers, cfg.votePowers)

		wnatAddr := testcontracts.DeployWNAT(auth, be, votePowerAddr, cfg.totalSupply)

		rewardAddr := testcontracts.DeployReward(auth, be, wnatAddr, cfg.epochs, cfg.providers, cfg.unclaimedRewards)

		ftsoManagerAddr := testcontracts.DeployManager(auth, be, rewardAddr, cfg.rewardEpochDurationSeconds, cfg.rewardEpochsStartTs,
			cfg.currentRewardEpoch, cfg.fraction, cfg.epochs, cfg.rewardEpochPowerHeights, cfg.rewardEpochStartHeights, cfg.rewardEpochStartTimes)

		ftsoRegistryAddr := testcontracts.DeployFTSORegistry(auth, be, cfg.supportedIndices)

		voterWhitelisterAddr := testcontracts.DeployWhitelist(auth, be, cfg.supportedIndices, cfg.providers)

		validationAddr := testcontracts.DeployValidator(auth, be, cfg.providers, cfg.nodes)

		submitterAddr := testcontracts.DeploySubmitter(auth, be, voterWhitelisterAddr, ftsoRegistryAddr, ftsoManagerAddr)

		be.Commit(true)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		latestBlock := be.Blockchain().LastAcceptedBlock()

		contracts, err := ftsoSystem.Contracts(latestBlock.Hash())
		require.NoError(t, err)

		assert.Equal(t, ftsoRegistryAddr, contracts.Registry.address)
		assert.Equal(t, ftsoManagerAddr, contracts.Manager.address)
		assert.Equal(t, rewardAddr, contracts.Rewards.address)
		assert.Equal(t, voterWhitelisterAddr, contracts.Whitelist.address)
		assert.Equal(t, votePowerAddr, contracts.Votepower.address)
		assert.Equal(t, validationAddr, contracts.Validation.address)
	})

	t.Run("handles could not get manager address error", func(t *testing.T) {
		_, be := simulatedBlockchain(t)
		defer be.Close()

		be.Commit(true)

		ftsoSystem := testFTSOSystem(t, be, common.Address{}, common.Address{})

		latestBlock := be.Blockchain().LastAcceptedBlock()

		_, err := ftsoSystem.Contracts(latestBlock.Hash())
		assert.EqualError(t, err, "could not get manager address: no return data")
	})

	t.Run("handles FTSO not testcontracts.Deployed error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		// Submitter contract
		submitterAddr := testcontracts.DeploySubmitter(auth, be, common.Address{}, common.Address{}, common.Address{})

		be.Commit(true)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, common.Address{})

		latestBlock := be.Blockchain().LastAcceptedBlock()

		_, err := ftsoSystem.Contracts(latestBlock.Hash())
		assert.ErrorIs(t, err, errFTSONotDeployed)
	})

	t.Run("handles FTSO not active error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		cfg := defaultTestContractsConfig
		withEpochs([]*big.Int{big.NewInt(0)})(&cfg)
		withRewardEpochStartHeights([]*big.Int{big.NewInt(0)})(&cfg)

		ftsoManagerAddr := testcontracts.DeployManager(auth, be, common.Address{}, big.NewInt(1), big.NewInt(1), big.NewInt(1),
			cfg.fraction, cfg.epochs, []*big.Int{big.NewInt(100)}, cfg.rewardEpochStartHeights, []*big.Int{big.NewInt(100)})

		submitterAddr := testcontracts.DeploySubmitter(auth, be, common.Address{}, common.Address{}, ftsoManagerAddr)

		be.Commit(true)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, common.Address{})

		latestBlock := be.Blockchain().LastAcceptedBlock()

		_, err := ftsoSystem.Contracts(latestBlock.Hash())
		assert.ErrorIs(t, err, errFTSONotActive)
	})

	t.Run("handles could not get WNAT address error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		cfg := defaultTestContractsConfig

		rewardsAddr := common.Address{}

		ftsoManagerAddr := testcontracts.DeployManager(auth, be, rewardsAddr, cfg.rewardEpochDurationSeconds, cfg.rewardEpochsStartTs,
			cfg.currentRewardEpoch, cfg.fraction, cfg.epochs, cfg.rewardEpochPowerHeights, cfg.rewardEpochStartHeights, cfg.rewardEpochStartTimes)

		ftsoRegistryAddr := testcontracts.DeployFTSORegistry(auth, be, cfg.supportedIndices)

		voterWhitelisterAddr := testcontracts.DeployWhitelist(auth, be, cfg.supportedIndices, cfg.providers)

		validationAddr := testcontracts.DeployValidator(auth, be, cfg.providers, cfg.nodes)

		submitterAddr := testcontracts.DeploySubmitter(auth, be, voterWhitelisterAddr, ftsoRegistryAddr, ftsoManagerAddr)

		be.Commit(true)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		latestBlock := be.Blockchain().LastAcceptedBlock()

		_, err := ftsoSystem.Contracts(latestBlock.Hash())
		assert.EqualError(t, err, "could not get WNAT address: no return data")
	})

	t.Run("handles could not get votepower address error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		cfg := defaultTestContractsConfig

		wnatAddr := common.Address{}

		rewardAddr := testcontracts.DeployReward(auth, be, wnatAddr, cfg.epochs, cfg.providers, cfg.unclaimedRewards)

		ftsoManagerAddr := testcontracts.DeployManager(auth, be, rewardAddr, cfg.rewardEpochDurationSeconds, cfg.rewardEpochsStartTs,
			cfg.currentRewardEpoch, cfg.fraction, cfg.epochs, cfg.rewardEpochPowerHeights, cfg.rewardEpochStartHeights, cfg.rewardEpochStartTimes)

		ftsoRegistryAddr := testcontracts.DeployFTSORegistry(auth, be, cfg.supportedIndices)

		voterWhitelisterAddr := testcontracts.DeployWhitelist(auth, be, cfg.supportedIndices, cfg.providers)

		validationAddr := testcontracts.DeployValidator(auth, be, cfg.providers, cfg.nodes)

		submitterAddr := testcontracts.DeploySubmitter(auth, be, voterWhitelisterAddr, ftsoRegistryAddr, ftsoManagerAddr)

		be.Commit(true)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		latestBlock := be.Blockchain().LastAcceptedBlock()

		_, err := ftsoSystem.Contracts(latestBlock.Hash())

		assert.EqualError(t, err, "could not get votepower address: no return data")
	})
}

func TestFTSOSystem_Details(t *testing.T) {

	t.Run("nominal case", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		cfg := defaultTestContractsConfig

		validationAddr, submitterAddr := deployAllContracts(t, auth, be, cfg)

		be.Commit(true)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		ftsoEpoch, err := ftsoSystem.Details(0)
		require.NoError(t, err)

		assert.Equal(t, cfg.rewardEpochPowerHeights[0].Uint64(), ftsoEpoch.PowerHeight)
		assert.Equal(t, cfg.rewardEpochStartHeights[0].Uint64(), ftsoEpoch.StartHeight)
		assert.Equal(t, cfg.rewardEpochStartTimes[0].Uint64(), ftsoEpoch.StartTime)
	})

	t.Run("handles no current header error", func(t *testing.T) {
		_, be := simulatedBlockchain(t)
		defer be.Close()

		ftsoSystem := testFTSOSystem(t, be, common.Address{}, common.Address{})

		_, err := ftsoSystem.Details(0)
		assert.EqualError(t, err, fmt.Sprintf("could not get contracts (hash: %x): could not get manager address: no return data", ftsoSystem.blockchain.CurrentHeader().Hash()))
	})

}

func TestFTSOSystem_Snapshot(t *testing.T) {

	t.Run("nominal case", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		cfg := defaultTestContractsConfig
		withRewardEpochPowerHeights([]*big.Int{big.NewInt(2), big.NewInt(1)})(&cfg)
		withRewardEpochStartHeights([]*big.Int{big.NewInt(3), big.NewInt(4)})(&cfg)

		validationAddr, submitterAddr := deployAllContracts(t, auth, be, cfg)

		// to get current header to be '4'
		for i := 0; i < 4; i++ {
			be.Commit(true)
		}

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		_, err := ftsoSystem.Snapshot(cfg.epochs[0].Uint64())
		require.NoError(t, err)
	})

	t.Run("handles could not get current epoch details error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		cfg := defaultTestContractsConfig

		validationAddr, submitterAddr := deployAllContracts(t, auth, be, cfg)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		_, err := ftsoSystem.Snapshot(100)
		assert.EqualError(t, err, fmt.Sprintf("could not get current epoch details: could not get contracts (hash: %x): could not get manager address: no return data", ftsoSystem.blockchain.CurrentHeader().Hash()))
	})

	t.Run("handles unknown power block error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		cfg := defaultTestContractsConfig

		powerBlock := big.NewInt(200)
		withRewardEpochPowerHeights([]*big.Int{powerBlock, big.NewInt(1)})(&cfg)

		validationAddr, submitterAddr := deployAllContracts(t, auth, be, cfg)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		// to get current header to be '4'
		for i := 0; i < 4; i++ {
			be.Commit(true)
		}

		_, err := ftsoSystem.Snapshot(cfg.epochs[0].Uint64())
		assert.EqualError(t, err, fmt.Sprintf("unknown power block (height: %d)", powerBlock.Uint64()))
	})

	t.Run("handles unknown current block error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		cfg := defaultTestContractsConfig

		currentBlock := big.NewInt(200)
		withRewardEpochStartHeights([]*big.Int{currentBlock, big.NewInt(1)})(&cfg)
		withRewardEpochPowerHeights([]*big.Int{big.NewInt(2), big.NewInt(1)})(&cfg)

		validationAddr, submitterAddr := deployAllContracts(t, auth, be, cfg)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		// to get current header to be '4'
		for i := 0; i < 4; i++ {
			be.Commit(true)
		}

		_, err := ftsoSystem.Snapshot(cfg.epochs[0].Uint64())
		assert.EqualError(t, err, fmt.Sprintf("unknown current block (height: %d)", currentBlock.Uint64()))
	})

	t.Run("handles could not get next epoch details error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		cfg := defaultTestContractsConfig

		startHeight := big.NewInt(400)

		withEpochs([]*big.Int{big.NewInt(0), big.NewInt(1)})(&cfg)
		withRewardEpochPowerHeights([]*big.Int{big.NewInt(2), big.NewInt(1)})(&cfg)
		withRewardEpochStartHeights([]*big.Int{big.NewInt(3), startHeight})(&cfg)

		validationAddr, submitterAddr := deployAllContracts(t, auth, be, cfg)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		// to get current header to be '4'
		for i := 0; i < 4; i++ {
			be.Commit(true)
		}

		_, err := ftsoSystem.Snapshot(cfg.epochs[0].Uint64())
		assert.EqualError(t, err, fmt.Sprintf("unknown next block (height: %d)", startHeight.Uint64()))
	})

}

func TestFTSOSystem_Current(t *testing.T) {

	t.Run("nominal case", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		cfg := defaultTestContractsConfig

		validationAddr, submitterAddr := deployAllContracts(t, auth, be, cfg)

		be.Commit(true)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		latestBlock := be.Blockchain().LastAcceptedBlock()

		epoch, err := ftsoSystem.Current(latestBlock.Hash())
		require.NoError(t, err)

		assert.Equal(t, cfg.currentRewardEpoch.Uint64(), epoch)
	})

	t.Run("handles could not get contracts error", func(t *testing.T) {
		_, be := simulatedBlockchain(t)
		defer be.Close()

		be.Commit(true)

		ftsoSystem := testFTSOSystem(t, be, common.Address{}, common.Address{})

		latestBlock := be.Blockchain().LastAcceptedBlock()

		_, err := ftsoSystem.Current(latestBlock.Hash())
		assert.EqualError(t, err, "could not get contracts: could not get manager address: no return data")
	})
}

func testFTSOSystem(t *testing.T, be *backends.SimulatedBackend, submitterAddr, validationAddr common.Address) *FTSOSystem {
	t.Helper()

	submitter := EVMContract{
		address: submitterAddr,
		abi:     testcontracts.SubmitterABI(),
	}

	validation := EVMContract{
		address: validationAddr,
		abi:     testcontracts.ValidatorABI(),
	}

	abis := FTSOABIs{
		Registry:  testcontracts.FTSORegistryABI(),
		Manager:   testcontracts.ManagerABI(),
		Rewards:   testcontracts.RewardABI(),
		WNAT:      testcontracts.WnatABI(),
		Whitelist: testcontracts.WhitelistABI(),
		Votepower: testcontracts.VotepowerABI(),
	}

	f := FTSOSystem{
		blockchain: be.Blockchain(),
		submitter:  submitter,
		validation: validation,
		abis:       abis,
	}

	return &f
}

func deployAllContracts(t *testing.T, auth *bind.TransactOpts, be *backends.SimulatedBackend,
	cfg testContractsConfig) (validationAddr, submitterAddr common.Address) {

	t.Helper()

	votePowerAddr := testcontracts.DeployVotepower(auth, be, cfg.providers, cfg.votePowers)

	wnatAddr := testcontracts.DeployWNAT(auth, be, votePowerAddr, cfg.totalSupply)

	rewardAddr := testcontracts.DeployReward(auth, be, wnatAddr, cfg.epochs, cfg.providers, cfg.unclaimedRewards)

	ftsoManagerAddr := testcontracts.DeployManager(auth, be, rewardAddr, cfg.rewardEpochDurationSeconds, cfg.rewardEpochsStartTs,
		cfg.currentRewardEpoch, cfg.fraction, cfg.epochs, cfg.rewardEpochPowerHeights, cfg.rewardEpochStartHeights, cfg.rewardEpochStartTimes)

	ftsoRegistryAddr := testcontracts.DeployFTSORegistry(auth, be, cfg.supportedIndices)

	voterWhitelisterAddr := testcontracts.DeployWhitelist(auth, be, cfg.supportedIndices, cfg.providers)

	validationAddr = testcontracts.DeployValidator(auth, be, cfg.providers, cfg.nodes)

	submitterAddr = testcontracts.DeploySubmitter(auth, be, voterWhitelisterAddr, ftsoRegistryAddr, ftsoManagerAddr)

	return
}

var defaultTestContractsConfig = testContractsConfig{
	supportedIndices:           []*big.Int{big.NewInt(123), big.NewInt(321)},
	nodes:                      [][20]byte{{1, 2}, {3, 4}},
	providers:                  []common.Address{common.BigToAddress(big.NewInt(4441)), common.BigToAddress(big.NewInt(4442))},
	votePowers:                 []*big.Int{big.NewInt(1000), big.NewInt(2000)},
	unclaimedRewards:           []*big.Int{big.NewInt(88), big.NewInt(99)},
	epochs:                     []*big.Int{big.NewInt(0), big.NewInt(1)},
	rewardEpochDurationSeconds: big.NewInt(301),
	rewardEpochsStartTs:        big.NewInt(1),
	currentRewardEpoch:         big.NewInt(2),
	fraction:                   big.NewInt(10),
	rewardEpochPowerHeights:    []*big.Int{big.NewInt(302), big.NewInt(303)},
	rewardEpochStartHeights:    []*big.Int{big.NewInt(304), big.NewInt(305)},
	rewardEpochStartTimes:      []*big.Int{big.NewInt(306), big.NewInt(307)},
	totalSupply:                big.NewInt(1000),
}

type testContractsConfig struct {
	supportedIndices           []*big.Int
	nodes                      [][20]byte
	providers                  []common.Address
	votePowers                 []*big.Int
	unclaimedRewards           []*big.Int
	epochs                     []*big.Int
	rewardEpochDurationSeconds *big.Int
	rewardEpochsStartTs        *big.Int
	currentRewardEpoch         *big.Int
	fraction                   *big.Int
	rewardEpochPowerHeights    []*big.Int
	rewardEpochStartHeights    []*big.Int
	rewardEpochStartTimes      []*big.Int
	totalSupply                *big.Int
}

type testContractsConfigOption func(*testContractsConfig)

func withSupportedIndices(supportedIndices []*big.Int) testContractsConfigOption {
	return func(cfg *testContractsConfig) {
		cfg.supportedIndices = supportedIndices
	}
}

func withNodes(nodes [][20]byte) testContractsConfigOption {
	return func(cfg *testContractsConfig) {
		cfg.nodes = nodes
	}
}

func withProviders(providers []common.Address) testContractsConfigOption {
	return func(cfg *testContractsConfig) {
		cfg.providers = providers
	}
}

func withVotePowers(votePowers []*big.Int) testContractsConfigOption {
	return func(cfg *testContractsConfig) {
		cfg.votePowers = votePowers
	}
}

func withUnclaimedRewards(unclaimedRewards []*big.Int) testContractsConfigOption {
	return func(cfg *testContractsConfig) {
		cfg.unclaimedRewards = unclaimedRewards
	}
}

func withEpochs(epochs []*big.Int) testContractsConfigOption {
	return func(cfg *testContractsConfig) {
		cfg.epochs = epochs
	}
}

func withRewardEpochDurationSeconds(rewardEpochDurationSeconds *big.Int) testContractsConfigOption {
	return func(cfg *testContractsConfig) {
		cfg.rewardEpochDurationSeconds = rewardEpochDurationSeconds
	}
}

func withRewardEpochsStartTs(rewardEpochsStartTs *big.Int) testContractsConfigOption {
	return func(cfg *testContractsConfig) {
		cfg.rewardEpochsStartTs = rewardEpochsStartTs
	}
}

func withCurrentRewardEpoch(currentRewardEpoch *big.Int) testContractsConfigOption {
	return func(cfg *testContractsConfig) {
		cfg.currentRewardEpoch = currentRewardEpoch
	}
}

func withRewardEpochPowerHeights(rewardEpochPowerHeights []*big.Int) testContractsConfigOption {
	return func(cfg *testContractsConfig) {
		cfg.rewardEpochPowerHeights = rewardEpochPowerHeights
	}
}

func withRewardEpochStartHeights(rewardEpochStartHeights []*big.Int) testContractsConfigOption {
	return func(cfg *testContractsConfig) {
		cfg.rewardEpochStartHeights = rewardEpochStartHeights
	}
}

func withRewardEpochStartTimes(rewardEpochStartTimes []*big.Int) testContractsConfigOption {
	return func(cfg *testContractsConfig) {
		cfg.rewardEpochStartTimes = rewardEpochStartTimes
	}
}

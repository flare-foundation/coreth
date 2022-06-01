//go:build integration
// +build integration

package ftso

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/coreth/accounts/abi/bind"
	"github.com/flare-foundation/coreth/accounts/abi/bind/backends"
	"github.com/flare-foundation/coreth/core"
	corevm "github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/params"
	"github.com/flare-foundation/coreth/plugin/evm/ftso/testcontracts"
)

func TestSystem_New(t *testing.T) {
	t.Run("handles no genesis price submitter", func(t *testing.T) {
		_, be := simulatedBackend(t)
		defer be.Close()

		currentSubmitterAddr := params.SubmitterAddress
		defer func() {
			params.SubmitterAddress = currentSubmitterAddr
		}()

		params.SubmitterAddress = common.BigToAddress(big.NewInt(3030303))

		be.Commit(true)

		blockchain := be.Blockchain()

		evm := testEVM(t, blockchain)

		_, err := NewSystem(evm)
		require.ErrorIs(t, err, ErrNoPriceSubmitter)
	})

	t.Run("handles FTSO system not deployed", func(t *testing.T) {
		cfg := defaultTestContractsConfig
		auth, be := simulatedBackend(t)
		defer be.Close()

		ftsoRegistryAddr := testcontracts.DeployFTSORegistry(auth, be, cfg.supportedIndices)

		voterWhitelisterAddr := testcontracts.DeployWhitelist(auth, be, cfg.supportedIndices, cfg.providers)

		submitterAddr := testcontracts.DeploySubmitter(auth, be, voterWhitelisterAddr, ftsoRegistryAddr, common.Address{})

		currentSubmitterAddr := params.SubmitterAddress
		defer func() {
			params.SubmitterAddress = currentSubmitterAddr
		}()

		params.SubmitterAddress = submitterAddr

		be.Commit(true)

		blockchain := be.Blockchain()
		evm := testEVM(t, blockchain)

		_, err := NewSystem(evm)
		require.ErrorIs(t, err, ErrFTSONotDeployed)
	})

	t.Run("FTSO system not activated", func(t *testing.T) {
		cfg := defaultTestContractsConfig
		withRewardEpochStartHeights([]*big.Int{big.NewInt(0), big.NewInt(305)})(&cfg)

		auth, be := simulatedBackend(t)
		defer be.Close()

		votePowerAddr := testcontracts.DeployVotepower(auth, be, cfg.providers, cfg.heights, cfg.votePowers)

		wnatAddr := testcontracts.DeployWNAT(auth, be, votePowerAddr, cfg.totalSupply)

		rewardAddr := testcontracts.DeployReward(auth, be, wnatAddr, cfg.epochs, cfg.providers, cfg.unclaimedRewards)

		ftsoManagerAddr := testcontracts.DeployManager(auth, be, rewardAddr, cfg.rewardEpochDurationSeconds, cfg.rewardEpochsStartTs,
			cfg.currentRewardEpoch, cfg.fraction, cfg.epochs, cfg.rewardEpochPowerHeights, cfg.rewardEpochStartHeights, cfg.rewardEpochStartTimes)

		ftsoRegistryAddr := testcontracts.DeployFTSORegistry(auth, be, cfg.supportedIndices)

		voterWhitelisterAddr := testcontracts.DeployWhitelist(auth, be, cfg.supportedIndices, cfg.providers)

		submitterAddr := testcontracts.DeploySubmitter(auth, be, voterWhitelisterAddr, ftsoRegistryAddr, ftsoManagerAddr)

		currentSubmitterAddr := params.SubmitterAddress
		defer func() {
			params.SubmitterAddress = currentSubmitterAddr
		}()

		params.SubmitterAddress = submitterAddr

		be.Commit(true)

		blockchain := be.Blockchain()

		evm := testEVM(t, blockchain)

		_, err := NewSystem(evm)
		require.ErrorIs(t, err, ErrFTSONotActive)
	})
}

func TestSystem_Current(t *testing.T) {
	cfg := defaultTestContractsConfig
	system := testSystem(t, cfg)

	t.Run("nominal case", func(t *testing.T) {
		epoch, err := system.Current()
		require.NoError(t, err)
		assert.Equal(t, cfg.currentRewardEpoch.Uint64(), epoch)
	})

	t.Run("handles wrong manager address", func(t *testing.T) {
		system.contracts.Manager.address = system.contracts.WNAT.address
		_, err := system.Current()
		require.Error(t, err)
	})
}

func TestSystem_Cap(t *testing.T) {
	cfg := defaultTestContractsConfig
	system := testSystem(t, cfg)

	t.Run("nominal case", func(t *testing.T) {
		res, err := system.Cap()
		require.NoError(t, err)

		capInt := big.NewInt(0).Div(cfg.totalSupply, cfg.fraction)
		capFloat := big.NewFloat(0).SetInt(capInt)
		expCap, _ := capFloat.Float64()

		assert.Equal(t, expCap, res)
	})

	t.Run("handles wrong WNAT address", func(t *testing.T) {
		currentWNATAddress := system.contracts.WNAT.address
		system.contracts.WNAT.address = system.contracts.Manager.address
		defer func() {
			system.contracts.WNAT.address = currentWNATAddress
		}()
		_, err := system.Cap()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "could not get total supply")
	})

	t.Run("handles wrong manager address", func(t *testing.T) {
		system.contracts.Manager.address = system.contracts.WNAT.address
		_, err := system.Cap()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "could not get votepower threshold fraction")
	})
}

func TestSystem_Whitelist(t *testing.T) {
	cfg := defaultTestContractsConfig
	system := testSystem(t, cfg)

	t.Run("nominal case", func(t *testing.T) {
		res, err := system.Whitelist()
		require.NoError(t, err)
		require.Len(t, res, len(cfg.providers))
		assert.Equal(t, res[0], cfg.providers[0])
		assert.Equal(t, res[1], cfg.providers[1])
	})

	t.Run("handles wrong registry address", func(t *testing.T) {
		currentRegistryAddress := system.contracts.Registry.address
		system.contracts.Registry.address = system.contracts.Manager.address
		defer func() {
			system.contracts.Registry.address = currentRegistryAddress
		}()
		_, err := system.Whitelist()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "could not get series indices")
	})

	t.Run("handles wrong whitelist address", func(t *testing.T) {
		currentWhitelistAddress := system.contracts.Whitelist.address
		system.contracts.Whitelist.address = system.contracts.WNAT.address
		defer func() {
			system.contracts.Whitelist.address = currentWhitelistAddress
		}()
		_, err := system.Whitelist()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "could not get provider addresses")
	})

	t.Run("handles no providers in the result set", func(t *testing.T) {
		auth, be := simulatedBackend(t)
		defer be.Close()

		votePowerAddr := testcontracts.DeployVotepower(auth, be, cfg.providers, cfg.heights, cfg.votePowers)

		wnatAddr := testcontracts.DeployWNAT(auth, be, votePowerAddr, cfg.totalSupply)

		rewardAddr := testcontracts.DeployReward(auth, be, wnatAddr, cfg.epochs, cfg.providers, cfg.unclaimedRewards)

		ftsoManagerAddr := testcontracts.DeployManager(auth, be, rewardAddr, cfg.rewardEpochDurationSeconds, cfg.rewardEpochsStartTs,
			cfg.currentRewardEpoch, cfg.fraction, cfg.epochs, cfg.rewardEpochPowerHeights, cfg.rewardEpochStartHeights, cfg.rewardEpochStartTimes)

		indices := []*big.Int{big.NewInt(88888), big.NewInt(99999)}
		ftsoRegistryAddr := testcontracts.DeployFTSORegistry(auth, be, indices)

		voterWhitelisterAddr := testcontracts.DeployWhitelist(auth, be, cfg.supportedIndices, cfg.providers)

		submitterAddr := testcontracts.DeploySubmitter(auth, be, voterWhitelisterAddr, ftsoRegistryAddr, ftsoManagerAddr)

		currentSubmitterAddr := params.SubmitterAddress
		defer func() {
			params.SubmitterAddress = currentSubmitterAddr
		}()

		params.SubmitterAddress = submitterAddr

		be.Commit(true)

		blockchain := be.Blockchain()

		evm := testEVM(t, blockchain)

		system, err := NewSystem(evm)
		require.NoError(t, err)

		res, err := system.Whitelist()
		require.NoError(t, err)
		require.Len(t, res, 0)
	})
}

func TestSystem_Votepower(t *testing.T) {
	cfg := defaultTestContractsConfig
	system := testSystem(t, cfg)

	t.Run("nominal case", func(t *testing.T) {
		res, err := system.Votepower(cfg.providers[1])
		require.NoError(t, err)

		vpFloat := big.NewFloat(0).SetInt(cfg.votePowers[1])
		votepower, _ := vpFloat.Float64()
		assert.Equal(t, votepower, res)
	})

	t.Run("handles zero votepower", func(t *testing.T) {
		res, err := system.Votepower(cfg.providers[0])
		require.NoError(t, err)
		assert.Equal(t, 0.0, res)
	})

	t.Run("handles could not get epoch info", func(t *testing.T) {
		withCurrentRewardEpoch(big.NewInt(10))(&cfg)
		system := testSystem(t, cfg)

		_, err := system.Votepower(cfg.providers[1])
		require.Error(t, err)
		assert.Contains(t, err.Error(), "could not get epoch info")
	})
}

func TestSystem_Rewards(t *testing.T) {
	cfg := defaultTestContractsConfig
	withCurrentRewardEpoch(big.NewInt(1))(&cfg)
	system := testSystem(t, cfg)

	t.Run("nominal case", func(t *testing.T) {
		res, err := system.Rewards(cfg.providers[1])
		require.NoError(t, err)

		rewardsFloat := big.NewFloat(0).SetInt(cfg.unclaimedRewards[1])
		rewards, _ := rewardsFloat.Float64()
		assert.Equal(t, rewards, res)
	})

	t.Run("no reward case", func(t *testing.T) {
		cfg := defaultTestContractsConfig
		system := testSystem(t, cfg)
		res, err := system.Rewards(cfg.providers[1])
		require.NoError(t, err)
		assert.Equal(t, 0.0, res)
	})

	t.Run("handles wrong manager address", func(t *testing.T) {
		currentManagerAddress := system.contracts.Manager.address
		system.contracts.Manager.address = system.contracts.WNAT.address
		defer func() {
			system.contracts.Manager.address = currentManagerAddress
		}()
		_, err := system.Rewards(cfg.providers[1])
		require.Error(t, err)
		assert.Contains(t, err.Error(), "could not get current epoch")
	})
}

func testSystem(t *testing.T, cfg testContractsConfig) *System {
	t.Helper()

	auth, be := simulatedBackend(t)
	defer be.Close()

	submitterAddr := deployAllContracts(t, auth, be, cfg)

	currentSubmitterAddr := params.SubmitterAddress
	defer func() {
		params.SubmitterAddress = currentSubmitterAddr
	}()

	params.SubmitterAddress = submitterAddr

	be.Commit(true)

	blockchain := be.Blockchain()

	evm := testEVM(t, blockchain)

	system, err := NewSystem(evm)
	require.NoError(t, err)

	return system
}

func testEVM(t *testing.T, blockchain *core.BlockChain) *corevm.EVM {
	t.Helper()

	stateDB, err := blockchain.State()
	require.NoError(t, err)

	chainConfig := blockchain.Config()

	blkContext := core.NewEVMBlockContext(blockchain.CurrentHeader(), blockchain, nil)
	evm := corevm.NewEVM(blkContext, corevm.TxContext{}, stateDB, chainConfig, corevm.Config{NoBaseFee: true})

	return evm
}

func deployAllContracts(t *testing.T, auth *bind.TransactOpts, be *backends.SimulatedBackend,
	cfg testContractsConfig) (submitterAddr common.Address) {

	t.Helper()

	votePowerAddr := testcontracts.DeployVotepower(auth, be, cfg.providers, cfg.heights, cfg.votePowers)

	wnatAddr := testcontracts.DeployWNAT(auth, be, votePowerAddr, cfg.totalSupply)

	rewardAddr := testcontracts.DeployReward(auth, be, wnatAddr, cfg.epochs, cfg.providers, cfg.unclaimedRewards)

	ftsoManagerAddr := testcontracts.DeployManager(auth, be, rewardAddr, cfg.rewardEpochDurationSeconds, cfg.rewardEpochsStartTs,
		cfg.currentRewardEpoch, cfg.fraction, cfg.epochs, cfg.rewardEpochPowerHeights, cfg.rewardEpochStartHeights, cfg.rewardEpochStartTimes)

	ftsoRegistryAddr := testcontracts.DeployFTSORegistry(auth, be, cfg.supportedIndices)

	voterWhitelisterAddr := testcontracts.DeployWhitelist(auth, be, cfg.supportedIndices, cfg.providers)

	submitterAddr = testcontracts.DeploySubmitter(auth, be, voterWhitelisterAddr, ftsoRegistryAddr, ftsoManagerAddr)

	return
}

var defaultTestContractsConfig = testContractsConfig{
	supportedIndices:           []*big.Int{big.NewInt(123), big.NewInt(321)},
	providers:                  []common.Address{common.BigToAddress(big.NewInt(4441)), common.BigToAddress(big.NewInt(4442))},
	heights:                    []*big.Int{big.NewInt(302), big.NewInt(303)},
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
	providers                  []common.Address
	heights                    []*big.Int
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

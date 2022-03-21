//go:build integration
// +build integration

package evm

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/coreth/accounts/abi/bind/backends"
)

func TestFTSOSystem_Contracts(t *testing.T) {
	initTestContracts(t)

	t.Run("nominal case", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		vpInt := big.NewInt(100)

		unclaimedReward := big.NewInt(200)

		rewardEpochDurationSeconds := big.NewInt(301)
		rewardEpochPowerHeight := big.NewInt(302)
		rewardEpochStartHeight := big.NewInt(303)
		rewardEpochStartTime := big.NewInt(304)
		rewardEpochsStartTs := big.NewInt(1) // this should be less than header.Time()

		index1 := big.NewInt(123)
		index2 := big.NewInt(321)
		supportedIndices := []*big.Int{index1, index2}

		indices := []*big.Int{big.NewInt(401)}
		priceProvidersAddress1 := common.BigToAddress(big.NewInt(402))
		priceProvidersAddresses := []common.Address{priceProvidersAddress1}

		providerAddresses := []common.Address{common.BigToAddress(big.NewInt(500))}
		nodes := [][20]byte{{1, 2}}

		// VotePower contract
		votePowerAddr := deployTestContract(t, auth, be, testAbiVotePower, testAbiVotePowerBin, vpInt)

		// Wnat contract
		wnatAddr := deployTestContract(t, auth, be, testAbiWnat, testAbiWnatBin, votePowerAddr)

		// Reward contract
		rewardAddr := deployTestContract(t, auth, be, testAbiReward, testAbiRewardBin, wnatAddr, unclaimedReward)

		// Manager contract
		ftsoManagerAddr := deployTestContract(t, auth, be, testAbiManager, testAbiManagerBin, rewardAddr, rewardEpochDurationSeconds, rewardEpochPowerHeight, rewardEpochStartHeight, rewardEpochStartTime, rewardEpochsStartTs)

		// Registry contract
		ftsoRegistryAddr := deployTestContract(t, auth, be, testAbiFtsoRegistry, testAbiFtsoRegistryBin, supportedIndices)

		// Whitelist contract
		voterWhitelisterAddr := deployTestContract(t, auth, be, testAbiVoterWhitelister, testAbiVoterWhitelisterBin, indices, priceProvidersAddresses)

		// Validation contract
		validationAddr := deployTestContract(t, auth, be, testAbiValidation, testAbiValidationBin, providerAddresses, nodes)

		// Submitter contract
		submitterAddr := deployTestContract(t, auth, be, testAbiSubmitter, testAbiSubmitterBin, voterWhitelisterAddr, ftsoRegistryAddr, ftsoManagerAddr)

		be.Commit(true)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		latestBlock := be.Blockchain().LastAcceptedBlock()

		contracts, err := ftsoSystem.Contracts(latestBlock.Hash())
		assert.NoError(t, err)

		assert.Equal(t, ftsoRegistryAddr, contracts.Registry.address)
		assert.Equal(t, ftsoManagerAddr, contracts.Manager.address)
		assert.Equal(t, rewardAddr, contracts.Rewards.address)
		assert.Equal(t, voterWhitelisterAddr, contracts.Whitelist.address)
		assert.Equal(t, votePowerAddr, contracts.Votepower.address)
		assert.Equal(t, validationAddr, contracts.Validation.address)
	})

	t.Run("handles FTSO not deployed error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		// Submitter contract
		submitterAddr := deployTestContract(t, auth, be, testAbiSubmitter, testAbiSubmitterBin, common.Address{}, common.Address{}, common.Address{})

		be.Commit(true)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, common.Address{})

		latestBlock := be.Blockchain().LastAcceptedBlock()

		_, err := ftsoSystem.Contracts(latestBlock.Hash())
		assert.ErrorIs(t, err, errFTSONotDeployed)
	})

	t.Run("handles FTSO not active error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		rewardEpochDurationSeconds := big.NewInt(301)
		rewardEpochPowerHeight := big.NewInt(302)
		rewardEpochStartHeight := big.NewInt(303)
		rewardEpochStartTime := big.NewInt(304)
		rewardEpochsStartTs := big.NewInt(1000000)

		// Manager contract
		ftsoManagerAddr := deployTestContract(t, auth, be, testAbiManager, testAbiManagerBin, common.Address{}, rewardEpochDurationSeconds, rewardEpochPowerHeight, rewardEpochStartHeight, rewardEpochStartTime, rewardEpochsStartTs)

		// Submitter contract
		submitterAddr := deployTestContract(t, auth, be, testAbiSubmitter, testAbiSubmitterBin, common.Address{}, common.Address{}, ftsoManagerAddr)

		be.Commit(true)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, common.Address{})

		latestBlock := be.Blockchain().LastAcceptedBlock()

		_, err := ftsoSystem.Contracts(latestBlock.Hash())
		assert.ErrorIs(t, err, errFTSONotActive)
	})
}

func TestFTSOSystem_Details(t *testing.T) {
	initTestContracts(t)

	auth, be := simulatedBlockchain(t)
	defer be.Close()

	vpInt := big.NewInt(100)

	unclaimedReward := big.NewInt(200)

	rewardEpochDurationSeconds := big.NewInt(301)
	rewardEpochPowerHeight := big.NewInt(302)
	rewardEpochStartHeight := big.NewInt(303)
	rewardEpochStartTime := big.NewInt(304)
	rewardEpochsStartTs := big.NewInt(1) // this should be less than header.Time()

	index1 := big.NewInt(123)
	index2 := big.NewInt(321)
	supportedIndices := []*big.Int{index1, index2}

	indices := []*big.Int{big.NewInt(401)}
	priceProvidersAddress1 := common.BigToAddress(big.NewInt(402))
	priceProvidersAddresses := []common.Address{priceProvidersAddress1}

	providerAddresses := []common.Address{common.BigToAddress(big.NewInt(500))}
	nodes := [][20]byte{{1, 2}}

	// VotePower contract
	votePowerAddr := deployTestContract(t, auth, be, testAbiVotePower, testAbiVotePowerBin, vpInt)

	// Wnat contract
	wnatAddr := deployTestContract(t, auth, be, testAbiWnat, testAbiWnatBin, votePowerAddr)

	// Reward contract
	rewardAddr := deployTestContract(t, auth, be, testAbiReward, testAbiRewardBin, wnatAddr, unclaimedReward)

	// Manager contract
	ftsoManagerAddr := deployTestContract(t, auth, be, testAbiManager, testAbiManagerBin, rewardAddr, rewardEpochDurationSeconds, rewardEpochPowerHeight, rewardEpochStartHeight, rewardEpochStartTime, rewardEpochsStartTs)

	// Registry contract
	ftsoRegistryAddr := deployTestContract(t, auth, be, testAbiFtsoRegistry, testAbiFtsoRegistryBin, supportedIndices)

	// Whitelist contract
	voterWhitelisterAddr := deployTestContract(t, auth, be, testAbiVoterWhitelister, testAbiVoterWhitelisterBin, indices, priceProvidersAddresses)

	// Validation contract
	validationAddr := deployTestContract(t, auth, be, testAbiValidation, testAbiValidationBin, providerAddresses, nodes)

	// Submitter contract
	submitterAddr := deployTestContract(t, auth, be, testAbiSubmitter, testAbiSubmitterBin, voterWhitelisterAddr, ftsoRegistryAddr, ftsoManagerAddr)

	be.Commit(true)

	ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

	epoch := big.NewInt(1).Uint64()
	details, err := ftsoSystem.Details(epoch)
	require.NoError(t, err)

	assert.Equal(t, rewardEpochPowerHeight.Uint64(), details.PowerHeight)
	assert.Equal(t, rewardEpochStartHeight.Uint64(), details.StartHeight)
	assert.Equal(t, rewardEpochStartTime.Uint64(), details.StartTime)
	assert.Equal(t, rewardEpochStartTime.Uint64()+rewardEpochDurationSeconds.Uint64(), details.EndTime)
}

func TestFTSOSystem_Snapshot(t *testing.T) {
	initTestContracts(t)

	t.Run("nominal case", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		vpInt := big.NewInt(100)

		unclaimedReward := big.NewInt(200)

		rewardEpochDurationSeconds := big.NewInt(301)
		rewardEpochPowerHeight := big.NewInt(5)
		rewardEpochStartHeight := big.NewInt(4)
		rewardEpochStartTime := big.NewInt(304)
		rewardEpochsStartTs := big.NewInt(1) // this should be less than header.Time()

		index1 := big.NewInt(123)
		index2 := big.NewInt(321)
		supportedIndices := []*big.Int{index1, index2}

		indices := []*big.Int{big.NewInt(401)}
		priceProvidersAddress1 := common.BigToAddress(big.NewInt(402))
		priceProvidersAddresses := []common.Address{priceProvidersAddress1}

		providerAddresses := []common.Address{common.BigToAddress(big.NewInt(500))}
		nodes := [][20]byte{{1, 2}}

		// VotePower contract
		votePowerAddr := deployTestContract(t, auth, be, testAbiVotePower, testAbiVotePowerBin, vpInt)

		// Wnat contract
		wnatAddr := deployTestContract(t, auth, be, testAbiWnat, testAbiWnatBin, votePowerAddr)

		// Reward contract
		rewardAddr := deployTestContract(t, auth, be, testAbiReward, testAbiRewardBin, wnatAddr, unclaimedReward)

		// Manager contract
		ftsoManagerAddr := deployTestContract(t, auth, be, testAbiManager, testAbiManagerBin, rewardAddr, rewardEpochDurationSeconds, rewardEpochPowerHeight, rewardEpochStartHeight, rewardEpochStartTime, rewardEpochsStartTs)

		// Registry contract
		ftsoRegistryAddr := deployTestContract(t, auth, be, testAbiFtsoRegistry, testAbiFtsoRegistryBin, supportedIndices)

		// Whitelist contract
		voterWhitelisterAddr := deployTestContract(t, auth, be, testAbiVoterWhitelister, testAbiVoterWhitelisterBin, indices, priceProvidersAddresses)

		// Validation contract
		validationAddr := deployTestContract(t, auth, be, testAbiValidation, testAbiValidationBin, providerAddresses, nodes)

		// Submitter contract
		submitterAddr := deployTestContract(t, auth, be, testAbiSubmitter, testAbiSubmitterBin, voterWhitelisterAddr, ftsoRegistryAddr, ftsoManagerAddr)

		be.Commit(true)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		epoch := big.NewInt(1).Uint64()

		// crate four additional headers
		for i := 0; i < 4; i++ {
			be.Commit(true)
		}

		sn, err := ftsoSystem.Snapshot(epoch)
		require.NoError(t, err)

		powerHeader := be.Blockchain().GetHeaderByNumber(rewardEpochPowerHeight.Uint64()).Hash()
		startHeader := be.Blockchain().GetHeaderByNumber(rewardEpochStartHeight.Uint64()).Hash()

		snapshot := sn.(*FTSOSnapshot)
		assert.Equal(t, snapshot.power, powerHeader)
		assert.Equal(t, snapshot.start, startHeader)
		assert.Equal(t, snapshot.end, startHeader)
	})

	t.Run("handles unknown power block error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		vpInt := big.NewInt(100)

		unclaimedReward := big.NewInt(200)

		rewardEpochDurationSeconds := big.NewInt(301)
		rewardEpochPowerHeight := big.NewInt(500)
		rewardEpochStartHeight := big.NewInt(4)
		rewardEpochStartTime := big.NewInt(304)
		rewardEpochsStartTs := big.NewInt(1) // this should be less than header.Time()

		index1 := big.NewInt(123)
		index2 := big.NewInt(321)
		supportedIndices := []*big.Int{index1, index2}

		indices := []*big.Int{big.NewInt(401)}
		priceProvidersAddress1 := common.BigToAddress(big.NewInt(402))
		priceProvidersAddresses := []common.Address{priceProvidersAddress1}

		providerAddresses := []common.Address{common.BigToAddress(big.NewInt(500))}
		nodes := [][20]byte{{1, 2}}

		// VotePower contract
		votePowerAddr := deployTestContract(t, auth, be, testAbiVotePower, testAbiVotePowerBin, vpInt)

		// Wnat contract
		wnatAddr := deployTestContract(t, auth, be, testAbiWnat, testAbiWnatBin, votePowerAddr)

		// Reward contract
		rewardAddr := deployTestContract(t, auth, be, testAbiReward, testAbiRewardBin, wnatAddr, unclaimedReward)

		// Manager contract
		ftsoManagerAddr := deployTestContract(t, auth, be, testAbiManager, testAbiManagerBin, rewardAddr, rewardEpochDurationSeconds, rewardEpochPowerHeight, rewardEpochStartHeight, rewardEpochStartTime, rewardEpochsStartTs)

		// Registry contract
		ftsoRegistryAddr := deployTestContract(t, auth, be, testAbiFtsoRegistry, testAbiFtsoRegistryBin, supportedIndices)

		// Whitelist contract
		voterWhitelisterAddr := deployTestContract(t, auth, be, testAbiVoterWhitelister, testAbiVoterWhitelisterBin, indices, priceProvidersAddresses)

		// Validation contract
		validationAddr := deployTestContract(t, auth, be, testAbiValidation, testAbiValidationBin, providerAddresses, nodes)

		// Submitter contract
		submitterAddr := deployTestContract(t, auth, be, testAbiSubmitter, testAbiSubmitterBin, voterWhitelisterAddr, ftsoRegistryAddr, ftsoManagerAddr)

		be.Commit(true)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		epoch := big.NewInt(1).Uint64()

		_, err := ftsoSystem.Snapshot(epoch)
		require.EqualError(t, err, `unknown power block (height: 500)`)
	})

	t.Run("handles unknown current block error", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)
		defer be.Close()

		vpInt := big.NewInt(100)

		unclaimedReward := big.NewInt(200)

		rewardEpochDurationSeconds := big.NewInt(301)
		rewardEpochPowerHeight := big.NewInt(5)
		rewardEpochStartHeight := big.NewInt(400)
		rewardEpochStartTime := big.NewInt(304)
		rewardEpochsStartTs := big.NewInt(1) // this should be less than header.Time()

		index1 := big.NewInt(123)
		index2 := big.NewInt(321)
		supportedIndices := []*big.Int{index1, index2}

		indices := []*big.Int{big.NewInt(401)}
		priceProvidersAddress1 := common.BigToAddress(big.NewInt(402))
		priceProvidersAddresses := []common.Address{priceProvidersAddress1}

		providerAddresses := []common.Address{common.BigToAddress(big.NewInt(500))}
		nodes := [][20]byte{{1, 2}}

		// VotePower contract
		votePowerAddr := deployTestContract(t, auth, be, testAbiVotePower, testAbiVotePowerBin, vpInt)

		// Wnat contract
		wnatAddr := deployTestContract(t, auth, be, testAbiWnat, testAbiWnatBin, votePowerAddr)

		// Reward contract
		rewardAddr := deployTestContract(t, auth, be, testAbiReward, testAbiRewardBin, wnatAddr, unclaimedReward)

		// Manager contract
		ftsoManagerAddr := deployTestContract(t, auth, be, testAbiManager, testAbiManagerBin, rewardAddr, rewardEpochDurationSeconds, rewardEpochPowerHeight, rewardEpochStartHeight, rewardEpochStartTime, rewardEpochsStartTs)

		// Registry contract
		ftsoRegistryAddr := deployTestContract(t, auth, be, testAbiFtsoRegistry, testAbiFtsoRegistryBin, supportedIndices)

		// Whitelist contract
		voterWhitelisterAddr := deployTestContract(t, auth, be, testAbiVoterWhitelister, testAbiVoterWhitelisterBin, indices, priceProvidersAddresses)

		// Validation contract
		validationAddr := deployTestContract(t, auth, be, testAbiValidation, testAbiValidationBin, providerAddresses, nodes)

		// Submitter contract
		submitterAddr := deployTestContract(t, auth, be, testAbiSubmitter, testAbiSubmitterBin, voterWhitelisterAddr, ftsoRegistryAddr, ftsoManagerAddr)

		be.Commit(true)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		epoch := big.NewInt(1).Uint64()

		// crate four additional headers
		for i := 0; i < 4; i++ {
			be.Commit(true)
		}

		_, err := ftsoSystem.Snapshot(epoch)
		require.EqualError(t, err, `unknown current block (height: 400)`)
	})
}

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

func testFTSOSystem(t *testing.T, be *backends.SimulatedBackend, submitterAddr, validationAddr common.Address) *FTSOSystem {
	t.Helper()

	submitter := EVMContract{
		address: submitterAddr,
		abi:     testAbiSubmitter,
	}

	validation := EVMContract{
		address: validationAddr,
		abi:     testAbiValidation,
	}

	abis := FTSOABIs{
		Registry:  testAbiFtsoRegistry,
		Manager:   testAbiManager,
		Rewards:   testAbiReward,
		WNAT:      testAbiWnat,
		Whitelist: testAbiVoterWhitelister,
		Votepower: testAbiVotePower,
	}

	f := FTSOSystem{
		blockchain: be.Blockchain(),
		submitter:  submitter,
		validation: validation,
		abis:       abis,
	}

	return &f
}

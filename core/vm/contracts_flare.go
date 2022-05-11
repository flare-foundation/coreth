package vm

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/coreth/internal/ethapi"
	"github.com/flare-foundation/coreth/params"
	"github.com/flare-foundation/coreth/plugin/evm/ftso"
	"github.com/flare-foundation/flare/ids"
)

var (
	errNoPriceSubmitter    = errors.New("no price submitter")
	errFTSONotDeployed     = errors.New("FTSO not deployed")
	errFTSONotActive       = errors.New("FTSO not active")
	errRegistryNotDeployed = errors.New("validators registry not deployed")
	errNoReturnData        = errors.New("no return data")
)

const (
	RootDegree      = 4
	RatioMultiplier = 100.0
)

type validatorRegistry struct {
}

// 1) Initialize FTSO system around the EVM.
// 2) Retrieve current reward epoch from EVM.
// 3) Retrieve last reward epoch from FTSO system.
// 4) Check against last validator list in precompiled contract state (just some DB wrapped around the EVM DB).
// 5) If the reward epoch increased, retrieve all provider details and compute FTSO provider weights.
// 6) If new weights calculated, insert into the DB (by epoch) and update latest epoch.
// 7) If we stored a new latest epoch, set precompiled contract address' byte code to hash of validator state.

// TODO do we need to implement these from original validator contract?
// getDataProviderForNodeId
// getNodeIdForDataProvider
// registerNodeIdAsDataProvider
// PRICE_SUBMITTER

func (v *validatorRegistry) Run(evm *EVM, caller ContractRef, address common.Address, input []byte, gas uint64, read bool) ([]byte, uint64, error) {
	// TODO calculate remaining gas

	// 1) Initialize FTSO system around the EVM.
	validator, err := ftso.NewValidator(evm)

	// TODO 2) Retrieve current reward epoch from EVM.
	var currentRewardEpoch uint64

	// 3) Retrieve last reward epoch from FTSO system.
	lastRewardEpoch, err := validator.LastRewardEpoch()
	if err != nil {
		return nil, 0, err
	}

	// TODO 4) Check against last validator list in precompiled contract state (
	//  just some DB wrapped around the EVM DB).
	//
	var providers []common.Address

	// 5) If the reward epoch increased, retrieve all provider details and compute FTSO provider weights.
	if currentRewardEpoch != lastRewardEpoch {
		cap, err := validator.Cap()
		if err != nil {
			return nil, 0, fmt.Errorf("could not get votepower cap: %w", err)
		}

		validators := make(map[ids.ShortID]uint64)

		for _, provider := range providers {

			validatorID, err := validator.Validator(provider)
			if err != nil {
				return nil, 0, fmt.Errorf("could not get FTSO validator (provider: %s): %w", provider, err)
			}
			if validatorID == ids.ShortEmpty {
				continue
			}

			votepower, err := validator.Votepower(provider, lastRewardEpoch)
			if err != nil {
				return nil, 0, fmt.Errorf("could not get vote power (provider: %s): %w", provider,
					err)
			}
			if votepower == 0 {
				continue
			}

			if votepower > cap {
				votepower = cap
			}

			rewards, err := validator.Rewards(provider, lastRewardEpoch)
			if err != nil {
				return nil, 0, fmt.Errorf("could not get rewards (provider: %s): %w", provider, err)
			}
			if rewards == 0 {
				continue
			}

			weight := uint64(math.Pow(votepower, 1.0/float64(RootDegree)) * (RatioMultiplier * rewards / votepower))

			validators[validatorID] = weight
		}

		// TODO	6) If new weights calculated, insert into the DB (by epoch) and update latest epoch.
		err = validator.UpdateLastRewardEpoch(lastRewardEpoch)
		if err != nil {
			return nil, 0, err
		}
	}

	// TODO serialize validators set
	var validatorsRaw []byte

	return validatorsRaw, 0, nil
}

type EVMContract struct {
	address common.Address
	abi     abi.ABI
}

const (
	jsonSubmitter        = `[{"type":"constructor","stateMutability":"nonpayable","inputs":[]},{"type":"event","name":"GovernanceProposed","inputs":[{"type":"address","name":"proposedGovernance","internalType":"address","indexed":false}],"anonymous":false},{"type":"event","name":"GovernanceUpdated","inputs":[{"type":"address","name":"oldGovernance","internalType":"address","indexed":false},{"type":"address","name":"newGoveranance","internalType":"address","indexed":false}],"anonymous":false},{"type":"event","name":"PriceHashesSubmitted","inputs":[{"type":"address","name":"submitter","internalType":"address","indexed":true},{"type":"uint256","name":"epochId","internalType":"uint256","indexed":true},{"type":"address[]","name":"ftsos","internalType":"contract IFtsoGenesis[]","indexed":false},{"type":"bytes32[]","name":"hashes","internalType":"bytes32[]","indexed":false},{"type":"uint256","name":"timestamp","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"PricesRevealed","inputs":[{"type":"address","name":"voter","internalType":"address","indexed":true},{"type":"uint256","name":"epochId","internalType":"uint256","indexed":true},{"type":"address[]","name":"ftsos","internalType":"contract IFtsoGenesis[]","indexed":false},{"type":"uint256[]","name":"prices","internalType":"uint256[]","indexed":false},{"type":"uint256[]","name":"randoms","internalType":"uint256[]","indexed":false},{"type":"uint256","name":"timestamp","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"claimGovernance","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"getFtsoManager","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"contract IFtsoRegistryGenesis"}],"name":"getFtsoRegistry","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address[]","name":"","internalType":"address[]"}],"name":"getTrustedAddresses","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"getVoterWhitelister","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"governance","inputs":[]},{"type":"function","stateMutability":"pure","outputs":[],"name":"initialise","inputs":[{"type":"address","name":"_governance","internalType":"address"}]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"initialiseFixedAddress","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"proposeGovernance","inputs":[{"type":"address","name":"_governance","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"proposedGovernance","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"revealPrices","inputs":[{"type":"uint256","name":"_epochId","internalType":"uint256"},{"type":"uint256[]","name":"_ftsoIndices","internalType":"uint256[]"},{"type":"uint256[]","name":"_prices","internalType":"uint256[]"},{"type":"uint256[]","name":"_randoms","internalType":"uint256[]"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"setContractAddresses","inputs":[{"type":"address","name":"_ftsoRegistry","internalType":"contract IFtsoRegistryGenesis"},{"type":"address","name":"_voterWhitelister","internalType":"address"},{"type":"address","name":"_ftsoManager","internalType":"address"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"setTrustedAddresses","inputs":[{"type":"address[]","name":"_trustedAddresses","internalType":"address[]"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"submitPriceHashes","inputs":[{"type":"uint256","name":"_epochId","internalType":"uint256"},{"type":"uint256[]","name":"_ftsoIndices","internalType":"uint256[]"},{"type":"bytes32[]","name":"_hashes","internalType":"bytes32[]"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"transferGovernance","inputs":[{"type":"address","name":"_governance","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"voterWhitelistBitmap","inputs":[{"type":"address","name":"_voter","internalType":"address"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"voterWhitelisted","inputs":[{"type":"address","name":"_voter","internalType":"address"},{"type":"uint256","name":"_ftsoIndex","internalType":"uint256"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"votersRemovedFromWhitelist","inputs":[{"type":"address[]","name":"_removedVoters","internalType":"address[]"},{"type":"uint256","name":"_ftsoIndex","internalType":"uint256"}]}]`
	managerAddressMethod = "getFtsoManager"

	jsonManager        = `[{"type":"constructor","stateMutability":"nonpayable","inputs":[{"type":"address","name":"_governance","internalType":"address"},{"type":"address","name":"_flareDaemon","internalType":"contract FlareDaemon"},{"type":"address","name":"_priceSubmitter","internalType":"contract IIPriceSubmitter"},{"type":"uint256","name":"_firstEpochStartTs","internalType":"uint256"},{"type":"uint256","name":"_priceEpochDurationSeconds","internalType":"uint256"},{"type":"uint256","name":"_revealEpochDurationSeconds","internalType":"uint256"},{"type":"uint256","name":"_rewardEpochsStartTs","internalType":"uint256"},{"type":"uint256","name":"_rewardEpochDurationSeconds","internalType":"uint256"},{"type":"uint256","name":"_votePowerIntervalFraction","internalType":"uint256"}]},{"type":"event","name":"CleanupBlockNumberManagerFailedForBlock","inputs":[{"type":"uint256","name":"blockNumber","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"CleanupBlockNumberManagerUnset","inputs":[],"anonymous":false},{"type":"event","name":"ClosingExpiredRewardEpochFailed","inputs":[{"type":"uint256","name":"_rewardEpoch","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"ContractRevertError","inputs":[{"type":"address","name":"theContract","internalType":"address","indexed":false},{"type":"uint256","name":"atBlock","internalType":"uint256","indexed":false},{"type":"string","name":"theMessage","internalType":"string","indexed":false}],"anonymous":false},{"type":"event","name":"DistributingRewardsFailed","inputs":[{"type":"address","name":"ftso","internalType":"address","indexed":false},{"type":"uint256","name":"epochId","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"FallbackMode","inputs":[{"type":"bool","name":"fallbackMode","internalType":"bool","indexed":false}],"anonymous":false},{"type":"event","name":"FinalizingPriceEpochFailed","inputs":[{"type":"address","name":"ftso","internalType":"contract IIFtso","indexed":false},{"type":"uint256","name":"epochId","internalType":"uint256","indexed":false},{"type":"uint8","name":"failingType","internalType":"enum IFtso.PriceFinalizationType","indexed":false}],"anonymous":false},{"type":"event","name":"FtsoAdded","inputs":[{"type":"address","name":"ftso","internalType":"contract IIFtso","indexed":false},{"type":"bool","name":"add","internalType":"bool","indexed":false}],"anonymous":false},{"type":"event","name":"FtsoFallbackMode","inputs":[{"type":"address","name":"ftso","internalType":"contract IIFtso","indexed":false},{"type":"bool","name":"fallbackMode","internalType":"bool","indexed":false}],"anonymous":false},{"type":"event","name":"GovernanceProposed","inputs":[{"type":"address","name":"proposedGovernance","internalType":"address","indexed":false}],"anonymous":false},{"type":"event","name":"GovernanceUpdated","inputs":[{"type":"address","name":"oldGovernance","internalType":"address","indexed":false},{"type":"address","name":"newGoveranance","internalType":"address","indexed":false}],"anonymous":false},{"type":"event","name":"InitializingCurrentEpochStateForRevealFailed","inputs":[{"type":"address","name":"ftso","internalType":"contract IIFtso","indexed":false},{"type":"uint256","name":"epochId","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"PriceEpochFinalized","inputs":[{"type":"address","name":"chosenFtso","internalType":"address","indexed":false},{"type":"uint256","name":"rewardEpochId","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"RewardEpochFinalized","inputs":[{"type":"uint256","name":"votepowerBlock","internalType":"uint256","indexed":false},{"type":"uint256","name":"startBlock","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"MAX_TRUSTED_ADDRESSES_LENGTH","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"activate","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"active","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"addFtso","inputs":[{"type":"address","name":"_ftso","internalType":"contract IIFtso"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"addRevertError","inputs":[{"type":"address","name":"revertedContract","internalType":"address"},{"type":"string","name":"message","internalType":"string"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"claimGovernance","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"contract CleanupBlockNumberManager"}],"name":"cleanupBlockNumberManager","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"daemonize","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint192","name":"totalRevertedErrors","internalType":"uint192"},{"type":"uint64","name":"lastErrorTypeIndex","internalType":"uint64"}],"name":"errorData","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"contract FlareDaemon"}],"name":"flareDaemon","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"contract IIFtsoRegistry"}],"name":"ftsoRegistry","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"priceEpochId","internalType":"uint256"},{"type":"uint256","name":"priceEpochStartTimestamp","internalType":"uint256"},{"type":"uint256","name":"priceEpochEndTimestamp","internalType":"uint256"},{"type":"uint256","name":"priceEpochRevealEndTimestamp","internalType":"uint256"},{"type":"uint256","name":"currentTimestamp","internalType":"uint256"}],"name":"getCurrentPriceEpochData","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"getCurrentRewardEpoch","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"bool","name":"_fallbackMode","internalType":"bool"},{"type":"address[]","name":"_ftsos","internalType":"contract IIFtso[]"},{"type":"bool[]","name":"_ftsoInFallbackMode","internalType":"bool[]"}],"name":"getFallbackMode","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address[]","name":"_ftsos","internalType":"contract IIFtso[]"}],"name":"getFtsos","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"_firstPriceEpochStartTs","internalType":"uint256"},{"type":"uint256","name":"_priceEpochDurationSeconds","internalType":"uint256"},{"type":"uint256","name":"_revealEpochDurationSeconds","internalType":"uint256"}],"name":"getPriceEpochConfiguration","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"contract IPriceSubmitter"}],"name":"getPriceSubmitter","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"getRewardEpochVotePowerBlock","inputs":[{"type":"uint256","name":"_rewardEpoch","internalType":"uint256"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"getVotePowerIntervalFraction","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"governance","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"initialise","inputs":[{"type":"address","name":"_governance","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"lastRewardedFtsoAddress","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"contract IIPriceSubmitter"}],"name":"priceSubmitter","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"proposeGovernance","inputs":[{"type":"address","name":"_governance","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"proposedGovernance","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"removeFtso","inputs":[{"type":"address","name":"_ftso","internalType":"contract IIFtso"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"replaceFtso","inputs":[{"type":"address","name":"_ftsoToRemove","internalType":"contract IIFtso"},{"type":"address","name":"_ftsoToAdd","internalType":"contract IIFtso"},{"type":"bool","name":"_copyCurrentPrice","internalType":"bool"},{"type":"bool","name":"_copyAssetOrAssetFtsos","internalType":"bool"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"rewardEpochDurationSeconds","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"votepowerBlock","internalType":"uint256"},{"type":"uint256","name":"startBlock","internalType":"uint256"},{"type":"uint256","name":"startTimestamp","internalType":"uint256"}],"name":"rewardEpochs","inputs":[{"type":"uint256","name":"","internalType":"uint256"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"rewardEpochsStartTs","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"contract IIFtsoRewardManager"}],"name":"rewardManager","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"setContractAddresses","inputs":[{"type":"address","name":"_rewardManager","internalType":"contract IIFtsoRewardManager"},{"type":"address","name":"_ftsoRegistry","internalType":"contract IIFtsoRegistry"},{"type":"address","name":"_voterWhitelister","internalType":"contract IIVoterWhitelister"},{"type":"address","name":"_supply","internalType":"contract IISupply"},{"type":"address","name":"_cleanupBlockNumberManager","internalType":"contract CleanupBlockNumberManager"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"setFallbackMode","inputs":[{"type":"bool","name":"_fallbackMode","internalType":"bool"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"setFtsoAsset","inputs":[{"type":"address","name":"_ftso","internalType":"contract IIFtso"},{"type":"address","name":"_asset","internalType":"contract IIVPToken"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"setFtsoAssetFtsos","inputs":[{"type":"address","name":"_ftso","internalType":"contract IIFtso"},{"type":"address[]","name":"_assetFtsos","internalType":"contract IIFtso[]"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"setFtsoFallbackMode","inputs":[{"type":"address","name":"_ftso","internalType":"contract IIFtso"},{"type":"bool","name":"_fallbackMode","internalType":"bool"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"setGovernanceParameters","inputs":[{"type":"uint256","name":"_maxVotePowerNatThresholdFraction","internalType":"uint256"},{"type":"uint256","name":"_maxVotePowerAssetThresholdFraction","internalType":"uint256"},{"type":"uint256","name":"_lowAssetUSDThreshold","internalType":"uint256"},{"type":"uint256","name":"_highAssetUSDThreshold","internalType":"uint256"},{"type":"uint256","name":"_highAssetTurnoutThresholdBIPS","internalType":"uint256"},{"type":"uint256","name":"_lowNatTurnoutThresholdBIPS","internalType":"uint256"},{"type":"uint256","name":"_rewardExpiryOffsetSeconds","internalType":"uint256"},{"type":"address[]","name":"_trustedAddresses","internalType":"address[]"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"maxVotePowerNatThresholdFraction","internalType":"uint256"},{"type":"uint256","name":"maxVotePowerAssetThresholdFraction","internalType":"uint256"},{"type":"uint256","name":"lowAssetUSDThreshold","internalType":"uint256"},{"type":"uint256","name":"highAssetUSDThreshold","internalType":"uint256"},{"type":"uint256","name":"highAssetTurnoutThresholdBIPS","internalType":"uint256"},{"type":"uint256","name":"lowNatTurnoutThresholdBIPS","internalType":"uint256"},{"type":"uint256","name":"rewardExpiryOffsetSeconds","internalType":"uint256"},{"type":"bool","name":"changed","internalType":"bool"},{"type":"bool","name":"initialized","internalType":"bool"}],"name":"settings","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256[]","name":"_lastErrorBlock","internalType":"uint256[]"},{"type":"uint256[]","name":"_numErrors","internalType":"uint256[]"},{"type":"string[]","name":"_errorString","internalType":"string[]"},{"type":"address[]","name":"_erroringContract","internalType":"address[]"},{"type":"uint256","name":"_totalRevertedErrors","internalType":"uint256"}],"name":"showLastRevertedError","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256[]","name":"_lastErrorBlock","internalType":"uint256[]"},{"type":"uint256[]","name":"_numErrors","internalType":"uint256[]"},{"type":"string[]","name":"_errorString","internalType":"string[]"},{"type":"address[]","name":"_erroringContract","internalType":"address[]"},{"type":"uint256","name":"_totalRevertedErrors","internalType":"uint256"}],"name":"showRevertedErrors","inputs":[{"type":"uint256","name":"startIndex","internalType":"uint256"},{"type":"uint256","name":"numErrorTypesToShow","internalType":"uint256"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"contract IISupply"}],"name":"supply","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"switchToFallbackMode","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"transferGovernance","inputs":[{"type":"address","name":"_governance","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"contract IIVoterWhitelister"}],"name":"voterWhitelister","inputs":[]}]`
	currentEpochMethod = "getCurrentRewardEpoch"
)

func currentEpoch(evm *EVM) (uint64, error) {
	abiSubmitter, err := abi.JSON(strings.NewReader(jsonSubmitter))
	if err != nil {
		panic(err)
	}
	submitter := EVMContract{
		address: params.SubmitterAddress,
		abi:     abiSubmitter,
	}

	var managerAddr common.Address
	err = submitter.Execute(evm, managerAddressMethod).Decode(&managerAddr)
	if errors.Is(err, errNoReturnData) {
		return 0, errNoPriceSubmitter
	}
	if err != nil {
		return 0, err
	}

	empty := common.Address{}
	if managerAddr == empty {
		return 0, errFTSONotDeployed
	}

	abiManager, err := abi.JSON(strings.NewReader(jsonManager))
	if err != nil {
		panic(err)
	}

	manager := EVMContract{
		address: managerAddr,
		abi:     abiManager,
	}

	epoch := &big.Int{}
	err = manager.Execute(evm, currentEpochMethod).Decode(&epoch)
	if err != nil {
		return 0, fmt.Errorf("could not execute current epoch retrieval: %w", err)
	}

	return epoch.Uint64(), nil
}

func (e *EVMContract) Execute(evm *EVM, method string, params ...interface{}) *ContractReturn {

	data, err := e.abi.Pack(method, params...)
	if err != nil {
		return &ContractReturn{err: fmt.Errorf("could not pack parameters: %w", err)}
	}

	input := hexutil.Bytes(data)
	args := ethapi.TransactionArgs{To: &e.address, Input: &input}
	msg, err := args.ToMessage(0, nil)
	if err != nil {
		return &ContractReturn{err: fmt.Errorf("could not convert arguments to message: %w", err)}
	}

	ret, _, err := evm.Call(vm.AccountRef(msg.From()), *msg.To(), msg.Data(), 2000000, msg.Value())
	if err != nil {
		return &ContractReturn{err: fmt.Errorf("could not make evm call: %w", err)}
	}
	if len(ret) == 0 {
		return &ContractReturn{err: errNoReturnData}
	}

	values, err := e.abi.Unpack(method, ret)
	if err != nil {
		return &ContractReturn{err: fmt.Errorf("could not unpack return data: %w", err)}
	}

	return &ContractReturn{values: values}
}

type ContractReturn struct {
	values []interface{}
	err    error
}

func (e *ContractReturn) Decode(values ...interface{}) error {

	if e.err != nil {
		return e.err
	}

	if len(e.values) != len(values) {
		return fmt.Errorf("invalid number of decode values (have: %d, want: %d)", len(values), len(e.values))
	}

	for i, val := range values {

		if val == nil {
			continue
		}

		ret := e.values[i]

		vv := reflect.ValueOf(val)
		if vv.IsNil() {
			continue
		}
		if vv.Kind() != reflect.Ptr {
			return fmt.Errorf("invalid non-pointer (index: %d, type: %T)", i, val)
		}

		iv := reflect.Indirect(vv)
		rv := reflect.ValueOf(ret)
		if iv.Kind() != rv.Kind() {
			return fmt.Errorf("invalid type for return value (have: %T, want: %T)", val, ret)
		}

		iv.Set(rv)
	}

	return nil
}

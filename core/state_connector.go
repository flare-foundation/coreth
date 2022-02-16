// (c) 2021, Flare Networks Limited. All rights reserved.
// Please see the file LICENSE for licensing terms.

package core

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/log"
	"github.com/flare-foundation/coreth/plugin/evm"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/coreth/core/vm"
)

var (
	flareChainID    = new(big.Int).SetUint64(14) // https://github.com/ethereum-lists/chains/blob/master/_data/chains/eip155-14.json
	songbirdChainID = new(big.Int).SetUint64(19) // https://github.com/ethereum-lists/chains/blob/master/_data/chains/eip155-19.json

	flareStateConnectorActivationTime    = new(big.Int).SetUint64(1000000000000)
	songbirdStateConnectorActivationTime = new(big.Int).SetUint64(1000000000000)
)

type AttestationVotes struct {
	reachedMajority    bool
	majorityDecision   string
	majorityAttestors  []common.Address
	divergentAttestors []common.Address
	abstainedAttestors []common.Address
}

func GetTestingChain(chainID *big.Int) bool {
	return chainID.Cmp(flareChainID) != 0 && chainID.Cmp(songbirdChainID) != 0
}

func GetStateConnectorActivated(chainID *big.Int, blockTime *big.Int) bool {
	if GetTestingChain(chainID) {
		return true
	} else if chainID.Cmp(flareChainID) == 0 {
		return blockTime.Cmp(flareStateConnectorActivationTime) >= 0
	} else if chainID.Cmp(songbirdChainID) == 0 {
		return blockTime.Cmp(songbirdStateConnectorActivationTime) >= 0
	}
	return false
}

func GetStateConnectorContract(chainID *big.Int, blockTime *big.Int) common.Address {
	switch {
	case GetStateConnectorActivated(chainID, blockTime) && chainID.Cmp(songbirdChainID) == 0:
		return common.HexToAddress("0x6b5DEa84F71052c1302b5fe652e17FD442D126a9")
	default:
		return common.HexToAddress("0x1000000000000000000000000000000000000001")
	}
}

func GetStateConnectorCoinbaseSignalAddr(chainID *big.Int, blockTime *big.Int) common.Address {
	switch {
	default:
		return common.HexToAddress("0x000000000000000000000000000000000000dEaD")
	}
}

func SubmitAttestationSelector(chainID *big.Int, blockTime *big.Int) []byte {
	switch {
	default:
		return []byte{0xcf, 0xd1, 0xfd, 0xad}
	}
}

func GetAttestationSelector(chainID *big.Int, blockTime *big.Int) []byte {
	switch {
	default:
		return []byte{0x29, 0xbe, 0x4d, 0xb2}
	}
}

func FinaliseRoundSelector(chainID *big.Int, blockTime *big.Int) []byte {
	switch {
	default:
		return []byte{0xea, 0xeb, 0xf6, 0xd3}
	}
}

func GetVoterWhitelisterSelector(chainID *big.Int, blockTime *big.Int) []byte {
	switch {
	default:
		return []byte{0x71, 0xe1, 0xfa, 0xd9}
	}
}

func GetFtsoWhitelistedPriceProvidersSelector(chainID *big.Int, blockTime *big.Int) []byte {
	switch {
	default:
		return []byte{0x09, 0xfc, 0xb4, 0x00}
	}
}

// The default attestors are the FTSO price providers
func (st *StateTransition) GetDefaultAttestors(chainID *big.Int, timestamp *big.Int) ([]common.Address, error) {
	log.Info("GetDefaultAttestors called...")
	if os.Getenv("TESTING_ATTESTATION_PROVIDERS") != "" && GetTestingChain(chainID) {
		log.Info("GetDefaultAttestors called...1")
		return GetEnvAttestationProviders("TESTING"), nil
	} else {
		log.Info("GetDefaultAttestors called...2")
		// Get VoterWhitelister contract
		voterWhitelisterContractBytes, _, err := st.evm.Call(
			vm.AccountRef(st.msg.From()),
			common.HexToAddress(GetPrioritisedFTSOContract(timestamp)),
			GetVoterWhitelisterSelector(chainID, timestamp),
			GetKeeperGasMultiplier(st.evm.Context.BlockNumber)*st.evm.Context.GasLimit,
			big.NewInt(0))
		g := GetKeeperGasMultiplier(st.evm.Context.BlockNumber) * st.evm.Context.GasLimit
		log.Info("Gas in evm call: ", "gas", g)
		log.Info("Gas in evm call: ", "GetKeeperGasMultiplier", GetKeeperGasMultiplier(st.evm.Context.BlockNumber))
		log.Info("Gas in evm call: ", "GasLimit", st.evm.Context.GasLimit)
		if err != nil {
			return []common.Address{}, err
		}
		// Get FTSO prive providers
		voterWhitelisterContract := common.BytesToAddress(voterWhitelisterContractBytes)
		priceProvidersBytes, _, err := st.evm.Call(
			vm.AccountRef(st.msg.From()),
			voterWhitelisterContract,
			GetFtsoWhitelistedPriceProvidersSelector(chainID, timestamp),
			GetKeeperGasMultiplier(st.evm.Context.BlockNumber)*st.evm.Context.GasLimit,
			big.NewInt(0))
		if err != nil {
			return []common.Address{}, err
		}
		NUM_ATTESTORS := len(priceProvidersBytes) / 32
		var attestors []common.Address
		for i := 0; i < NUM_ATTESTORS; i++ {
			attestors = append(attestors, common.BytesToAddress(priceProvidersBytes[i*32:(i+1)*32]))
		}
		return attestors, nil
	}
}

func GetEnvAttestationProviders(attestorType string) []common.Address {
	envAttestationProvidersString := os.Getenv(attestorType + "_ATTESTATION_PROVIDERS")
	if envAttestationProvidersString == "" {
		return []common.Address{}
	}
	envAttestationProviders := strings.Split(envAttestationProvidersString, ",")
	NUM_ATTESTORS := len(envAttestationProviders)
	var attestors []common.Address
	for i := 0; i < NUM_ATTESTORS; i++ {
		attestors = append(attestors, common.HexToAddress(envAttestationProviders[i]))
	}
	return attestors
}

func (st *StateTransition) GetAttestation(attestor common.Address, instructions []byte) (string, error) {
	merkleRootHash, _, err := st.evm.Call(vm.AccountRef(attestor), st.to(), instructions, 20000, big.NewInt(0))
	return hex.EncodeToString(merkleRootHash), err
}

func (st *StateTransition) CountAttestations(attestors []common.Address, instructions []byte) (AttestationVotes, error) {
	var attestationVotes AttestationVotes
	hashFrequencies := make(map[string][]common.Address)
	for i, a := range attestors {
		h, err := st.GetAttestation(a, instructions)
		if err != nil {
			attestationVotes.abstainedAttestors = append(attestationVotes.abstainedAttestors, a)
		}
		hashFrequencies[h] = append(hashFrequencies[h], attestors[i])
	}
	// Find the plurality
	var pluralityNum int
	var pluralityKey string
	for key, val := range hashFrequencies {
		if len(val) > pluralityNum {
			pluralityNum = len(val)
			pluralityKey = key
		}
	}
	if pluralityNum > len(attestors)/2 {
		attestationVotes.reachedMajority = true
		attestationVotes.majorityDecision = pluralityKey
		attestationVotes.majorityAttestors = hashFrequencies[pluralityKey]
	}
	for key, val := range hashFrequencies {
		if key != pluralityKey {
			attestationVotes.divergentAttestors = append(attestationVotes.divergentAttestors, val...)
		}
	}
	return attestationVotes, nil
}

func (st *StateTransition) FinalisePreviousRound(chainID *big.Int, timestamp *big.Int, currentRoundNumber []byte) error {
	getAttestationSelector := GetAttestationSelector(chainID, timestamp)
	instructions := append(getAttestationSelector[:], currentRoundNumber[:]...)
	defaultAttestors, err := st.GetDefaultAttestors(chainID, timestamp)
	if err != nil {
		return err
	}
	defaultAttestationVotes, err := st.CountAttestations(defaultAttestors, instructions)
	if err != nil {
		return err
	}
	localAttestors := GetEnvAttestationProviders("LOCAL")
	var finalityReached bool
	if len(localAttestors) > 0 {
		localAttestationVotes, err := st.CountAttestations(localAttestors, instructions)
		if defaultAttestationVotes.reachedMajority && localAttestationVotes.reachedMajority && defaultAttestationVotes.majorityDecision == localAttestationVotes.majorityDecision {
			finalityReached = true
		} else if err != nil || (defaultAttestationVotes.reachedMajority && defaultAttestationVotes.majorityDecision != localAttestationVotes.majorityDecision) {
			// Make a back-up of the current state database, because this node is about to branch from the default set
		}
	} else if defaultAttestationVotes.reachedMajority {
		finalityReached = true
	}
	if finalityReached {
		// Finalise defaultAttestationVotes.majorityDecision
		finaliseRoundSelector := FinaliseRoundSelector(chainID, timestamp)
		finalisedData := append(finaliseRoundSelector[:], currentRoundNumber[:]...)
		merkleRootHashBytes, err := hex.DecodeString(defaultAttestationVotes.majorityDecision)
		if err != nil {
			return err
		}
		finalisedData = append(finalisedData[:], merkleRootHashBytes[:]...)
		coinbaseSignal := GetStateConnectorCoinbaseSignalAddr(chainID, timestamp)
		originalCoinbase := st.evm.Context.Coinbase
		defer func() {
			st.evm.Context.Coinbase = originalCoinbase
		}()
		st.evm.Context.Coinbase = coinbaseSignal

		_, _, err = st.evm.Call(vm.AccountRef(coinbaseSignal), st.to(), finalisedData, st.evm.Context.GasLimit, new(big.Int).SetUint64(0))
		if err != nil {
			return err
		}

		// Issue rewards to defaultAttestationVotes.majorityAttestors here:
	}
	return nil
}

func (st *StateTransition) GetFTSOManagerContract(blockTime *big.Int, chainID *big.Int) (common.Address, error) { //"0xbfA12e4E1411B62EdA8B035d71735667422A6A9e"
	FTSOManagerContractBytes, _, err := st.evm.Call(
		vm.AccountRef(st.msg.From()),
		common.HexToAddress(GetPrioritisedFTSOContract(blockTime)),
		evm.GetFtsoManagerSelector(chainID, blockTime),
		GetKeeperGasMultiplier(st.evm.Context.BlockNumber)*st.evm.Context.GasLimit,
		big.NewInt(0))
	if err != nil {
		return common.Address{}, err
	}

	FTSOManagerContract := common.BytesToAddress(FTSOManagerContractBytes)

	return FTSOManagerContract, nil
}

func (st *StateTransition) GetRewardRate() (float64, error) {

	ftsoManagerContractAddress, err := st.GetFTSOManagerContract(nil, nil) // todo fill the fields
	if err != nil {
		return -1, err
	}

	getCurrentRewardEpochBytes, _, err := st.evm.Call(
		vm.AccountRef(st.msg.From()),
		ftsoManagerContractAddress,
		evm.GetCurrentRewardEpochSelector(),
		GetKeeperGasMultiplier(st.evm.Context.BlockNumber)*st.evm.Context.GasLimit,
		big.NewInt(0)) // uint256

	currentRewardEpoch := int64(binary.BigEndian.Uint64(getCurrentRewardEpochBytes))

	getRewardEpochVotePowerBlockBytes, _, err := st.evm.Call(
		vm.AccountRef(st.msg.From()),
		ftsoManagerContractAddress,
		evm.GetRewardEpochVotePowerBlockSelector(currentRewardEpoch), // todo fill the field with currentRewardEpoch
		GetKeeperGasMultiplier(st.evm.Context.BlockNumber)*st.evm.Context.GasLimit,
		big.NewInt(0)) // uint256

	rewardEpochVotePowerBlock := int64(binary.BigEndian.Uint64(getRewardEpochVotePowerBlockBytes))

	ftsosBytes, _, err := st.evm.Call(
		vm.AccountRef(st.msg.From()),
		ftsoManagerContractAddress,
		evm.GetFtsosSelector(),
		GetKeeperGasMultiplier(st.evm.Context.BlockNumber)*st.evm.Context.GasLimit,
		big.NewInt(0)) // uint256

	if err != nil {
		return -1, err
	}
	NUM_FTSOS := len(ftsosBytes) / 32
	var ftsos []common.Address
	for i := 0; i < NUM_FTSOS; i++ {
		ftsos = append(ftsos, common.BytesToAddress(ftsosBytes[i*32:(i+1)*32]))
	}

	ftsoContract := ftsos[0]
	// Now from this ftso contract retrieve wNat
	wNatBytes, _, err := st.evm.Call(
		vm.AccountRef(st.msg.From()),
		ftsoContract,
		evm.GetWnatSelector(), //todo check if this is the correct way. If not, now to get public variables?
		GetKeeperGasMultiplier(st.evm.Context.BlockNumber)*st.evm.Context.GasLimit,
		big.NewInt(0)) // uint256

	if err != nil {
		return -1, err
	}

	wNatContract := common.BytesToAddress(wNatBytes)

	delegateWeiBytes, _, err := st.evm.Call(
		vm.AccountRef(st.msg.From()),
		wNatContract,
		evm.VotePowerOfAtSelector(), //todo add parameters as this is empty!
		GetKeeperGasMultiplier(st.evm.Context.BlockNumber)*st.evm.Context.GasLimit,
		big.NewInt(0)) // uint256

	if err != nil {
		return -1, err
	}

	delegatedAmount := int64(binary.BigEndian.Uint64(delegateWeiBytes)) / 1000000000000000000.0

	totalVotePowerBytes, _, err := st.evm.Call(
		vm.AccountRef(st.msg.From()),
		wNatContract,
		evm.TotalVotePowerSelector(),
		GetKeeperGasMultiplier(st.evm.Context.BlockNumber)*st.evm.Context.GasLimit,
		big.NewInt(0)) // uint256

	if err != nil {
		return -1, err
	}

	totalVotePowerAmount := int64(binary.BigEndian.Uint64(totalVotePowerBytes)) / 1000000000000000000.0

	votePower := delegatedAmount / totalVotePowerAmount * 100

	// todo get rewardManager contract
	FTSORewardManagerContractBytes, _, err := st.evm.Call(
		vm.AccountRef(st.msg.From()),
		ftsoManagerContractAddress,
		evm.GetRewardManagerSelector(),
		GetKeeperGasMultiplier(st.evm.Context.BlockNumber)*st.evm.Context.GasLimit,
		big.NewInt(0)) // uint256

	if err != nil {
		return -1, err
	}

	FTSORewardManagerContract := common.BytesToAddress(FTSORewardManagerContractBytes)

	dataProviderCurrentFeePercentageBytes, _, err := st.evm.Call(
		vm.AccountRef(st.msg.From()),
		FTSORewardManagerContract,
		evm.GetDataProviderCurrentFeePercentageSelector(),
		GetKeeperGasMultiplier(st.evm.Context.BlockNumber)*st.evm.Context.GasLimit,
		big.NewInt(0)) // uint256

	if err != nil {
		return -1, err
	}

	dataProviderCurrentFeePercentage := int64(binary.BigEndian.Uint64(dataProviderCurrentFeePercentageBytes))
	fee := dataProviderCurrentFeePercentage / 100.0

	rewardsBytes, _, err := st.evm.Call(
		vm.AccountRef(st.msg.From()),
		FTSORewardManagerContract,
		evm.GetUnclaimedRewardSelector(), // todo fill the params
		GetKeeperGasMultiplier(st.evm.Context.BlockNumber)*st.evm.Context.GasLimit,
		big.NewInt(0)) // uint256

	if err != nil {
		return -1, err
	}

	rewards := int64(binary.BigEndian.Uint64(rewardsBytes))

	reward_rate := rewards / delegatedAmount * (1 - fee/100)

	return float64(reward_rate), nil
}

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
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/coreth/core/vm"
)

const (
	defaultAttestorEnv = "DEFAULT_ATTESTATORS"
	localAttestorEnv   = "LOCAL_ATTESTATORS"
)

var (
	costonChainID   = new(big.Int).SetUint64(16) // https://github.com/ethereum-lists/chains/blob/master/_data/chains/eip155-16.json
	songbirdChainID = new(big.Int).SetUint64(19) // https://github.com/ethereum-lists/chains/blob/master/_data/chains/eip155-19.json
	flareChainID    = new(big.Int).SetUint64(14) // https://github.com/ethereum-lists/chains/blob/master/_data/chains/eip155-14.json

	costonActivationTime   = big.NewInt(time.Date(2022, time.February, 25, 17, 0, 0, 0, time.UTC).Unix())
	songbirdActivationTime = big.NewInt(time.Date(2200, time.January, 1, 0, 0, 0, 0, time.UTC).Unix())
	flareActivationTime    = big.NewInt(time.Date(2200, time.January, 1, 0, 0, 0, 0, time.UTC).Unix())

	costonDefaultAttestors = []common.Address{
		common.HexToAddress("0x3a6e101103ec3d9267d08f484a6b70e1440a8255"),
	}
	songbirdDefaultAttestors = []common.Address{}
	flareDefaultAttestors    = []common.Address{}
)

type AttestationVotes struct {
	reachedMajority    bool
	majorityDecision   string
	majorityAttestors  []common.Address
	divergentAttestors []common.Address
	abstainedAttestors []common.Address
}

func GetStateConnectorActivated(chainID *big.Int, blockTime *big.Int) bool {
	switch {
	case chainID.Cmp(costonChainID) == 0:
		return blockTime.Cmp(costonActivationTime) >= 0
	case chainID.Cmp(songbirdChainID) == 0:
		return blockTime.Cmp(songbirdActivationTime) >= 0
	case chainID.Cmp(flareChainID) == 0:
		return blockTime.Cmp(flareActivationTime) >= 0
	default:
		return true
	}
}

func GetStateConnectorContract(chainID *big.Int, blockTime *big.Int) common.Address {
	switch {
	case chainID.Cmp(costonChainID) == 0:
		return common.HexToAddress("0x947c76694491d3fD67a73688003c4d36C8780A97")
	case chainID.Cmp(songbirdChainID) == 0:
		return common.HexToAddress("0x3A1b3220527aBA427d1e13e4b4c48c31460B4d91")
	case chainID.Cmp(flareChainID) == 0:
		return common.HexToAddress("0x1000000000000000000000000000000000000001")
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

func GetDefaultAttestors(chainID *big.Int) []common.Address {
	defaultAttestorList := os.Getenv(defaultAttestorEnv)
	if defaultAttestorList != "" {
		defaultAttestorEntries := strings.Split(defaultAttestorList, ",")
		defaultAttestors := make([]common.Address, 0, len(defaultAttestorEntries))
		for _, defaultAttestorEntry := range defaultAttestorEntries {
			defaultAttestors = append(defaultAttestors, common.HexToAddress(defaultAttestorEntry))
		}
		return defaultAttestors
	}
	switch {
	case chainID.Cmp(costonChainID) == 0:
		return costonDefaultAttestors
	case chainID.Cmp(songbirdChainID) == 0:
		return songbirdDefaultAttestors
	case chainID.Cmp(flareChainID) == 0:
		return flareDefaultAttestors
	default:
		return nil
	}
}

func GetLocalAttestors() []common.Address {
	localAttestorList := os.Getenv(localAttestorEnv)
	if localAttestorList != "" {
		localAttestorEntries := strings.Split(localAttestorList, ",")
		localAttestors := make([]common.Address, 0, len(localAttestorEntries))
		for _, localAttestorEntry := range localAttestorEntries {
			localAttestors = append(localAttestors, common.HexToAddress(localAttestorEntry))
		}
	}
	return nil
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
	defaultAttestors := GetDefaultAttestors(chainID)
	defaultAttestationVotes, err := st.CountAttestations(defaultAttestors, instructions)
	if err != nil {
		return err
	}
	localAttestors := GetLocalAttestors()
	var finalityReached bool
	if len(localAttestors) > 0 {
		localAttestationVotes, err := st.CountAttestations(localAttestors, instructions)
		if defaultAttestationVotes.reachedMajority && localAttestationVotes.reachedMajority && defaultAttestationVotes.majorityDecision == localAttestationVotes.majorityDecision {
			finalityReached = true
		} else if err != nil || (defaultAttestationVotes.reachedMajority && defaultAttestationVotes.majorityDecision != localAttestationVotes.majorityDecision) {
			// Make a back-up of the current state database, because this node is about to fork from the default set
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

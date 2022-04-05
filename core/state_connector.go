// (c) 2021, Flare Networks Limited. All rights reserved.
// Please see the file LICENSE for licensing terms.

package core

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/coreth/params"
)

const (
	defaultAttestorEnv = "DEFAULT_ATTESTATORS"
	localAttestorEnv   = "LOCAL_ATTESTATORS"
)

var (
	costonDefaultAttestors = []common.Address{
		common.HexToAddress("0x3a6e101103ec3d9267d08f484a6b70e1440a8255"),
	}
	songbirdDefaultAttestors = []common.Address{
		common.HexToAddress("0x0c19f3B4927abFc596353B0f9Ddad5D817736F70"),
	}
	flareDefaultAttestors []common.Address
)

// Caller is a light wrapper around Ethereum Virtual Machine
type Caller interface {
	Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error)
	SetBlockContext(bc vm.BlockContext)
	BlockContext() vm.BlockContext
}

// stateConnector is responsible for calling state connector smart contract based on votes from attestors.
type stateConnector struct {
	caller Caller
	msg    Message
}

func newConnector(caller Caller, msg Message) *stateConnector {
	s := &stateConnector{
		caller: caller,
		msg:    msg,
	}
	return s
}

func (c *stateConnector) finalizePreviousRound(chainID *big.Int, timestamp *big.Int, currentRoundNumber []byte) error {
	instructions := append(attestationSelector(chainID, timestamp), currentRoundNumber[:]...)
	defaultAttestors := defaultAttestors(chainID)

	defaultAttestationResult := c.countAttestations(defaultAttestors, instructions)

	reached := c.isFinalityReached(defaultAttestationResult, instructions)
	if !reached {
		return nil
	}

	// Finalize defaultAttestationResult.majorityDecision
	finalizeRoundSelector := finalizeRoundSelector(chainID, timestamp)
	finalizedData := append(finalizeRoundSelector[:], currentRoundNumber[:]...)
	merkleRootHashBytes, err := hex.DecodeString(defaultAttestationResult.majorityDecision)
	if err != nil {
		return fmt.Errorf("could not decode majority decision into hex: %w", err)
	}
	finalizedData = append(finalizedData[:], merkleRootHashBytes[:]...)

	bc := c.caller.BlockContext()

	// switch the address to be able to call the finalized function
	originalBC := bc
	defer func() {
		c.caller.SetBlockContext(originalBC)
	}()
	coinbaseSignal := stateConnectorCoinbaseSignalAddr(chainID, timestamp)
	bc.Coinbase = coinbaseSignal
	c.caller.SetBlockContext(bc)

	_, _, err = c.caller.Call(vm.AccountRef(coinbaseSignal), c.to(), finalizedData, bc.GasLimit, new(big.Int).SetUint64(0))
	if err != nil {
		return err
	}

	return nil
}

// defaultAttestors returns list of default attestors set by environment variable or based on chainID.
func defaultAttestors(chainID *big.Int) []common.Address {
	defaultAttestors := envAttestors(defaultAttestorEnv)
	if len(defaultAttestors) != 0 {
		return defaultAttestors
	}
	switch {
	case chainID.Cmp(params.CostonChainID) == 0:
		return costonDefaultAttestors
	case chainID.Cmp(params.SongbirdChainID) == 0:
		return songbirdDefaultAttestors
	case chainID.Cmp(params.FlareChainID) == 0:
		return flareDefaultAttestors
	default:
		return nil
	}
}

type attestationResult struct {
	reachedMajority  bool
	majorityDecision string
}

// isFinalityReached checks if finality is reached based on attestation votes.
func (c *stateConnector) isFinalityReached(defaultAttestationResult attestationResult, instructions []byte) bool {
	var finalityReached bool

	localAttestors := envAttestors(localAttestorEnv)
	if len(localAttestors) > 0 {
		localAttestationResult := c.countAttestations(localAttestors, instructions)
		if defaultAttestationResult.reachedMajority && localAttestationResult.reachedMajority && defaultAttestationResult.majorityDecision == localAttestationResult.majorityDecision {
			finalityReached = true
		} else if defaultAttestationResult.reachedMajority && defaultAttestationResult.majorityDecision != localAttestationResult.majorityDecision {
			// FIXME Make a back-up of the current state database, because this node is about to branch from the default set
		}
	} else if defaultAttestationResult.reachedMajority {
		finalityReached = true
	}

	return finalityReached
}

// countAttestations counts the number of the votes and determines whether majority is reached
func (c *stateConnector) countAttestations(attestors []common.Address, instructions []byte) attestationResult {
	var av attestationResult

	hashFrequencies := make(map[string][]common.Address, len(attestors))
	for i := range attestors {
		h, err := c.attestationResult(attestors[i], instructions)
		if err != nil {
			// FIXME: how to handle this error??
			continue
		}
		hashFrequencies[h] = append(hashFrequencies[h], attestors[i])
	}
	var majorityNum int
	var majorityKey string
	for key, val := range hashFrequencies {
		if len(val) > majorityNum {
			majorityNum = len(val)
			majorityKey = key
		}
	}
	if majorityNum > len(attestors)/2 {
		av.reachedMajority = true
		av.majorityDecision = majorityKey
	}

	return av
}

// attestationResult returns resulting hash from the attestor.
func (c *stateConnector) attestationResult(attestor common.Address, instructions []byte) (string, error) {
	rootHash, _, err := c.caller.Call(vm.AccountRef(attestor), c.to(), instructions, 20000, big.NewInt(0))
	return hex.EncodeToString(rootHash), err
}

// to returns the recipient of the message.
func (c *stateConnector) to() common.Address {
	// empty message or receiver means contract creation
	if c.msg == nil || c.msg.To() == nil {
		return common.Address{}
	}
	return *c.msg.To()
}

// envAttestors returns list of attestors from environment variable using provided key.
// Returns an empty list of value if the key does not exist in the environment.
func envAttestors(key string) []common.Address {
	envAttestationProvidersString := os.Getenv(key)
	if envAttestationProvidersString == "" {
		return nil
	}
	envAttestationProviders := strings.Split(envAttestationProvidersString, ",")
	attestors := make([]common.Address, len(envAttestationProviders))
	for i := range envAttestationProviders {
		attestors[i] = common.HexToAddress(envAttestationProviders[i])
	}
	return attestors
}

// unused 'blockTime' might be used for hard forks in the future.
func stateConnectorContract(chainID *big.Int, blockTime *big.Int) common.Address {
	switch {
	case chainID.Cmp(params.CostonChainID) == 0:
		return common.HexToAddress("0x947c76694491d3fD67a73688003c4d36C8780A97")
	case chainID.Cmp(params.SongbirdChainID) == 0:
		return common.HexToAddress("0x3A1b3220527aBA427d1e13e4b4c48c31460B4d91")
	case chainID.Cmp(params.FlareChainID) == 0:
		return common.HexToAddress("0x1000000000000000000000000000000000000001")
	default:
		return common.HexToAddress("0x1000000000000000000000000000000000000001")
	}
}

func isStateConnectorActivated(chainID *big.Int, blockTime *big.Int) bool {
	switch {
	case chainID.Cmp(params.CostonChainID) == 0:
		return blockTime.Cmp(costonActivationTime) >= 0
	case chainID.Cmp(params.SongbirdChainID) == 0:
		return blockTime.Cmp(songbirdActivationTime) >= 0
	case chainID.Cmp(params.FlareChainID) == 0:
		return blockTime.Cmp(flareActivationTime) >= 0
	default:
		return true
	}
}

// unused 'chainID' and 'blockTime' might be used for hard forks in the future.
func submitAttestationSelectorFunc(chainID *big.Int, blockTime *big.Int) []byte {
	switch {
	default:
		return []byte{0xcf, 0xd1, 0xfd, 0xad}
	}
}

// unused 'chainID' and 'blockTime' might be used for hard forks in the future.
func stateConnectorCoinbaseSignalAddr(chainID *big.Int, blockTime *big.Int) common.Address {
	switch {
	default:
		return common.HexToAddress("0x000000000000000000000000000000000000dEaD")
	}
}

// unused 'chainID' and 'blockTime' might be used for hard forks in the future.
func attestationSelector(chainID *big.Int, blockTime *big.Int) []byte {
	switch {
	default:
		return []byte{0x29, 0xbe, 0x4d, 0xb2}
	}
}

// unused 'chainID' and 'blockTime' might be used for hard forks in the future.
func finalizeRoundSelector(chainID *big.Int, blockTime *big.Int) []byte {
	switch {
	default:
		return []byte{0xea, 0xeb, 0xf6, 0xd3}
	}
}

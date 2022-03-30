// (c) 2021, Flare Networks Limited. All rights reserved.
// Please see the file LICENSE for licensing terms.

package core

import (
	"encoding/hex"
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
	WithBlockContext(bc vm.BlockContext)
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

	defaultAttestationVotes := c.countAttestations(defaultAttestors, instructions)

	reached := c.isFinalityReached(defaultAttestationVotes, instructions)
	if !reached {
		return nil
	}

	// Finalise defaultAttestationVotes.majorityDecision
	finaliseRoundSelector := finalizeRoundSelector(chainID, timestamp)
	finalizedData := append(finaliseRoundSelector[:], currentRoundNumber[:]...)
	merkleRootHashBytes, err := hex.DecodeString(defaultAttestationVotes.majorityDecision)
	if err != nil {
		return err
	}
	finalizedData = append(finalizedData[:], merkleRootHashBytes[:]...)

	bc := c.caller.BlockContext()

	// switch the address to be able to call the finalized function
	originalBC := bc
	defer func() {
		c.caller.WithBlockContext(originalBC)
	}()
	coinbaseSignal := stateConnectorCoinbaseSignalAddr(chainID, timestamp)
	bc.Coinbase = coinbaseSignal
	c.caller.WithBlockContext(bc)

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

type attestationVotes struct {
	reachedMajority    bool
	majorityDecision   string
	majorityAttestors  []common.Address
	divergentAttestors []common.Address
	abstainedAttestors []common.Address
}

// isFinalityReached checks if finality is reached based on attestation votes.
func (c *stateConnector) isFinalityReached(defaultAttestationVotes attestationVotes, instructions []byte) bool {
	var finalityReached bool

	localAttestors := envAttestors(localAttestorEnv)
	if len(localAttestors) > 0 {
		localAttestationVotes := c.countAttestations(localAttestors, instructions)
		if defaultAttestationVotes.reachedMajority && localAttestationVotes.reachedMajority && defaultAttestationVotes.majorityDecision == localAttestationVotes.majorityDecision {
			finalityReached = true
		} else if defaultAttestationVotes.reachedMajority && defaultAttestationVotes.majorityDecision != localAttestationVotes.majorityDecision {
			// FIXME Make a back-up of the current state database, because this node is about to branch from the default set
		}
	} else if defaultAttestationVotes.reachedMajority {
		finalityReached = true
	}

	return finalityReached
}

// countAttestations counts the number of the votes and determines whether majority is reached
func (c *stateConnector) countAttestations(attestors []common.Address, instructions []byte) attestationVotes {
	var av attestationVotes
	hashFrequencies := make(map[string][]common.Address, len(attestors))
	for i := range attestors {
		h, err := c.attestationResult(attestors[i], instructions)
		if err != nil {
			av.abstainedAttestors = append(av.abstainedAttestors, attestors[i])
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
		av.majorityAttestors = hashFrequencies[majorityKey]
	}
	for key, val := range hashFrequencies {
		if key != majorityKey {
			av.divergentAttestors = append(av.divergentAttestors, val...)
		}
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

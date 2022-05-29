package validatordb

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"

	"github.com/flare-foundation/flare/ids"

	"github.com/flare-foundation/coreth/core/state"
)

// State represents the validator state database at a specific point in time,
// as determined by the root hash of the state trie it encapsulates.
type State struct {
	enc  cbor.EncMode
	dec  cbor.DecMode
	trie state.Trie
}

// SetEpoch sets the active validator epoch. It is used to transition from a previous
// validator epoch to a subsequent validator epoch.
func (s *State) SetEpoch(epoch uint64) error {

	key := []byte{codeEpoch}

	val := make([]byte, 8)
	binary.BigEndian.PutUint64(val, epoch)

	err := s.trie.TryUpdate(key, val)
	if err != nil {
		return fmt.Errorf("could not update epoch: %w", err)
	}

	return nil
}

// GetEpoch gets the active validator epoch.
func (s *State) GetEpoch() (uint64, error) {

	key := []byte{codeEpoch}

	val, err := s.trie.TryGet(key)
	if err != nil {
		return 0, fmt.Errorf("could not get epoch: %w", err)
	}

	return binary.BigEndian.Uint64(val), nil
}

// SetMapping sets the mapping between an FTSO data provider address and the corresponding
// validator node ID. It is used by FTSO data providers to indicate which validator node
// ID they would like to assign their validation weight to.
func (s *State) SetMapping(provider common.Address, nodeID ids.ShortID) error {

	key := make([]byte, 21)
	key[0] = codeMapping
	copy(key[1:], provider[:])

	val := nodeID[:]

	err := s.trie.TryUpdate(key, val)
	if err != nil {
		return fmt.Errorf("could not set mapping: %w", err)
	}

	return nil
}

// GetMapping gets the mapping between an FTSO data provider address and the corresponding
// validator node ID.
func (s *State) GetMapping(provider common.Address) (ids.ShortID, error) {

	key := make([]byte, 21)
	key[0] = codeMapping
	copy(key[1:], provider[:])

	val, err := s.trie.TryGet(key)
	if err != nil {
		return ids.ShortID{}, fmt.Errorf("could not get mapping: %w", err)
	}

	validatorID, err := ids.ToShortID(val)
	if err != nil {
		return ids.ShortID{}, fmt.Errorf("could not parse mapping: %w", err)
	}

	return validatorID, nil
}

// AllMappings gets all mappings between FTSO data provider addresses and the corresponding
// validator node IDs.
func (s *State) AllMappings() (map[common.Address]ids.ShortID, error) {

	start := []byte{codeMapping}
	end := []byte{codeMapping + 1}

	validators := make(map[common.Address]ids.ShortID)
	it := s.trie.NodeIterator(start)
	for it.Next(true) {

		if !it.Leaf() {
			continue
		}

		key := it.LeafKey()
		if bytes.Compare(key, end) >= 0 {
			break
		}

		val := it.LeafBlob()
		validatorID, err := ids.ToShortID(val)
		if err != nil {
			return nil, fmt.Errorf("could not parse validator (val: %x): %w", val, err)
		}

		provider := common.BytesToAddress(key[1:])

		validators[provider] = validatorID
	}

	err := it.Error()
	if err != nil {
		return nil, fmt.Errorf("could not iterate validators: %w", err)
	}

	return validators, nil
}

// SetCandidates sets the validator candidates for the next validator epoch. It is
// used in a daemon contract call upon epoch switchover to prepare the validator
// information needed for the upcoming validator epoch.
func (s *State) SetCandidates(candidates []*Candidate) error {

	key := []byte{codeCandidates}

	val, err := s.enc.Marshal(candidates)
	if err != nil {
		return fmt.Errorf("could not encode candidates: %w", err)
	}

	err = s.trie.TryUpdate(key, val)
	if err != nil {
		return fmt.Errorf("could not put candidates: %w", err)
	}

	return nil
}

// GetCandidates gets the validator candidates for the next validator epoch.
func (s *State) GetCandidates() ([]*Candidate, error) {

	key := []byte{codeCandidates}

	val, err := s.trie.TryGet(key)
	if err != nil {
		return nil, fmt.Errorf("could not get candidates: %w", err)
	}

	var candidates []*Candidate
	err = s.dec.Unmarshal(val, &candidates)
	if err != nil {
		return nil, fmt.Errorf("could not decode candidates: %w", err)
	}

	return candidates, nil
}

// SetValidators sets the active validators for the current validator epoch. It is
// used in a daemon contract call upon epoch switchover to set the validator information
// for the active validator epoch.
func (s *State) SetValidators(validators []*Validator) error {

	key := []byte{codeValidators}

	val, err := s.enc.Marshal(validators)
	if err != nil {
		return fmt.Errorf("could not encode validators: %w", err)
	}

	err = s.trie.TryUpdate(key, val)
	if err != nil {
		return fmt.Errorf("could not put validators: %w", err)
	}

	return nil
}

// GetValidators gets the active validators for the current validator epoch.
func (s *State) GetValidators() ([]*Validator, error) {

	key := []byte{codeValidators}

	val, err := s.trie.TryGet(key)
	if err != nil {
		return nil, fmt.Errorf("could not get validators: %w", err)
	}

	var validators []*Validator
	err = s.dec.Unmarshal(val, &validators)
	if err != nil {
		return nil, fmt.Errorf("could not decode validators: %w", err)
	}

	return validators, nil
}

// RootHash calculates the root hash of the validator state. It is used to link the
// validator state to the EVM state and thus make it change in lock-step with potential
// forks in the chain history.
func (s *State) RootHash() (common.Hash, error) {

	root, _, err := s.trie.Commit(nil)
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not commit trie: %w", err)
	}

	return root, nil
}

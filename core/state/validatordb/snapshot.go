package validatordb

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/ids"

	"github.com/flare-foundation/coreth/core/state"
)

type Snapshot struct {
	state *State
	trie  state.Trie
}

func (s *Snapshot) SetEpoch(epoch uint64) error {

	key := []byte{codeEpoch}

	val := make([]byte, 8)
	binary.BigEndian.PutUint64(val, epoch)

	err := s.trie.TryUpdate(key, val)
	if err != nil {
		return fmt.Errorf("could not update epoch: %w", err)
	}

	return nil
}

func (s *Snapshot) GetEpoch() (uint64, error) {

	key := []byte{codeEpoch}

	val, err := s.trie.TryGet(key)
	if err != nil {
		return 0, fmt.Errorf("could not get epoch: %w", err)
	}

	return binary.BigEndian.Uint64(val), nil
}

func (s *Snapshot) SetMapping(provider common.Address, nodeID ids.ShortID) error {

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

func (s *Snapshot) GetMapping(provider common.Address) (ids.ShortID, error) {

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

func (s *Snapshot) AllMappings() (map[common.Address]ids.ShortID, error) {

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

func (s *Snapshot) SetCandidates(candidates []*Candidate) error {

	key := []byte{codeCandidates}

	val, err := s.state.enc.Marshal(candidates)
	if err != nil {
		return fmt.Errorf("could not encode candidates: %w", err)
	}

	err = s.trie.TryUpdate(key, val)
	if err != nil {
		return fmt.Errorf("could not put candidates: %w", err)
	}

	return nil
}

func (s *Snapshot) GetCandidates() ([]*Candidate, error) {

	key := []byte{codeCandidates}

	val, err := s.trie.TryGet(key)
	if err != nil {
		return nil, fmt.Errorf("could not get candidates: %w", err)
	}

	var candidates []*Candidate
	err = s.state.dec.Unmarshal(val, &candidates)
	if err != nil {
		return nil, fmt.Errorf("could not decode candidates: %w", err)
	}

	return candidates, nil
}

func (s *Snapshot) SetValidators(validators []*Validator) error {

	key := []byte{codeValidators}

	val, err := s.state.enc.Marshal(validators)
	if err != nil {
		return fmt.Errorf("could not encode validators: %w", err)
	}

	err = s.trie.TryUpdate(key, val)
	if err != nil {
		return fmt.Errorf("could not put validators: %w", err)
	}

	return nil
}

func (s *Snapshot) GetValidators() ([]*Validator, error) {

	key := []byte{codeValidators}

	val, err := s.trie.TryGet(key)
	if err != nil {
		return nil, fmt.Errorf("could not get validators: %w", err)
	}

	var validators []*Validator
	err = s.state.dec.Unmarshal(val, &validators)
	if err != nil {
		return nil, fmt.Errorf("could not decode validators: %w", err)
	}

	return validators, nil
}

func (s *Snapshot) RootHash() (common.Hash, error) {

	root, _, err := s.trie.Commit(nil)
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not commit trie: %w", err)
	}

	return root, nil
}

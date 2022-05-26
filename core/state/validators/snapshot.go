package validators

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

	key := make([]byte, 1)
	key[0] = codeEpoch

	val := make([]byte, 8)
	binary.BigEndian.PutUint64(val, epoch)

	err := s.trie.TryUpdate(key, val)
	if err != nil {
		return fmt.Errorf("could not update epoch: %w", err)
	}

	return nil
}

func (s *Snapshot) GetEpoch() (uint64, error) {

	key := make([]byte, 1)
	key[0] = codeEpoch

	val, err := s.trie.TryGet(key)
	if err != nil {
		return 0, fmt.Errorf("could not get epoch: %w", err)
	}

	return binary.BigEndian.Uint64(val), nil
}

func (s *Snapshot) SetEntries(entries []Entry) error {

	key := make([]byte, 1)
	key[0] = codeEntries

	val, err := s.state.enc.Marshal(entries)
	if err != nil {
		return fmt.Errorf("could not encode entries: %w", err)
	}

	err = s.trie.TryUpdate(key, val)
	if err != nil {
		return fmt.Errorf("could not put entries: %w", err)
	}

	return nil
}

func (s *Snapshot) GetEntries() ([]Entry, error) {

	key := make([]byte, 1)
	key[0] = codeEntries

	val, err := s.trie.TryGet(key)
	if err != nil {
		return nil, fmt.Errorf("could not get entries: %w", err)
	}

	var entries []Entry
	err = s.state.dec.Unmarshal(val, &entries)
	if err != nil {
		return nil, fmt.Errorf("could not decode entries: %w", err)
	}

	return entries, nil
}

func (s *Snapshot) SetPending(provider common.Address, nodeID ids.ShortID) error {
	return s.setValidator(codePending, provider, nodeID)
}

func (s *Snapshot) OnePending(provider common.Address) (ids.ShortID, error) {
	return s.oneValidator(codePending, provider)
}

func (s *Snapshot) AllPending() (map[common.Address]ids.ShortID, error) {
	return s.allValidators(codePending)
}

func (s *Snapshot) LookupPending(nodeID ids.ShortID) (common.Address, error) {
	return s.lookupValidator(codePending, nodeID)
}

func (s *Snapshot) UnsetPending(provider common.Address) error {
	return s.unsetValidator(codePending, provider)
}

func (s *Snapshot) DropPending() error {
	return s.dropValidators(codePending)
}

func (s *Snapshot) SetActive(provider common.Address, nodeID ids.ShortID) error {
	return s.setValidator(codeActive, provider, nodeID)
}

func (s *Snapshot) GetActive(provider common.Address) (ids.ShortID, error) {
	return s.oneValidator(codeActive, provider)
}

func (s *Snapshot) AllActive() (map[common.Address]ids.ShortID, error) {
	return s.allValidators(codeActive)
}

func (s *Snapshot) LookupActive(nodeID ids.ShortID) (common.Address, error) {
	return s.lookupValidator(codeActive, nodeID)
}

func (s *Snapshot) UnsetActive(provider common.Address) error {
	return s.unsetValidator(codeActive, provider)
}

func (s *Snapshot) DropActive() error {
	return s.dropValidators(codeActive)
}

func (s *Snapshot) setValidator(code byte, provider common.Address, nodeID ids.ShortID) error {

	key := make([]byte, 21)
	key[0] = code
	copy(key[1:], provider[:])

	val := nodeID[:]

	err := s.trie.TryUpdate(key, val)
	if err != nil {
		return fmt.Errorf("could not put validator: %w", err)
	}

	return nil
}

func (s *Snapshot) oneValidator(code byte, provider common.Address) (ids.ShortID, error) {

	key := make([]byte, 21)
	key[0] = code
	copy(key[1:], provider[:])

	val, err := s.trie.TryGet(key)
	if err != nil {
		return ids.ShortID{}, fmt.Errorf("could not get validator: %w", err)
	}

	validatorID, err := ids.ToShortID(val)
	if err != nil {
		return ids.ShortID{}, fmt.Errorf("could not parse validator: %w", err)
	}

	return validatorID, nil

}

func (s *Snapshot) allValidators(code byte) (map[common.Address]ids.ShortID, error) {

	start := []byte{code}
	end := []byte{code + 1}

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

func (s *Snapshot) lookupValidator(code byte, nodeID ids.ShortID) (common.Address, error) {

	start := []byte{code}
	end := []byte{code + 1}

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
			return common.Address{}, fmt.Errorf("could not parse validator (val: %x): %w", val, err)
		}

		if validatorID != nodeID {
			continue
		}

		provider := common.BytesToAddress(key[1:])

		return provider, nil

	}

	err := it.Error()
	if err != nil {
		return common.Address{}, fmt.Errorf("could not iterate validators: %w", err)
	}

	return common.Address{}, ErrNotFound
}

func (s *Snapshot) unsetValidator(code byte, provider common.Address) error {

	key := make([]byte, 21)
	key[0] = code
	copy(key[1:], provider[:])

	err := s.trie.TryDelete(key)
	if err != nil {
		return fmt.Errorf("could not delete validator: %w", err)
	}

	return nil
}

func (s *Snapshot) dropValidators(code byte) error {

	key := make([]byte, 1)
	key[0] = code

	err := s.trie.TryDelete(key)
	if err != nil {
		return fmt.Errorf("could not drop validators: %w", err)
	}
	return nil
}

func (s *Snapshot) SetWeight(epoch uint64, nodeID ids.ShortID, weight uint64) error {

	key := make([]byte, 29)
	key[0] = codeWeight
	binary.BigEndian.PutUint64(key[1:9], epoch)
	copy(key[9:29], nodeID[:])

	val := make([]byte, 8)
	binary.BigEndian.PutUint64(val[:], weight)

	err := s.trie.TryUpdate(key, val)
	if err != nil {
		return fmt.Errorf("could not put weight: %w", err)
	}

	return nil
}

func (s *Snapshot) OneWeight(epoch uint64, nodeID ids.ShortID) (uint64, error) {

	key := make([]byte, 29)
	key[0] = codeWeight
	binary.BigEndian.PutUint64(key[1:9], epoch)
	copy(key[9:29], nodeID[:])

	val, err := s.trie.TryGet(key)
	if err != nil {
		return 0, fmt.Errorf("could not get weight: %w", err)
	}

	return binary.BigEndian.Uint64(val), nil
}

func (s *Snapshot) AllWeights(epoch uint64) (map[ids.ShortID]uint64, error) {

	start := make([]byte, 9)
	start[0] = codeWeight
	binary.BigEndian.PutUint64(start[1:], epoch)

	end := make([]byte, 9)
	end[0] = codeWeight
	binary.BigEndian.PutUint64(end[1:], epoch+1)

	it := s.trie.NodeIterator(start)
	weights := make(map[ids.ShortID]uint64)
	for it.Next(true) {

		if !it.Leaf() {
			continue
		}

		key := it.LeafKey()
		if bytes.Compare(key, end) >= 0 {
			break
		}

		val := it.LeafBlob()
		weight := binary.BigEndian.Uint64(val)

		validatorID, err := ids.ToShortID(key[9:])
		if err != nil {
			return nil, fmt.Errorf("could not parse validator (key: %x): %w", key, err)
		}

		weights[validatorID] = weight
	}

	err := it.Error()
	if err != nil {
		return nil, fmt.Errorf("could not iterate weights: %w", err)
	}

	return weights, nil
}

func (s *Snapshot) RootHash() (common.Hash, error) {

	root, _, err := s.trie.Commit(nil)
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not commit trie: %w", err)
	}

	return root, nil
}

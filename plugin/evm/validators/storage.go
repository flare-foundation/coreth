package validators

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"

	"github.com/flare-foundation/flare/database"
	"github.com/flare-foundation/flare/ids"
)

var (
	epochKey          = []byte("epoch")
	activePrefix      = []byte("active")
	pendingPrefix     = []byte("pending")
	weightEpochPrefix = []byte("weightepoch")

	errNoEntries = errors.New("no entries")
)

type Storage struct {
	database database.Database
	enc      cbor.EncMode
	dec      cbor.DecMode
}

func NewStorage(database database.Database) *Storage {

	enc, err := cbor.EncOptions{
		Sort:        cbor.SortCoreDeterministic,
		IndefLength: cbor.IndefLengthForbidden,
		TagsMd:      cbor.TagsForbidden,
	}.EncMode()
	if err != nil {
		panic(fmt.Sprintf("invalid encoding options (%s)", err))
	}

	dec, err := cbor.DecOptions{
		DupMapKey:         cbor.DupMapKeyEnforcedAPF,
		IndefLength:       cbor.IndefLengthAllowed,
		TagsMd:            cbor.TagsForbidden,
		ExtraReturnErrors: cbor.ExtraDecErrorUnknownField,
	}.DecMode()
	if err != nil {
		panic(fmt.Sprintf("invalid decoding option (%s)", err))
	}

	s := Storage{
		database: database,
		enc:      enc,
		dec:      dec,
	}

	return &s
}

// SetEntries sets list of entries for the epoch
func (s *Storage) SetEntries(epoch uint64, entries []Entry) error {

	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, epoch)

	data, err := s.enc.Marshal(entries)
	if err != nil {
		return fmt.Errorf("could not encode entries: %w", err)
	}

	err = s.database.Put(key, data)
	if err != nil {
		return fmt.Errorf("could not put etries data: %w", err)
	}

	return nil
}

func (s *Storage) GetEntries(epoch uint64) ([]Entry, error) {

	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, epoch)

	data, err := s.database.Get(key)
	if err != nil {
		return nil, fmt.Errorf("could not get entries: %w", err)
	}

	var entries []Entry
	err = s.dec.Unmarshal(data, &entries)
	if err != nil {
		return nil, fmt.Errorf("could not decode entries: %w", err)
	}

	if len(entries) == 0 {
		return nil, errNoEntries
	}

	return entries, nil
}

func (s *Storage) SetPending(provider common.Address, nodeID ids.ShortID) error {
	err := s.database.Put(bytes.Join([][]byte{pendingPrefix, provider.Bytes()}, nil), nodeID.Bytes())
	if err != nil {
		return fmt.Errorf("could not set pending provider %s: %w", provider, err)
	}
	return nil
}

func (s *Storage) GetPending(provider common.Address) (ids.ShortID, error) {
	data, err := s.database.Get(bytes.Join([][]byte{pendingPrefix, provider.Bytes()}, nil))
	if err != nil {
		return ids.ShortID{}, fmt.Errorf("could not get pending provider %s: %w", provider, err)
	}
	return ids.ToShortID(data)
}

func (s *Storage) SetActive(provider common.Address, nodeID ids.ShortID) error {
	err := s.database.Put(bytes.Join([][]byte{activePrefix, provider.Bytes()}, nil), nodeID.Bytes())
	if err != nil {
		return fmt.Errorf("could not set active provider %s: %w", provider, err)
	}
	return nil
}

func (s *Storage) GetActive(provider common.Address) (ids.ShortID, error) {
	data, err := s.database.Get(bytes.Join([][]byte{activePrefix, provider.Bytes()}, nil))
	if err != nil {
		return ids.ShortID{}, fmt.Errorf("could not get active provider %s: %w", provider, err)
	}
	return ids.ToShortID(data)
}

func (s *Storage) Epoch() (uint64, error) {

	data, err := s.database.Get(epochKey)
	if err != nil {
		return 0, fmt.Errorf("could not get epoch: %w", err)
	}

	return big.NewInt(0).SetBytes(data).Uint64(), nil
}

func (s *Storage) SetEpoch(epoch uint64) error {
	err := s.database.Put(epochKey, big.NewInt(0).SetUint64(epoch).Bytes())
	if err != nil {
		return fmt.Errorf("could not set epoch: %w", err)
	}

	return nil
}

func (s *Storage) Pending() (map[common.Address]ids.ShortID, error) {
	it := s.database.NewIteratorWithPrefix(pendingPrefix)
	defer it.Release()

	m := make(map[common.Address]ids.ShortID)

	for it.Next() {
		id, _ := ids.ToShortID(it.Value())
		m[common.BytesToAddress(bytes.TrimPrefix(it.Key(), pendingPrefix))] = id
	}
	if err := it.Error(); err != nil {
		return nil, fmt.Errorf("could not get list of pending providers: %w", err)
	}

	return m, nil
}

func (s *Storage) Active() (map[common.Address]ids.ShortID, error) {
	it := s.database.NewIteratorWithPrefix(activePrefix)
	defer it.Release()

	m := make(map[common.Address]ids.ShortID)

	for it.Next() {
		id, _ := ids.ToShortID(it.Value())
		m[common.BytesToAddress(bytes.TrimPrefix(it.Key(), activePrefix))] = id
	}
	if err := it.Error(); err != nil {
		return nil, fmt.Errorf("could not get list of active providers: %w", err)
	}

	return m, nil
}

func (s *Storage) Weights(epoch uint64) (map[ids.ShortID]uint64, error) {
	prefix := bytes.Join([][]byte{weightEpochPrefix, big.NewInt(0).SetUint64(epoch).Bytes()}, nil)

	it := s.database.NewIteratorWithPrefix(prefix)
	defer it.Release()

	m := make(map[ids.ShortID]uint64)

	for it.Next() {
		weight := big.NewInt(0).SetBytes(it.Value())
		key := it.Key()
		id := key[len(key)-len(ids.ShortID{}):]

		nodeID, err := ids.ToShortID(id)
		if err != nil {
			return nil, fmt.Errorf("could not parse short id %v: %w", id, err)
		}

		m[nodeID] = weight.Uint64()
	}
	if err := it.Error(); err != nil {
		return nil, fmt.Errorf("could not get list of weights for epoch %d: %w", epoch, err)
	}

	return m, nil
}

func (s *Storage) SetWeight(epoch uint64, nodeID ids.ShortID, weight uint64) error {

	key := bytes.Join([][]byte{weightEpochPrefix, big.NewInt(0).SetUint64(epoch).Bytes(), nodeID.Bytes()}, nil)

	err := s.database.Put(key, big.NewInt(0).SetUint64(weight).Bytes())
	if err != nil {
		return fmt.Errorf("could not set weight: %w", err)
	}

	return nil
}

func (s *Storage) Lookup(provider common.Address) (ids.ShortID, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Storage) UnsetPending() error {
	it := s.database.NewIteratorWithPrefix(pendingPrefix)
	defer it.Release()

	for it.Next() {
		err := s.database.Delete(it.Key())
		if err != nil {
			return fmt.Errorf("could not delete pending provider %s: %w",
				common.BytesToAddress(bytes.TrimPrefix(it.Key(), pendingPrefix)), err)
		}
	}
	if err := it.Error(); err != nil {
		return fmt.Errorf("could not unset pending providers: %w", err)
	}

	return nil
}

func (s *Storage) UnsetActive(provider common.Address) error {
	err := s.database.Delete(bytes.Join([][]byte{activePrefix, provider.Bytes()}, nil))
	if err != nil {
		return fmt.Errorf("could not unset active provider %s: %w", provider, err)
	}
	return nil
}

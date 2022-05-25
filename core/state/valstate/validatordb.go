package valstate

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/fxamacker/cbor/v2"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/database"
	"github.com/flare-foundation/flare/ids"
)

var (
	epochKey          = []byte("epoch")
	activePrefix      = []byte("active")
	pendingPrefix     = []byte("pending")
	weightEpochPrefix = []byte("weightepoch")
)

type ValidatorDB struct {
	database database.Database
	enc      cbor.EncMode
	dec      cbor.DecMode
}

func NewValidatorDB(database database.Database) *ValidatorDB {

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

	v := ValidatorDB{
		database: database,
		enc:      enc,
		dec:      dec,
	}

	return &v
}

// SetEntries sets list of entries for the epoch
func (v *ValidatorDB) SetEntries(epoch uint64, entries []Entry) error {

	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, epoch)

	data, err := v.enc.Marshal(entries)
	if err != nil {
		return fmt.Errorf("could not encode entries: %w", err)
	}

	err = v.database.Put(key, data)
	if err != nil {
		return fmt.Errorf("could not put etries data: %w", err)
	}

	return nil
}

func (v *ValidatorDB) GetEntries(epoch uint64) ([]Entry, error) {

	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, epoch)

	data, err := v.database.Get(key)
	if err != nil {
		return nil, fmt.Errorf("could not get entries: %w", err)
	}

	var entries []Entry
	err = v.dec.Unmarshal(data, &entries)
	if err != nil {
		return nil, fmt.Errorf("could not decode entries: %w", err)
	}

	if len(entries) == 0 {
		return nil, ErrNoEntries
	}

	return entries, nil
}

func (v *ValidatorDB) SetPending(provider common.Address, nodeID ids.ShortID) error {
	err := v.database.Put(bytes.Join([][]byte{pendingPrefix, provider.Bytes()}, nil), nodeID.Bytes())
	if err != nil {
		return fmt.Errorf("could not set pending provider %s: %w", provider, err)
	}
	return nil
}

func (v *ValidatorDB) GetPending(provider common.Address) (ids.ShortID, error) {
	data, err := v.database.Get(bytes.Join([][]byte{pendingPrefix, provider.Bytes()}, nil))
	if err != nil {
		return ids.ShortID{}, fmt.Errorf("could not get pending provider %s: %w", provider, err)
	}
	return ids.ToShortID(data)
}

func (v *ValidatorDB) SetActive(provider common.Address, nodeID ids.ShortID) error {
	err := v.database.Put(bytes.Join([][]byte{activePrefix, provider.Bytes()}, nil), nodeID.Bytes())
	if err != nil {
		return fmt.Errorf("could not set active provider %s: %w", provider, err)
	}
	return nil
}

func (v *ValidatorDB) GetActive(provider common.Address) (ids.ShortID, error) {
	data, err := v.database.Get(bytes.Join([][]byte{activePrefix, provider.Bytes()}, nil))
	if err != nil {
		return ids.ShortID{}, fmt.Errorf("could not get active provider %s: %w", provider, err)
	}
	return ids.ToShortID(data)
}

func (v *ValidatorDB) Epoch() (uint64, error) {

	data, err := v.database.Get(epochKey)
	if err != nil {
		return 0, fmt.Errorf("could not get epoch: %w", err)
	}

	return big.NewInt(0).SetBytes(data).Uint64(), nil
}

func (v *ValidatorDB) SetEpoch(epoch uint64) error {
	err := v.database.Put(epochKey, big.NewInt(0).SetUint64(epoch).Bytes())
	if err != nil {
		return fmt.Errorf("could not set epoch: %w", err)
	}

	return nil
}

func (v *ValidatorDB) Pending() (map[common.Address]ids.ShortID, error) {
	it := v.database.NewIteratorWithPrefix(pendingPrefix)
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

func (v *ValidatorDB) Active() (map[common.Address]ids.ShortID, error) {
	it := v.database.NewIteratorWithPrefix(activePrefix)
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

func (v *ValidatorDB) Weights(epoch uint64) (map[ids.ShortID]uint64, error) {
	prefix := bytes.Join([][]byte{weightEpochPrefix, big.NewInt(0).SetUint64(epoch).Bytes()}, nil)

	it := v.database.NewIteratorWithPrefix(prefix)
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

func (v *ValidatorDB) SetWeight(epoch uint64, nodeID ids.ShortID, weight uint64) error {

	key := bytes.Join([][]byte{weightEpochPrefix, big.NewInt(0).SetUint64(epoch).Bytes(), nodeID.Bytes()}, nil)

	err := v.database.Put(key, big.NewInt(0).SetUint64(weight).Bytes())
	if err != nil {
		return fmt.Errorf("could not set weight: %w", err)
	}

	return nil
}

func (v *ValidatorDB) LookupPending(nodeID ids.ShortID) (common.Address, error) {
	return v.lookup(nodeID, pendingPrefix)
}

func (v *ValidatorDB) LookupActive(nodeID ids.ShortID) (common.Address, error) {
	return v.lookup(nodeID, activePrefix)
}

func (v *ValidatorDB) lookup(nodeID ids.ShortID, prefix []byte) (common.Address, error) {

	it := v.database.NewIteratorWithPrefix(prefix)
	defer it.Release()

	for it.Next() {
		id, _ := ids.ToShortID(it.Value())
		if id == nodeID {
			return common.BytesToAddress(bytes.TrimPrefix(it.Key(), pendingPrefix)), nil
		}
	}
	if err := it.Error(); err != nil {
		return common.Address{}, fmt.Errorf("could not get validator for node %v: %w", nodeID, err)
	}

	return common.Address{}, ErrNotFound
}

func (v *ValidatorDB) UnsetPending() error {
	it := v.database.NewIteratorWithPrefix(pendingPrefix)
	defer it.Release()

	for it.Next() {
		err := v.database.Delete(it.Key())
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

func (v *ValidatorDB) UnsetActive(provider common.Address) error {
	err := v.database.Delete(bytes.Join([][]byte{activePrefix, provider.Bytes()}, nil))
	if err != nil {
		return fmt.Errorf("could not unset active provider %s: %w", provider, err)
	}
	return nil
}

func (v *ValidatorDB) RootHash() (common.Hash, error) {
	return common.Hash{}, fmt.Errorf("not implemented")
}

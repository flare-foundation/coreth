package valstate

import (
	"encoding/binary"
	"fmt"

	"github.com/fxamacker/cbor/v2"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/flare/database"
	"github.com/flare-foundation/flare/ids"
)

const (
	codeEpoch   = 0
	codeEntries = 1
	codePending = 2
	codeActive  = 3
	codeWeight  = 4
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

func (v *ValidatorDB) SetEpoch(epoch uint64) error {

	key := make([]byte, 1)
	key[0] = codeEpoch

	val := make([]byte, 8)
	binary.BigEndian.PutUint64(val, epoch)

	err := v.database.Put(key, val)
	if err != nil {
		return fmt.Errorf("could not put epoch: %w", err)
	}

	return nil
}

func (v *ValidatorDB) GetEpoch() (uint64, error) {

	key := make([]byte, 1)
	key[0] = codeEpoch

	val, err := v.database.Get(key)
	if err != nil {
		return 0, fmt.Errorf("could not get epoch: %w", err)
	}

	return binary.BigEndian.Uint64(val), nil
}

func (v *ValidatorDB) SetEntries(entries []Entry) error {

	key := make([]byte, 1)
	key[0] = codeEntries

	val, err := v.enc.Marshal(entries)
	if err != nil {
		return fmt.Errorf("could not encode entries: %w", err)
	}

	err = v.database.Put(key, val)
	if err != nil {
		return fmt.Errorf("could not put entries: %w", err)
	}

	return nil
}

func (v *ValidatorDB) GetEntries() ([]Entry, error) {

	key := make([]byte, 1)
	key[0] = codeEntries

	val, err := v.database.Get(key)
	if err != nil {
		return nil, fmt.Errorf("could not get entries: %w", err)
	}

	var entries []Entry
	err = v.dec.Unmarshal(val, &entries)
	if err != nil {
		return nil, fmt.Errorf("could not decode entries: %w", err)
	}

	return entries, nil
}

func (v *ValidatorDB) SetPending(provider common.Address, nodeID ids.ShortID) error {
	return v.setValidator(codePending, provider, nodeID)
}

func (v *ValidatorDB) OnePending(provider common.Address) (ids.ShortID, error) {
	return v.oneValidator(codePending, provider)
}

func (v *ValidatorDB) AllPending() (map[common.Address]ids.ShortID, error) {
	return v.allValidators(codePending)
}

func (v *ValidatorDB) LookupPending(nodeID ids.ShortID) (common.Address, error) {
	return v.lookupValidator(codePending, nodeID)
}

func (v *ValidatorDB) UnsetPending(provider common.Address) error {
	return v.unsetValidator(codePending, provider)
}

func (v *ValidatorDB) DropPending() error {
	return v.dropValidators(codePending)
}

func (v *ValidatorDB) SetActive(provider common.Address, nodeID ids.ShortID) error {
	return v.setValidator(codeActive, provider, nodeID)
}

func (v *ValidatorDB) GetActive(provider common.Address) (ids.ShortID, error) {
	return v.oneValidator(codeActive, provider)
}

func (v *ValidatorDB) AllActive() (map[common.Address]ids.ShortID, error) {
	return v.allValidators(codeActive)
}

func (v *ValidatorDB) LookupActive(nodeID ids.ShortID) (common.Address, error) {
	return v.lookupValidator(codeActive, nodeID)
}

func (v *ValidatorDB) UnsetActive(provider common.Address) error {
	return v.unsetValidator(codeActive, provider)
}

func (v *ValidatorDB) DropActive() error {
	return v.dropValidators(codeActive)
}

func (v *ValidatorDB) setValidator(code byte, provider common.Address, nodeID ids.ShortID) error {

	key := make([]byte, 21)
	key[0] = code
	copy(key[1:], provider[:])

	val := nodeID[:]

	err := v.database.Put(key, val)
	if err != nil {
		return fmt.Errorf("could not put validator: %w", err)
	}

	return nil
}

func (v *ValidatorDB) oneValidator(code byte, provider common.Address) (ids.ShortID, error) {

	key := make([]byte, 21)
	key[0] = code
	copy(key[1:], provider[:])

	val, err := v.database.Get(key)
	if err != nil {
		return ids.ShortID{}, fmt.Errorf("could not get validator: %w", err)
	}

	validatorID, err := ids.ToShortID(val)
	if err != nil {
		return ids.ShortID{}, fmt.Errorf("could not parse validator: %w", err)
	}

	return validatorID, nil

}

func (v *ValidatorDB) allValidators(code byte) (map[common.Address]ids.ShortID, error) {

	prefix := []byte{code}

	it := v.database.NewIteratorWithPrefix(prefix)
	defer it.Release()

	validators := make(map[common.Address]ids.ShortID)
	for it.Next() {

		val := it.Value()
		validatorID, err := ids.ToShortID(val)
		if err != nil {
			return nil, fmt.Errorf("could not parse validator (val: %x): %w", val, err)
		}

		key := it.Key()
		provider := common.BytesToAddress(key[1:])

		validators[provider] = validatorID
	}

	err := it.Error()
	if err != nil {
		return nil, fmt.Errorf("could not iterate validators: %w", err)
	}

	return validators, nil
}

func (v *ValidatorDB) lookupValidator(code byte, nodeID ids.ShortID) (common.Address, error) {

	prefix := make([]byte, 1)
	prefix[0] = code

	it := v.database.NewIteratorWithPrefix(prefix)
	defer it.Release()

	for it.Next() {

		val := it.Value()
		validatorID, err := ids.ToShortID(it.Value())
		if err != nil {
			return common.Address{}, fmt.Errorf("could not parse validator (val: %x): %w", val, err)
		}

		if validatorID != nodeID {
			continue
		}

		key := it.Key()
		provider := common.BytesToAddress(key[1:])

		return provider, nil

	}

	err := it.Error()
	if err != nil {
		return common.Address{}, fmt.Errorf("could not iterate validators: %w", err)
	}

	return common.Address{}, ErrNotFound
}

func (v *ValidatorDB) unsetValidator(code byte, provider common.Address) error {

	key := make([]byte, 21)
	key[0] = code
	copy(key[1:], provider[:])

	err := v.database.Delete(key)
	if err != nil {
		return fmt.Errorf("could not delete validator: %w", err)
	}

	return nil
}

func (v *ValidatorDB) dropValidators(code byte) error {

	prefix := make([]byte, 1)
	prefix[0] = code

	it := v.database.NewIteratorWithPrefix(prefix)
	defer it.Release()

	for it.Next() {
		key := it.Key()
		err := v.database.Delete(key)
		if err != nil {
			return fmt.Errorf("could not delete validator (key: %x): %w", key, err)
		}
	}

	err := it.Error()
	if err != nil {
		return fmt.Errorf("could not iterate validators: %w", err)
	}

	return nil
}

func (v *ValidatorDB) SetWeight(epoch uint64, nodeID ids.ShortID, weight uint64) error {

	key := make([]byte, 29)
	key[0] = codeWeight
	binary.BigEndian.PutUint64(key[1:9], epoch)
	copy(key[9:29], nodeID[:])

	val := make([]byte, 8)
	binary.BigEndian.PutUint64(val[:], weight)

	err := v.database.Put(key, val)
	if err != nil {
		return fmt.Errorf("could not put weight: %w", err)
	}

	return nil
}

func (v *ValidatorDB) OneWeight(epoch uint64, nodeID ids.ShortID) (uint64, error) {

	key := make([]byte, 29)
	key[0] = codeWeight
	binary.BigEndian.PutUint64(key[1:9], epoch)
	copy(key[9:29], nodeID[:])

	val, err := v.database.Get(key)
	if err != nil {
		return 0, fmt.Errorf("could not get weight: %w", err)
	}

	return binary.BigEndian.Uint64(val), nil
}

func (v *ValidatorDB) AllWeights(epoch uint64) (map[ids.ShortID]uint64, error) {

	prefix := make([]byte, 9)
	prefix[0] = codeWeight
	binary.BigEndian.PutUint64(prefix[1:], epoch)

	it := v.database.NewIteratorWithPrefix(prefix)
	defer it.Release()

	weights := make(map[ids.ShortID]uint64)
	for it.Next() {

		val := it.Value()
		weight := binary.BigEndian.Uint64(val)

		key := it.Key()
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

func (v *ValidatorDB) RootHash() (common.Hash, error) {
	return common.Hash{}, fmt.Errorf("not implemented")
}

// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"encoding/binary"
	"fmt"

	"github.com/fxamacker/cbor/v2"

	"github.com/flare-foundation/coreth/ethdb"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"
)

type ValidatorsStore struct {
	log logging.Logger
	db  ethdb.Database
	enc cbor.EncMode
	dec cbor.DecMode
}

func NewValidatorsStore(log logging.Logger, db ethdb.Database) (*ValidatorsStore, error) {

	enc, err := cbor.EncOptions{
		Sort:        cbor.SortCoreDeterministic,
		IndefLength: cbor.IndefLengthForbidden,
		TagsMd:      cbor.TagsForbidden,
	}.EncMode()
	if err != nil {
		return nil, fmt.Errorf("invalid encoding options: %w", err)
	}

	dec, err := cbor.DecOptions{
		DupMapKey:         cbor.DupMapKeyEnforcedAPF,
		IndefLength:       cbor.IndefLengthAllowed,
		TagsMd:            cbor.TagsForbidden,
		ExtraReturnErrors: cbor.ExtraDecErrorUnknownField,
	}.DecMode()
	if err != nil {
		return nil, fmt.Errorf("invalid decoding options: %w", err)
	}

	v := ValidatorsStore{
		log: log,
		db:  db,
		enc: enc,
		dec: dec,
	}

	return &v, nil
}

func (v *ValidatorsStore) Persist(epoch uint64, validators map[ids.ShortID]uint64) error {

	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, epoch)

	data, err := v.enc.Marshal(validators)
	if err != nil {
		return fmt.Errorf("could not encode validators: %w", err)
	}

	err = v.db.Put(key, data)
	if err != nil {
		return fmt.Errorf("could not put validator data: %w", err)
	}

	v.log.Debug("persisted validators for epoch %d", epoch)

	return nil
}

func (v *ValidatorsStore) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {

	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, epoch)

	data, err := v.db.Get(key)
	if err != nil {
		return nil, fmt.Errorf("could not get validator data: %w", err)
	}

	validators := make(map[ids.ShortID]uint64)
	err = v.dec.Unmarshal(data, &validators)
	if err != nil {
		return nil, fmt.Errorf("could not decode validators: %w", err)
	}

	v.log.Debug("restored validators for epoch %d", epoch)

	return validators, nil
}

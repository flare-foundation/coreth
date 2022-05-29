package validatordb

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/coreth/core/state"
	"github.com/flare-foundation/coreth/ethdb"
	"github.com/flare-foundation/coreth/trie"
)

type Repository struct {
	db  state.Database
	enc cbor.EncMode
	dec cbor.DecMode
}

func NewRepository(diskdb ethdb.KeyValueStore, options ...Option) *Repository {

	cfg := DefaultConfig
	for _, option := range options {
		option(&cfg)
	}

	db := state.NewDatabaseWithConfig(diskdb,
		&trie.Config{
			Cache:     cfg.CacheSize,
			Preimages: false,
		},
	)

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

	v := Repository{
		db:  db,
		enc: enc,
		dec: dec,
	}

	return &v
}

func (r *Repository) WithRoot(root common.Hash) (*State, error) {

	trie, err := r.db.OpenTrie(root)
	if err != nil {
		return nil, fmt.Errorf("could not open trie: %w", err)
	}

	state := State{
		enc:  r.enc,
		dec:  r.dec,
		trie: trie,
	}

	return &state, nil
}

package validators

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/coreth/core/state"
	"github.com/flare-foundation/coreth/ethdb"
	"github.com/flare-foundation/coreth/trie"
)

const (
	codeEpoch   = 0
	codeEntries = 1
	codePending = 2
	codeActive  = 3
	codeWeight  = 4
)

type State struct {
	db  state.Database
	enc cbor.EncMode
	dec cbor.DecMode
}

func NewState(diskdb ethdb.KeyValueStore, options ...Option) *State {

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

	v := State{
		db:  db,
		enc: enc,
		dec: dec,
	}

	return &v
}

func (s *State) WithRoot(root common.Hash) (*Snapshot, error) {

	trie, err := s.db.OpenTrie(root)
	if err != nil {
		return nil, fmt.Errorf("could not open trie: %w", err)
	}

	snapshot := Snapshot{
		state: s,
		trie:  trie,
	}

	return &snapshot, nil
}

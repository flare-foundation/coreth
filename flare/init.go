package flare

import (
	"github.com/dgraph-io/badger/v3"
	bstore "gitlab.com/flarenetwork/coreth/flare/store/badger"
)

const (
	DatabaseDirectory = "cache"
)

func init() {
	opts := badger.DefaultOptions(DatabaseDirectory)
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}
	store = bstore.NewStore(db)
	connectors = make(map[uint32]Connector)
}

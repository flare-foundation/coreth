package validators

import (
	"github.com/flare-foundation/flare/database"

	"github.com/flare-foundation/coreth/core/vm"
)

type Storage struct {
	db database.Database
}

func NewStorage(db database.Database) *Storage {

	s := Storage{
		db: db,
	}

	return &s
}

func (s *Storage) WithEVM(evm *vm.EVM) vm.ValidatorManager {

	m := Manager{
		db:  s.db,
		evm: evm,
	}

	return &m
}

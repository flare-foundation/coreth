package validators

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/coreth/core/vm"
	"github.com/flare-foundation/flare/database"
	"github.com/flare-foundation/flare/ids"
)

var (
	errNoPriceSubmitter = errors.New("no price submitter")
	errFTSONotDeployed  = errors.New("FTSO not deployed")
	errFTSONotActive    = errors.New("FTSO not active")

	epochKey = []byte("epoch")
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

func (s *Storage) WithEVM(evm *vm.EVM) (vm.ValidatorManager, error) {

	ftso, err := NewFTSO(evm)
	if err != nil {
		return nil, fmt.Errorf("could not create FTSO: %w", err)
	}

	m := Manager{
		log:  nil,
		repo: s,
		ftso: ftso,
	}

	return &m, nil
}

func (s *Storage) Epoch() (uint64, error) {
	e, err := s.db.Get(epochKey)
	if err != nil {
		return 0, fmt.Errorf("could not get epoch: %w", err)
	}

	epochInt := big.NewInt(0).SetBytes(e)

	return epochInt.Uint64(), nil
}

func (s *Storage) Pending() (map[common.Address]ids.ShortID, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Storage) Active() (map[common.Address]ids.ShortID, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Storage) Weights(epoch uint64) (map[ids.ShortID]uint64, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Storage) Lookup(provider common.Address) (ids.ShortID, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Storage) SetPending(provider common.Address, nodeID ids.ShortID) error {
	// TODO implement me
	panic("implement me")
}

func (s *Storage) SetEpoch(epoch uint64) error {
	epochData := big.NewInt(0).SetUint64(epoch).Bytes()

	err := s.db.Put(epochKey, epochData)
	if err != nil {
		return fmt.Errorf("could not set epoch: %w", err)
	}

	return nil
}

func (s *Storage) SetActive(provider common.Address, nodeID ids.ShortID) error {
	// TODO implement me
	panic("implement me")
}

func (s *Storage) SetWeight(epoch uint64, nodeID ids.ShortID, weight uint64) error {
	// TODO implement me
	panic("implement me")
}

func (s *Storage) UnsetPending() error {
	// TODO implement me
	panic("implement me")
}

func (s *Storage) UnsetActive(address common.Address) error {
	// TODO implement me
	panic("implement me")
}

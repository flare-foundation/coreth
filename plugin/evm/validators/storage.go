package validators

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"

	"github.com/flare-foundation/coreth/ethdb"
	"github.com/flare-foundation/flare/ids"
)

var (
	errNoPriceSubmitter = errors.New("no price submitter")
	errFTSONotDeployed  = errors.New("FTSO not deployed")
	errFTSONotActive    = errors.New("FTSO not active")

	epochKey = []byte("epoch")
)

type Storage struct {
	read  ethdb.Reader
	write ethdb.Writer
	enc   cbor.EncMode
	dec   cbor.DecMode
}

func NewStorage(read ethdb.Reader, write ethdb.Writer) *Storage {

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
		read:  read,
		write: write,
		enc:   enc,
		dec:   dec,
	}

	return &s
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

	err := s.write.Put(epochKey, epochData)
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

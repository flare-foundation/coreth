package validators

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/flare/ids"
	"github.com/fxamacker/cbor/v2"

	"github.com/flare-foundation/coreth/ethdb"
)

var (
	epochKey = []byte("epoch")

	errNoEntries = errors.New("no entries")
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

// SetEntries sets list of entries for the epoch
func (s *Storage) SetEntries(epoch uint64, entries []Entry) error {

	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, epoch)

	data, err := s.enc.Marshal(entries)
	if err != nil {
		return fmt.Errorf("could not encode entries: %w", err)
	}

	err = s.write.Put(key, data)
	if err != nil {
		return fmt.Errorf("could not put etries data: %w", err)
	}

	return nil
}

func (s *Storage) GetEntries(epoch uint64) ([]Entry, error) {

	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, epoch)

	data, err := s.read.Get(key)
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

	// TODO implement me
	panic("implement me")
}

func (s *Storage) GetPending(provider common.Address) (ids.ShortID, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Storage) SetActive(provider common.Address, nodeID ids.ShortID) error {
	// TODO implement me
	panic("implement me")
}

func (s *Storage) GetActive(provider common.Address) (ids.ShortID, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Storage) Epoch() (uint64, error) {

	data, err := s.read.Get(epochKey)
	if err != nil {
		return 0, fmt.Errorf("could not get epoch: %w", err)
	}

	return big.NewInt(0).SetBytes(data).Uint64(), nil
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

func (s *Storage) SetEpoch(epoch uint64) error {
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

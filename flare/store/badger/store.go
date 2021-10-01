package badger

import (
	"bytes"
	"fmt"

	"github.com/dgraph-io/badger/v3"
	"golang.org/x/crypto/sha3"
)

var (
	valAccepted = []byte{1}
	valRejected = []byte{0}
)

type Store struct {
	db *badger.DB
}

func NewStore(db *badger.DB) *Store {

	s := Store{
		db: db,
	}

	return &s
}

func (s *Store) Accept(data []byte, ret []byte) error {

	key := s.key(data, ret)
	err := s.db.Update(func(tx *badger.Txn) error {
		err := tx.Set(key, valAccepted)
		if err != nil {
			return fmt.Errorf("could not set value: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("could not save accepeted: %w", err)
	}

	return nil
}

func (s *Store) Reject(data []byte, ret []byte) error {

	key := s.key(data, ret)
	err := s.db.Update(func(tx *badger.Txn) error {
		err := tx.Set(key, valRejected)
		if err != nil {
			return fmt.Errorf("could not set value: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("could not save rejected: %w", err)
	}

	return nil
}

func (s *Store) Accepted(data []byte, ret []byte) (bool, error) {

	accepted := false
	key := s.key(data, ret)
	err := s.db.View(func(tx *badger.Txn) error {
		item, err := tx.Get(key)
		if err != nil {
			return fmt.Errorf("could not get item: %w", err)
		}
		err = item.Value(func(val []byte) error {
			accepted = bytes.Equal(val, valAccepted)
			return nil
		})
		if err != nil {
			return fmt.Errorf("could not get value: %w", err)
		}
		return nil
	})
	if err != nil {
		return false, fmt.Errorf("could not retrieve accepted: %w", err)
	}

	return accepted, nil
}

func (s *Store) Rejected(data []byte, ret []byte) (bool, error) {

	rejected := false
	key := s.key(data, ret)
	err := s.db.View(func(tx *badger.Txn) error {
		item, err := tx.Get(key)
		if err != nil {
			return fmt.Errorf("could not get item: %w", err)
		}
		err = item.Value(func(val []byte) error {
			rejected = bytes.Equal(val, valRejected)
			return nil
		})
		if err != nil {
			return fmt.Errorf("could not get value: %w", err)
		}
		return nil
	})
	if err != nil {
		return false, fmt.Errorf("could not retrieve rejected: %w", err)
	}

	return rejected, nil
}

func (s *Store) key(data []byte, ret []byte) []byte {
	hasher := sha3.New256()
	hasher.Write(data[0:4])
	hasher.Write(ret[0:64])
	hasher.Write(ret[96:128])
	return hasher.Sum(nil)
}

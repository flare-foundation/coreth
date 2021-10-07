package ripple

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

type Ledger struct {
	Hash      [32]byte
	Height    uint32
	Validated bool
}

func (l Ledger) Fingerprint() [32]byte {
	hash := hex.EncodeToString(l.Hash[:])
	fingerprint := crypto.Keccak256Hash([]byte(hash))
	return fingerprint
}

package ripple

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Transaction struct {
	Height      uint32
	Hash        [32]byte
	Validated   bool
	Amount      uint64
	Destination string
	Tag         uint32
}

func (t Transaction) Fingerprint(currency string) [32]byte {
	part1 := crypto.Keccak256(t.Hash[:])
	part2 := crypto.Keccak256([]byte(t.Destination))
	part3 := crypto.Keccak256(common.LeftPadBytes(common.FromHex(hexutil.EncodeUint64(uint64(t.Tag))), 32))
	part4 := crypto.Keccak256(common.LeftPadBytes(common.FromHex(hexutil.EncodeUint64(t.Amount)), 32))
	part5 := crypto.Keccak256([]byte(currency))
	return crypto.Keccak256Hash(part1, part2, part3, part4, part5)
}

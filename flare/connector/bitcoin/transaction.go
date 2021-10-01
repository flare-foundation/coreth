package bitcoin

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Transaction struct {
	Block         [32]byte
	Hash          [32]byte
	Index         uint8
	Recipient     string
	Amount        uint64
	Confirmations uint64
}

func (t Transaction) Fingerprint(currency string) [32]byte {
	part1 := crypto.Keccak256(append([]byte(strconv.Itoa(int(t.Index))), t.Hash[:]...))
	part2 := crypto.Keccak256([]byte(t.Recipient))
	part3 := crypto.Keccak256(common.LeftPadBytes(common.FromHex(hexutil.EncodeUint64(t.Amount)), 32))
	part4 := crypto.Keccak256([]byte(currency))
	return crypto.Keccak256Hash(part1, part2, part3, part4)
}

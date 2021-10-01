package bitcoin

type Block struct {
	Hash          [32]byte
	Height        uint64
	Confirmations uint64
}

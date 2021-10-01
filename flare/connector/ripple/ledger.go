package ripple

type Ledger struct {
	Hash      [32]byte
	Index     uint32
	Validated bool
}

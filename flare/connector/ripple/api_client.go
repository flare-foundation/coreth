package ripple

type APIClient interface {
	Ledger(index uint32) (*Ledger, error)
	Transaction(hash [32]byte) (*Transaction, error)
}

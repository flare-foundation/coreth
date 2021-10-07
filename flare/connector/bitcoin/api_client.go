package bitcoin

type APIClient interface {
	Block(hash [32]byte) (*Block, error)
	Transaction(hash [32]byte, index uint8) (*Transaction, error)
}

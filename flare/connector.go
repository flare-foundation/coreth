package flare

type Connector interface {
	ProveAvailability(ret []byte) (bool, error)
	ProvePayment(ret []byte) (bool, error)
	DisprovePayment(ret []byte) (bool, error)
}

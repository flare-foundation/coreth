package flare

type Store interface {
	Accept(data []byte, ret []byte) error
	Reject(data []byte, ret []byte) error
	Accepted(data []byte, ret []byte) (bool, error)
	Rejected(data []byte, ret []byte) (bool, error)
}

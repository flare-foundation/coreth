package bitcoin

type HeaderResponse struct {
	Hash          string `json:"hash"`
	Confirmations uint64 `json:"confirmations"`
	Height        uint64 `json:"height"`
	Timestamp     uint64 `json:"time"`
}

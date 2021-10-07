package bitcoin

type TransactionResponse struct {
	ID            string         `json:"txid"`
	Height        uint64         `json:"blockheight"`
	Block         string         `json:"blockhash"`
	Confirmations uint64         `json:"confirmations"`
	Timestamp     uint64         `json:"time"`
	Decoded       RawTransaction `json:"decoded"`
}

type RawTransaction struct {
	Outputs []Output `json:"vout"`
}

type Output struct {
	Value uint64 `json:"value"`
	Index uint16 `json:"n"`
	Key   Key    `json:"scriptPubKey"`
}

type Key struct {
	NumSignatures uint     `json:"reqSigs"`
	Type          string   `json:"type"`
	Addresses     []string `json:"addresses"`
}

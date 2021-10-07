package ripple

type TransactionRequest struct {
	Transaction string `json:"transaction"`
	Binary      bool   `json:"binary"`
}

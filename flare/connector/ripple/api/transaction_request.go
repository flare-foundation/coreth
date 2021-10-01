package api

type TransactionRequest struct {
	Transaction string `json:"transaction"`
	Binary      bool   `json:"binary"`
}

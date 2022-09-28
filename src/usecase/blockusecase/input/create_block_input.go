package input

type CreateBlockInput struct {
	Nonce            uint32 `json:"nonce"`
	PreviousHash     string `json:"previousHash"`
	TransactionsHash string `json:"transactionsHash"`
}

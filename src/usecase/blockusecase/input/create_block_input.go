package input

type CreateBlockInput struct {
	Nonce            uint32 `json:"nonce"`
	TransactionsHash string `json:"transactionsHash"`
}

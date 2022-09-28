package output

type CreateBlock struct {
	ID               uint32 `json:"id"`
	timestamp        uint64 `json:"timestamp"`
	nonce            uint64 `json:"nonce"`
	previousHash     uint32 `json:"previousHash"`
	transactionsHash uint32 `json:"transactionsHash"`
}

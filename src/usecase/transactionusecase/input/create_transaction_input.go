package input

type CreateTransactionInput struct {
	SenderAddress    string `json:"senderAddress"`
	RecipientAddress string `json:"recipientAddress"`
	Amount           uint64 `json:"amount"`
}

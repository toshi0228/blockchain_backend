package input

type CreateTransactionInput struct {
	SenderAddress    string `json:"senderAddress"`
	RecipientAddress string `json:"recipientAddress"`
	Amount           uint64 `json:"amount"`
	PrivateKey       string `json:"privateKey"`
	PublicKey        string `json:"publicKey"`
	Signature        string `json:"signature"`
}

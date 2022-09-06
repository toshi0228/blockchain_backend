package entity

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	senderBlockChainAddress    string
	recipientBlockChainAddress string
	value                      float32
}

// NewTransaction トランザクションの作成
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s \n", strings.Repeat("-", 40))
	fmt.Printf("送信者 : %s\n", t.senderBlockChainAddress)
	fmt.Printf("受取人 : %s\n", t.recipientBlockChainAddress)
	fmt.Printf("金額 : %.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"senderBlockChainAddress"`
		Recipient string  `json:"recipientBlockChainAddress"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockChainAddress,
		Recipient: t.recipientBlockChainAddress,
		Value:     t.value,
	})
}

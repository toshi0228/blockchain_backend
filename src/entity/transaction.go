package entity

import (
	"fmt"
	"github.com/toshi0228/blockchain/src/entity/vo"
	"strings"
)

type Transaction struct {
	id               vo.ID `json:"id"`
	senderAddress    string
	recipientAddress string
	value            uint64
	createdAt        vo.CreatedAt `json:"createdAt"`
	updatedAt        vo.UpdatedAt `json:"updatedAt"`
}

func (t *Transaction) Id() vo.ID {
	return t.id
}

func (t *Transaction) SenderAddress() string {
	return t.senderAddress
}

func (t *Transaction) RecipientAddress() string {
	return t.recipientAddress
}

func (t *Transaction) Value() uint64 {
	return t.value
}

func (t *Transaction) CreatedAt() vo.CreatedAt {
	return t.createdAt
}

func (t *Transaction) UpdatedAt() vo.UpdatedAt {
	return t.updatedAt
}

// NewTransaction トランザクションの作成
func NewTransaction(senderAddress string, recipientAddress string, value uint64) *Transaction {

	return &Transaction{
		senderAddress:    senderAddress,
		recipientAddress: recipientAddress,
		value:            value,
	}
}

func GenWhenCreateTransaction(senderAddress string, recipientAddress string, value uint64) (*Transaction, error) {
	return &Transaction{
		id:               vo.NewID(),
		senderAddress:    senderAddress,
		recipientAddress: recipientAddress,
		value:            value,
		createdAt:        vo.NewCreatedAt(),
		updatedAt:        vo.NewUpdatedAt(),
	}, nil
}

func (t *Transaction) Print() {
	fmt.Printf("%s \n", strings.Repeat("-", 40))
	fmt.Printf("送信者 : %s\n", t.senderAddress)
	fmt.Printf("受取人 : %s\n", t.recipientAddress)
	fmt.Printf("金額 : %v\n", t.value)
}

//func (t *Transaction) MarshalJSON() ([]byte, error) {
//	return json.Marshal(struct {
//		Sender    string `json:"senderBlockChainAddress"`
//		Recipient string `json:"recipientBlockChainAddress"`
//		Value     uint64 `json:"value"`
//	}{
//		Sender:    t.senderAddress,
//		Recipient: t.recipientAddress,
//		Value:     t.value,
//	})
//}

package entity

import (
	"encoding/json"
	"github.com/toshi0228/blockchain/src/entity/vo"
	"strings"
	"time"
)

type TransactionPool struct {
	id               vo.ID        `json:"id"`
	senderAddress    string       `json:"senderAddress"`
	recipientAddress string       `json:"recipientAddress"`
	amount           uint64       `json:"amount"`
	createdAt        vo.CreatedAt `json:"createdAt"`
	updatedAt        vo.UpdatedAt `json:"updatedAt"`
}

func (tp TransactionPool) Id() vo.ID {
	return tp.id
}

func (tp TransactionPool) SenderAddress() string {
	return tp.senderAddress
}

func (tp TransactionPool) RecipientAddress() string {
	return tp.recipientAddress
}

func (tp TransactionPool) Amount() uint64 {
	return tp.amount
}

func (tp TransactionPool) CreatedAt() vo.CreatedAt {
	return tp.createdAt
}

func (tp TransactionPool) UpdatedAt() vo.UpdatedAt {
	return tp.updatedAt
}

func NewTransactionPool(id uint32, senderAddress string, recipientAddress string, amount uint64, createdAt time.Time, updatedAt time.Time) (*TransactionPool, error) {

	return &TransactionPool{
		id:               vo.ID(id),
		senderAddress:    senderAddress,
		recipientAddress: recipientAddress,
		amount:           amount,
		createdAt:        vo.CreatedAt(createdAt),
		updatedAt:        vo.UpdatedAt(updatedAt),
	}, nil
}

func (tp *TransactionPool) ToJSON() string {

	byteArr, _ := json.Marshal(struct {
		SenderAddress    string    `json:"senderAddress"`
		RecipientAddress string    `json:"recipientAddress"`
		Amount           uint64    `json:"amount"`
		CreatedAt        time.Time `json:"createdAt"`
		UpdatedAt        time.Time `json:"updatedAt"`
	}{
		SenderAddress:    tp.senderAddress,
		RecipientAddress: tp.RecipientAddress(),
		Amount:           tp.amount,
		CreatedAt:        tp.createdAt.Value(),
		UpdatedAt:        tp.updatedAt.Value(),
	})
	return string(byteArr)
}

func TransactionPoolToHash(txPoolJSON []string) string {

	stTxPoolJSON := strings.Join(txPoolJSON, ",")
	return vo.NewHashToHex([]byte(stTxPoolJSON)).Value()

}

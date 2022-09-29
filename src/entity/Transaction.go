package entity

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/toshi0228/blockchain/src/entity/vo"
	"log"
)

type Transaction struct {
	id               vo.ID `json:"id"`
	senderAddress    string
	recipientAddress string
	value            uint64
	signature        *Signature
	senderPrivateKey *ecdsa.PrivateKey
	senderPublicKey  *ecdsa.PublicKey
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

func NewTransaction() (*Transaction, error) {
	return &Transaction{}, nil
}

// GenWhenCreateTransactions トランザクションの新規登録の作成
func GenWhenCreateTransactions(senderAddress, recipientAddress, senderPrivateKeyHex, senderPublicKeyHex, signatureHex string, value uint64) (*Transaction, error) {

	publicKey := vo.PublicKeyFromString(senderPublicKeyHex)

	// {"recipientAddress":"JdDfxqdCryQyArgL7nsorExmoPMJB9RzP","senderAddress":"1PRJPrEMiXre3bTuQMEHaRnZR2v5Z1anw3","amount":12}

	//TODO　改竄されていないかチェック
	type tx struct {
		RecipientAddress string `json:"recipientAddress"`
		SenderAddress    string `json:"senderAddress"`
		Value            uint64 `json:"amount"`
	}
	var x tx

	x.RecipientAddress = recipientAddress
	x.SenderAddress = senderAddress
	x.Value = value

	v, _ := json.Marshal(x)
	log.Println(string(v))

	hash := sha256.Sum256(v)
	signature := SignatureFromString(signatureHex)

	// 改竄して場合
	//hash = sha256.Sum256([]byte("HI"))

	//改竄されていないか検証
	valid := ecdsa.Verify(publicKey, hash[:], signature.R, signature.S)
	fmt.Println("signature verified:", valid)

	return &Transaction{
		id:               vo.NewID(),
		senderAddress:    senderAddress,
		recipientAddress: recipientAddress,
		value:            value,
		senderPrivateKey: vo.PrivateKeyFromString(senderPrivateKeyHex, publicKey),
		senderPublicKey:  publicKey,
		signature:        SignatureFromString(signatureHex),
		createdAt:        vo.NewCreatedAt(),
		updatedAt:        vo.NewUpdatedAt(),
	}, nil
}

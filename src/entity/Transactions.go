package entity

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/toshi0228/blockchain/src/entity/vo"
	"log"
	"strings"
)

type Transactions struct {
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

func (t *Transactions) Id() vo.ID {
	return t.id
}

func (t *Transactions) SenderAddress() string {
	return t.senderAddress
}

func (t *Transactions) RecipientAddress() string {
	return t.recipientAddress
}

func (t *Transactions) Value() uint64 {
	return t.value
}

func (t *Transactions) CreatedAt() vo.CreatedAt {
	return t.createdAt
}

func (t *Transactions) UpdatedAt() vo.UpdatedAt {
	return t.updatedAt
}

// NewTransactions トランザクションの作成
func NewTransactions(senderAddress string, recipientAddress string, value uint64) *Transaction {

	return &Transaction{
		senderAddress:    senderAddress,
		recipientAddress: recipientAddress,
		value:            value,
	}
}

func (t *Transactions) Print() {
	fmt.Printf("%s \n", strings.Repeat("-", 40))
	fmt.Printf("送信者 : %s\n", t.senderAddress)
	fmt.Printf("受取人 : %s\n", t.recipientAddress)
	fmt.Printf("金額 : %v\n", t.value)
}

// GenWhenCreateTransactions トランザクションの新規登録の作成
func GenWhenCreateTransactions(senderAddress, recipientAddress, senderPrivateKeyHex, senderPublicKeyHex, signatureHex string, value uint64) (*Transactions, error) {

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

	hash := sha256.Sum256([]byte(string(v)))

	signature := SignatureFromString(signatureHex)

	//改竄されていないか検証
	valid := ecdsa.Verify(publicKey, hash[:], signature.R, signature.S)
	fmt.Println("signature verified:", valid)

	return &Transactions{
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

package entity

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
)

type SentTransaction struct {
	senderPrivateKey *ecdsa.PrivateKey
	senderPublicKey  *ecdsa.PublicKey

	senderBlockChainAddress    string
	recipientBlockChainAddress string
	value                      float32
}

// NewSentTransaction 送信されたトランザクション　※送信者=>ブロックチェーンノード
func NewSentTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, sender string, recipient string, value float32) *SentTransaction {
	return &SentTransaction{privateKey, publicKey, sender, recipient, value}
}

func (t *SentTransaction) Print() {
	fmt.Printf("%s \n", strings.Repeat("-", 40))
	fmt.Printf("送信者 : %s\n", t.senderBlockChainAddress)
	fmt.Printf("受取人 : %s\n", t.recipientBlockChainAddress)
	fmt.Printf("金額 : %.1f\n", t.value)
}

// GenerateSignature トランザクションの署名　※送信者=>ブロックチェーンノード
func (t *SentTransaction) GenerateSignature() *Signature {
	m, _ := json.Marshal(t)
	h := sha256.Sum256(m)
	r, s, _ := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
	return NewSignature(r, s)
}

func (t *SentTransaction) MarshalJSON() ([]byte, error) {
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

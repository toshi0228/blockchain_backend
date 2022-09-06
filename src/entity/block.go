package entity

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

const (
	MiningDifficulty = 3
	MiningSender     = "MiningSender"
	MiningReward     = 1.0
)

type Block struct {
	timestamp    int64
	nonce        int
	previousHash [32]byte
	transactions []*Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transaction []*Transaction) *Block {
	b := &Block{}
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transaction
	return b
}

// Hash ブロックのハッシュ値を求める
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m)
}

// Print ブロックを見やすいように表示
func (b *Block) Print() {
	fmt.Printf("timestamp       %d\n", b.timestamp)
	fmt.Printf("nonce           %d\n", b.nonce)
	fmt.Printf("previous_hash   %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previousHash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

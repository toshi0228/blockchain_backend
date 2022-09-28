package entity

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/toshi0228/blockchain/src/entity/vo"
)

const (
	MiningDifficulty = 3
	MiningSender     = "MiningSender"
	MiningReward     = 1.0
)

type Block struct {
	id               uint32
	nonce            uint32
	previousHash     [32]byte
	transactionsHash [32]byte
	timestamp        int64
}

func (b *Block) Id() uint32 {
	return b.id
}

func (b *Block) Nonce() uint32 {
	return b.nonce
}

func (b *Block) PreviousHashToHex() string {

	return fmt.Sprintf("%x", b.previousHash)
}

func (b *Block) TransactionsHashToHex() string {
	return fmt.Sprintf("%x", b.transactionsHash)
}

func (b *Block) Timestamp() int64 {
	return b.timestamp
}

func NewBlock(nonce uint32, previousHash string, transactionsHash string) (*Block, error) {

	return &Block{
		id:               vo.NewID().Value(),
		nonce:            nonce,
		previousHash:     vo.NewStringData(previousHash).Hash(),
		transactionsHash: vo.NewStringData(transactionsHash).Hash(),
		timestamp:        vo.NewUnixTimeNow().Value(),
	}, nil
}

// Hash ブロックのハッシュ値を求める
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m)
}

// Print ブロックを見やすいように表示
//func (b *Block) Print() {
//	fmt.Printf("timestamp       %d\n", b.timestamp)
//	fmt.Printf("nonce           %d\n", b.nonce)
//	fmt.Printf("previous_hash   %x\n", b.previousHash)
//	for _, t := range b.transactions {
//		t.Print()
//	}
//}

//func (b *Block) MarshalJSON() ([]byte, error) {
//	return json.Marshal(struct {
//		Timestamp    int64          `json:"timestamp"`
//		Nonce        int            `json:"nonce"`
//		PreviousHash [32]byte       `json:"previousHash"`
//		Transactions []*Transaction `json:"transactions"`
//	}{
//		Timestamp:    b.timestamp,
//		Nonce:        b.nonce,
//		PreviousHash: b.previousHash,
//		Transactions: b.transactions,
//	})
//}

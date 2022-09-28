package entity

import (
	"encoding/json"
	"fmt"
	"github.com/toshi0228/blockchain/src/entity/vo"
)

type Block struct {
	id           uint32
	nonce        uint32
	previousHash [32]byte
	transactions string
	timestamp    int64
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
	return b.transactions
}

func (b *Block) Timestamp() int64 {
	return b.timestamp
}

func NewBlock(nonce uint32, previousHash string, transactions []string) (*Block, error) {

	return &Block{
		id:           vo.NewID().Value(),
		nonce:        nonce,
		previousHash: vo.NewHexHashToByte32(previousHash).Value(),
		transactions: vo.NewTransactions(transactions).Value(),
		timestamp:    vo.NewUnixTimeNow().Value(),
	}, nil
}

// Hash ブロックのハッシュ値を求める
func (b *Block) Hash() string {

	m, _ := json.Marshal(
		struct {
			Id               uint32   `json:"id"`
			Nonce            uint32   `json:"nonce"`
			PreviousHash     [32]byte `json:"previousHash"`
			TransactionsHash string   `json:"transactionsHash"`
			Timestamp        int64    `json:"timestamp"`
		}{
			Id:               b.id,
			Nonce:            b.nonce,
			PreviousHash:     b.previousHash,
			TransactionsHash: b.transactions,
			Timestamp:        b.timestamp,
		},
	)

	return vo.NewHash(m).ValueToHex()
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

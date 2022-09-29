package entity

import (
	"encoding/json"
	"fmt"
	"github.com/toshi0228/blockchain/src/entity/vo"
	"log"
	"strings"
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

func NewBlock(previousHash string, transactions []string) (*Block, error) {
	b := &Block{}

	//TODO ナンスの計算を行う

	b.id = vo.NewID().Value()
	b.previousHash = vo.NewHexHashToByte32(previousHash).Value()
	b.transactions = vo.NewTransactions(transactions).Value()
	b.timestamp = vo.NewUnixTimeNow().Value()

	nonce := b.ProofOfWork(b.id, previousHash, transactions, b.timestamp)
	b.nonce = nonce
	//return Block{
	//	id:           id,
	//	nonce:        nonce,
	//	previousHash: vo.NewHexHashToByte32(previousHash).Value(),
	//	transactions: vo.NewTransactions(transactions).Value(),
	//	timestamp:    timestamp,
	//}, nil
	//

	return b, nil
}

// Hash ブロックのハッシュ値を求める
//func (b *Block) Hash() string {
//
//	m, _ := json.Marshal(
//		struct {
//			Id           uint32   `json:"id"`
//			Nonce        uint32   `json:"nonce"`
//			PreviousHash [32]byte `json:"previousHash"`
//			Transactions string   `json:"transactionsHash"`
//			Timestamp    int64    `json:"timestamp"`
//		}{
//			Id:           b.id,
//			Nonce:        b.nonce,
//			PreviousHash: b.previousHash,
//			Transactions: b.transactions,
//			Timestamp:    b.timestamp,
//		},
//	)
//	return vo.NewHash(m).ValueToHex()
//}

// ProofOfWork ナンスを導き出す
func (b *Block) ProofOfWork(id uint32, previousHash string, transactions []string, timestamp int64) uint32 {

	miningDifficulty := uint32(3) // ハッシュ値が 000になるまで計算
	nonce := uint32(0)

	for !b.ValidProof(id, nonce, previousHash, transactions, timestamp, miningDifficulty) {
		nonce += 1
	}
	return nonce
}

// ValidProof  ナンスの計算
func (b *Block) ValidProof(id uint32, nonce uint32, previousHash string, transactions []string, timestamp int64, difficulty uint32) bool {
	zeros := strings.Repeat("0", int(difficulty))

	guessBlock := &Block{
		id:           id,
		nonce:        nonce,
		previousHash: vo.NewHexHashToByte32(previousHash).Value(),
		transactions: vo.NewTransactions(transactions).Value(),
		timestamp:    timestamp,
	}

	guessHashStr := b.GuessBlockHash(guessBlock)
	log.Println(guessHashStr)

	return guessHashStr[:difficulty] == zeros
}

// Hash ブロックのハッシュ値を求める
func (b *Block) Hash() string {

	m, _ := json.Marshal(
		struct {
			Id           uint32   `json:"id"`
			Nonce        uint32   `json:"nonce"`
			PreviousHash [32]byte `json:"previousHash"`
			Transactions string   `json:"transactions"`
			Timestamp    int64    `json:"timestamp"`
		}{
			Id:           b.id,
			Nonce:        b.nonce,
			PreviousHash: b.previousHash,
			Transactions: b.transactions,
			Timestamp:    b.timestamp,
		},
	)
	return vo.NewHash(m).ValueToHex()
}

// GuessBlockHash ナンスを計算する際のブロックのハッシュ値を求める
func (b *Block) GuessBlockHash(argB *Block) string {

	m, _ := json.Marshal(
		struct {
			Id           uint32   `json:"id"`
			Nonce        uint32   `json:"nonce"`
			PreviousHash [32]byte `json:"previousHash"`
			Transactions string   `json:"transactions"`
			Timestamp    int64    `json:"timestamp"`
		}{
			Id:           argB.id,
			Nonce:        argB.nonce,
			PreviousHash: argB.previousHash,
			Transactions: argB.transactions,
			Timestamp:    argB.timestamp,
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

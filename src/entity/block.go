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

	b.id = vo.NewID().Value()
	b.previousHash = vo.NewHexHashToByte32(previousHash).Value()
	b.transactions = vo.NewTransactions(transactions).Value()
	b.timestamp = vo.NewUnixTimeNow().Value()
	b.nonce = b.ProofOfWork() // この値は一番下に書かないといけない。他のBlockのstructの値を使ってProofOfWorkを行うため

	return b, nil
}

// ProofOfWork ナンスを導き出す
func (b *Block) ProofOfWork() uint32 {

	miningDifficulty := uint32(4) // ハッシュ値が 000になるまで計算
	b.nonce = uint32(0)

	for !b.ValidProof(miningDifficulty) {
		b.nonce += 1
	}
	return b.nonce
}

// ValidProof  ナンスの計算
func (b *Block) ValidProof(difficulty uint32) bool {
	zeros := strings.Repeat("0", int(difficulty))

	guessHashStr := b.Hash()
	log.Printf("%v : %v", b.nonce, guessHashStr)

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

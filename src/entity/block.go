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
	hash         [32]byte // hash値以外の Blockの値をhashした結果
}

func (b *Block) Id() uint32 {
	return b.id
}

func (b *Block) Nonce() uint32 {
	return b.nonce
}

func (b *Block) PreviousHash() string {

	return fmt.Sprintf("%x", b.previousHash)
}

func (b *Block) Hash() string {

	return fmt.Sprintf("%x", b.hash)
}

func (b *Block) Transactions() string {
	return b.transactions
}

func (b *Block) Timestamp() int64 {
	return b.timestamp
}

//===========================================================
//　プリミティブなBlockのEntityを作成
//===========================================================

func NewBlock(id, nonce uint32, hash, previousHash, transactions string, timestamp int64) (*Block, error) {

	return &Block{
		id:           vo.ID(id).Value(),
		previousHash: vo.NewHexHashToByte32(previousHash).Value(),
		transactions: transactions,
		timestamp:    timestamp,
		nonce:        nonce,
		hash:         vo.NewHexHashToByte32(hash).Value(),
	}, nil

}

//===========================================================
// GenWhenCreateBlock 新しいBlockを作成する時に使用
//===========================================================

func GenWhenCreateBlock(previousHash string, transactions []string) (*Block, error) {
	b := &Block{}

	b.previousHash = vo.NewHexHashToByte32(previousHash).Value()
	b.transactions = vo.NewTransactions(transactions).Value()
	b.timestamp = vo.NewUnixTimeNow().Value()
	b.nonce = b.ProofOfWork() // この値は一番下に書かないといけない。他のBlockのstructの値を使ってProofOfWorkを行うため

	return b, nil
}

// ProofOfWork ナンスを導き出す
func (b *Block) ProofOfWork() uint32 {

	miningDifficulty := uint32(3) // ハッシュ値が 000になるまで計算
	b.nonce = uint32(0)

	for !b.ValidProof(miningDifficulty) {
		b.nonce += 1
	}
	return b.nonce
}

// ValidProof  ナンスの計算
func (b *Block) ValidProof(difficulty uint32) bool {
	zeros := strings.Repeat("0", int(difficulty))

	guessHashStr := b.CalcHash()
	log.Printf("%v : %v", b.nonce, guessHashStr)

	return guessHashStr[:difficulty] == zeros
}

// CalcHash ブロックのハッシュ値を求める
func (b *Block) CalcHash() string {

	m, _ := json.Marshal(
		struct {
			Nonce        uint32   `json:"nonce"`
			PreviousHash [32]byte `json:"previousHash"`
			Transactions string   `json:"transactions"`
			Timestamp    int64    `json:"timestamp"`
		}{
			Nonce:        b.nonce,
			PreviousHash: b.previousHash,
			Transactions: b.transactions,
			Timestamp:    b.timestamp,
		},
	)

	return vo.NewHash(m).ValueToHex()
}

// IsCleanData 改竄されているか検証
func (b *Block) IsCleanData() bool {

	// DBに保存してあるハッシュ値とDBにあるBlockのデータをハッシュ値を計算して同じ値になるか検証
	if b.Hash() != b.CalcHash() {
		log.Println("改竄されている")
		return false
	}
	return true
}

package entity

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type BlockChain struct {
	transactionPool   []*Transaction
	chain             []*Block
	BlockChainAddress string
}

func NewBlockChain(blockChainAddress string) *BlockChain {
	b := &Block{}
	bc := &BlockChain{}
	bc.BlockChainAddress = blockChainAddress
	bc.CreateBlock(0, b.Hash())
	return bc
}

// CreateBlock ブロックを作成
func (bc *BlockChain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

// LastBlock 一個前のブロックを表示
func (bc *BlockChain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

// AddTransaction トランザクションの作成
func (bc *BlockChain) AddTransaction(sender string, recipient string, value float32,
	senderPublicKey *ecdsa.PublicKey, s *Signature) bool {
	t := NewTransaction(sender, recipient, value)

	if sender == "MiningSender" {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	}

	if bc.VerifyTransactionSignature(senderPublicKey, s, t) {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	} else {
		log.Println("Error:　トランザクションの検証でエラー")
	}
	return false
}

// VerifyTransactionSignature 送信されてきたトランザクションが正しいか検証
func (bc *BlockChain) VerifyTransactionSignature(senderPublicKey *ecdsa.PublicKey, s *Signature, t *Transaction) bool {
	m, _ := json.Marshal(t)
	h := sha256.Sum256(m)
	return ecdsa.Verify(senderPublicKey, h[:], s.R, s.S)
}

// CopyTransactionPool トランザクションプールのコピー
func (bc *BlockChain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions, NewTransaction(t.senderBlockChainAddress, t.recipientBlockChainAddress, t.value))
	}
	return transactions
}

// ValidProof ナンスの計算
func (bc *BlockChain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nonce, previousHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

// Mining マイニングの処理
func (bc *BlockChain) Mining() bool {
	bc.AddTransaction(MiningSender, bc.BlockChainAddress, MiningReward, nil, nil)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining, status=success")
	return true
}

// CalculateTotalAmount アドレスごとの所持金額
func (bc *BlockChain) CalculateTotalAmount(blockChainAddress string) float32 {
	var totalAmount float32 = 0.0
	for _, b := range bc.chain {
		for _, t := range b.transactions {
			value := t.value
			if blockChainAddress == t.recipientBlockChainAddress {
				totalAmount += value
			}

			if blockChainAddress == t.senderBlockChainAddress {
				totalAmount -= value
			}
		}
	}

	return totalAmount
}

// ProofOfWork ナンスを導き出す
func (bc *BlockChain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := -0
	for !bc.ValidProof(nonce, previousHash, transactions, MiningDifficulty) {
		nonce += 1
	}
	return nonce
}

// Print ブロックチェーンを見やすいように表示
func (bc *BlockChain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s \n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s \n", strings.Repeat("*", 25))
}

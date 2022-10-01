package output

type BlockList struct {
	BlockChain []*Block `json:"blockChain"`
}

type Block struct {
	ID           uint32 `json:"id"`
	Nonce        uint32 `json:"nonce"`
	Transactions string `json:"transactions"`
	PreviousHash string `json:"previousHash"`
	Timestamp    int64  `json:"timestamp"`
	Hash         string `json:"hash"`
}

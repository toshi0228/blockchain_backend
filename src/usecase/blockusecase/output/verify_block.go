package output

type VerifyBlockList struct {
	VerifyBlockChain []*VerifyBlock `json:"verifyBlockChain"`
}

type VerifyBlock struct {
	ID           uint32 `json:"id"`
	Nonce        uint32 `json:"nonce"`
	Transactions string `json:"transactions"`
	PreviousHash string `json:"previousHash"`
	Timestamp    int64  `json:"timestamp"`
	Hash         string `json:"hash"`
	IsCleanData  bool   `json:"isCleanData"`
}

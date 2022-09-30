package datamodel

type BlocK struct {
	Id           uint32 `db:"id"`
	Nonce        uint32 `db:"nonce"`
	Transactions string `db:"transactions"`
	PreviousHash string `db:"previous_hash"`
	Timestamp    int64  `db:"timestamp"`
	Hash         string `db:"hash"`
}

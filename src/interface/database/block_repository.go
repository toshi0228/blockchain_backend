package database

import (
	_ "embed"
	"fmt"
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/infra/db"
	"github.com/toshi0228/blockchain/src/usecase/blockusecase/input"
	"log"
)

type BlockRepositoryImpl struct{}

func NewBlockRepositoryImpl() *BlockRepositoryImpl {
	return &BlockRepositoryImpl{}
}

//===========================================================
// Block Create
//===========================================================

//go:embed block_repository_create.sql
var createBlockSql string

//go:embed block_repository_find_prev_block.sql
var findPrevBlockSql string

func (repo *BlockRepositoryImpl) Create(in *input.CreateBlockInput) ([]*entity.Block, error) {

	// TODO ブロックがなければ作成
	// TODO トランザクションプールのデータを取得

	//一個前のブロックのハッシュを取得
	row := db.Conn().QueryRow(findPrevBlockSql)
	var prevHash string
	err := row.Scan(&prevHash)
	if err != nil {
		return nil, err
	}

	b, err := entity.NewBlock(in.Nonce, prevHash, in.TransactionsHash)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	_, err = db.Conn().Exec(
		createBlockSql,
		b.Id(),
		b.Nonce(),
		b.PreviousHashToHex(),
		b.TransactionsHashToHex(),
		b.Timestamp(),
		b.Hash(),
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return nil, nil
}

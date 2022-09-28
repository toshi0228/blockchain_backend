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

func (repo *BlockRepositoryImpl) Create(in *input.CreateBlockInput) ([]*entity.Block, error) {

	b, err := entity.NewBlock(in.Nonce, in.PreviousHash, in.TransactionsHash)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	//cmd := fmt.Sprintf(createBlockSql, ***)

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

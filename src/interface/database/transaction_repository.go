package database

import (
	_ "embed"
	"fmt"
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/infra/db"
	"github.com/toshi0228/blockchain/src/usecase/transactionusecase/input"
	"log"
)

type TransactionRepositoryImpl struct{}

func NewTransactionRepositoryImpl() *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{}
}

//===========================================================
// Transaction Create
//===========================================================

//go:embed transaction_repository_create.sql
var createTransactionSql string

//go:embed transaction_repository_create_pool.sql
var createTransactionPooLSql string

func (repo *TransactionRepositoryImpl) Create(in *input.CreateTransactionInput) (*entity.Transactions, error) {

	tx, err := entity.GenWhenCreateTransactions(in.SenderAddress, in.RecipientAddress, in.PrivateKey, in.PublicKey, in.Signature, in.Amount)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf(err.Error())
	}

	//トランザクションの作成
	_, err = db.Conn().Exec(
		createTransactionSql,
		tx.Id().Value(),
		tx.SenderAddress(),
		tx.RecipientAddress(),
		tx.Value(),
		tx.CreatedAt().Value(),
		tx.UpdatedAt().Value(),
	)
	if err != nil {
		return nil, err
	}

	//トランザクションプールの作成
	_, err = db.Conn().Exec(
		createTransactionPooLSql,
		tx.Id().Value(),
		tx.SenderAddress(),
		tx.RecipientAddress(),
		tx.Value(),
		tx.CreatedAt().Value(),
		tx.UpdatedAt().Value(),
	)

	return tx, nil
}

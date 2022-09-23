package database

import (
	_ "embed"
	"fmt"
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/infra/db"
	"github.com/toshi0228/blockchain/src/usecase/transactionusecase/input"
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

func (repo *TransactionRepositoryImpl) Create(in *input.CreateTransactionInput) (*entity.Transaction, error) {

	tx, err := entity.GenWhenCreateTransaction(in.SenderAddress, in.RecipientAddress, in.Amount)

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

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

	return tx, nil
}

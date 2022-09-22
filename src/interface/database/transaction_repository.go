package database

import (
	_ "embed"
	"github.com/toshi0228/blockchain/src/entity"
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

func (repo *TransactionRepositoryImpl) Create(in *input.CreateTransactionInput) ([]*entity.Transaction, error) {

	//_, err := entity.NewTransaction(in.Name)
	//if err != nil {
	//	return nil, fmt.Errorf(err.Error())
	//}

	//cmd := fmt.Sprintf(createTransactionSql, ***)

	//_, err = db.Conn().Exec(createTransactionSql, u.Id())

	return nil, nil
}

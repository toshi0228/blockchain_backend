package database

import (
	_ "embed"
	"fmt"
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/infra/datamodel"
	"github.com/toshi0228/blockchain/src/infra/db"
	"github.com/toshi0228/blockchain/src/usecase/transactionpoolusecase/input"
)

type TransactionPoolRepositoryImpl struct{}

func NewTransactionPoolRepositoryImpl() *TransactionPoolRepositoryImpl {
	return &TransactionPoolRepositoryImpl{}
}

//===========================================================
// 　トランザクションプールのデータを取得する
//===========================================================

//go:embed transaction_pool_repository_find_all.sql
var findAllTxPoolSql string

func (repo *TransactionPoolRepositoryImpl) FindAll(in *input.GetTransactionPoolInput) ([]*entity.TransactionPool, error) {

	//DBからトランザクションの取得
	var txPool []*datamodel.TransactionPool
	err := db.Conn().Select(&txPool, findAllTxPoolSql)
	if err != nil {
		return nil, err
	}

	// DBから取得したトランザクションをentityに代入
	var entityTxPool []*entity.TransactionPool
	for i, tx := range txPool {
		fmt.Println(i, tx)

		tx, err := entity.NewTransactionPool(tx.Id, tx.SenderAddress, tx.RecipientAddress, tx.Amount, tx.CreatedAt, tx.UpdatedAt)
		if err != nil {
			return nil, err
		}

		entityTxPool = append(entityTxPool, tx)
	}

	return entityTxPool, nil

}

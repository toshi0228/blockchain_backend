package dbtable

import (
	_ "embed"
	"fmt"
)

const TableNameTransactionPool = "transactionPool"

//go:embed transactionPool.sql
var transactionPoolTableSql string

func TransactionPools() string {

	cmd := fmt.Sprintf(transactionPoolTableSql, TableNameTransactionPool)
	return cmd
}

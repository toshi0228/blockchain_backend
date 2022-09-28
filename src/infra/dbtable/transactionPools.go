package dbtable

import (
	_ "embed"
	"fmt"
)

const TableNameTransactionPool = "transactionPools"

//go:embed transactionPools.sql
var transactionPoolsTableSql string

func TransactionPools() string {

	cmd := fmt.Sprintf(transactionPoolsTableSql, TableNameTransactionPool)
	return cmd
}

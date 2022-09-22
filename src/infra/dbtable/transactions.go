package dbtable

import (
	_ "embed"
	"fmt"
)

const TableNameTransaction = "transactions"

//go:embed transactions.sql
var transactionsTableSql string

func Transactions() string {

	cmd := fmt.Sprintf(transactionsTableSql, TableNameTransaction)
	return cmd
}

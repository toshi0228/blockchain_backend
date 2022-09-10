package dbtable

import (
	_ "embed"
	"fmt"
)

const TableNameWallet = "wallets"

//go:embed wallets.sql
var walletsTableSql string

func Wallets() string {

	cmd := fmt.Sprintf(walletsTableSql, TableNameWallet)
	return cmd
}

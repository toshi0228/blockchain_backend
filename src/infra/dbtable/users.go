package dbtable

import (
	_ "embed"
	"fmt"
)

const TableNameUser = "users"

//go:embed users.sql
var usersTableSql string

func Users() string {
	cmd := fmt.Sprintf(usersTableSql, TableNameUser)
	return cmd
}

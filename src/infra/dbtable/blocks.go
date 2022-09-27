package dbtable

import (
	_ "embed"
	"fmt"
)

const TableNameBlock = "blocks"

//go:embed blocks.sql
var blocksTableSql string

func Blocks() string {

	cmd := fmt.Sprintf(blocksTableSql, TableNameBlock)
	return cmd
}

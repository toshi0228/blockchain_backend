---
to: src/infra/dbtable/<%= h.changeCase.camel(table) %>s.go
unless_exists: true
---

package dbtable

import (
	_ "embed"
	"fmt"
)

const TableName<%= h.changeCase.pascal(table) %> = "<%= h.changeCase.camel(table) %>s"

//go:embed <%= h.changeCase.camel(table) %>s.sql
var <%= h.changeCase.camel(table) %>sTableSql string

func <%= h.changeCase.pascal(table) %>s() string {

	cmd := fmt.Sprintf(<%= h.changeCase.camel(table) %>sTableSql, TableName<%= h.changeCase.pascal(table) %>)
	return cmd
}

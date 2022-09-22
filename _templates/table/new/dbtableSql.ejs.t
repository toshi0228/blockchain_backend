---
to: src/infra/dbtable/<%= h.changeCase.camel(entity) %>s.go
unless_exists: true
---

package dbtable

import (
	_ "embed"
	"fmt"
)

const TableName<%= h.changeCase.pascal(entity) %> = "<%= h.changeCase.camel(entity) %>s"

//go:embed <%= h.changeCase.camel(entity) %>s.sql
var <%= h.changeCase.camel(entity) %>sTableSql string

func <%= h.changeCase.pascal(entity) %>s() string {

	cmd := fmt.Sprintf(<%= h.changeCase.camel(entity) %>sTableSql, TableName<%= h.changeCase.pascal(entity) %>)
	return cmd
}

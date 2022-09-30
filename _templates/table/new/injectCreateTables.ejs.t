---
to: src/infra/db/createTables.go
unless_exists: true
inject: true
after: createTableCMD
---
	"<%= h.changeCase.camel(table) %>s":    dbtable.<%= h.changeCase.pascal(table) %>s(),

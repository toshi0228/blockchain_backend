---
to: src/infra/db/createTables.go
unless_exists: true
inject: true
after: createTableCMD
---
	"<%= h.changeCase.camel(entity) %>s":    dbtable.<%= h.changeCase.camel(entity) %>s(),

---
to: src/infra/db/reset.go
unless_exists: true
inject: true
after: tableList
---
	"<%= h.changeCase.camel(table) %>s":    dbtable.TableName<%= h.changeCase.pascal(table) %>,

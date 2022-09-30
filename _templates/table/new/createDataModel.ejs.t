---
to: src/infra/datamodel/<%= h.changeCase.snake(table) %>_datamodel.go
unless_exists: true
---
package datamodel

import (
	"time"
)

//===========================================================
//　DBにある<%= h.changeCase.camel(table) %>のデータを格納するStruct
//===========================================================

type <%= h.changeCase.camel(table) %> struct {
	Id               uint32    `db:"id"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}

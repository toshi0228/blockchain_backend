---
to: src/entity/<%= h.changeCase.camel(table) %>.go
unless_exists: true
---
package entity

import (
	"github.com/toshi0228/blockchain/src/entity/vo"
)

type <%= h.changeCase.camel(table) %> struct {
	id           uint32
	createdAt vo.CreatedAt `json:"createdAt"`
	updatedAt vo.UpdatedAt `json:"updatedAt"`
}

//===========================================================
//　プリミティブな<%= h.changeCase.camel(table) %>のEntityを作成
//===========================================================

func New<%= h.changeCase.camel(table) %>(
	id uint32,
	createdAt time.Time,
	updatedAt time.Time,
) (*<%= h.changeCase.camel(table) %>, error) {

	return &Transaction{
		id:               vo.ID(id),
		createdAt:        vo.CreatedAt(createdAt),
		updatedAt:        vo.UpdatedAt(updatedAt),
	}, nil
}

//===========================================================
// GenWhenCreate<%= h.changeCase.camel(table) %> 新しい<%= h.changeCase.camel(table) %>を作成する時に使用
//===========================================================

func GenWhenCreate<%= h.changeCase.camel(table) %>() (*<%= h.changeCase.camel(table) %>, error) {

	return &<%= h.changeCase.camel(table) %>{
		id:               vo.ID(id),
		createdAt:        vo.CreatedAt(createdAt),
		updatedAt:        vo.UpdatedAt(updatedAt),
	}, nil
}

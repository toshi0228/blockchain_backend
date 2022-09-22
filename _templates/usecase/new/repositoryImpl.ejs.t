---
to: src/interface/database/<%= h.changeCase.snake(entity) %>_repository.go
unless_exists: true
---
package database

import (
	_ "embed"
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/usecase/<%=　h.changeCase.lower(entity) %>usecase/input"
)

type <%= entity %>RepositoryImpl struct{}


func New<%= entity %>RepositoryImpl() *<%= entity %>RepositoryImpl {
	return &<%= entity %>RepositoryImpl{}
}

//===========================================================
// <%= h.changeCase.pascal(entity) %> <%=　h.changeCase.pascal(method) %>
//===========================================================

//go:embed <%= h.changeCase.snake(entity) %>_repository_<%= h.changeCase.snake(method) %>.sql
var <%= h.changeCase.camel(method) %><%= h.changeCase.pascal(entity) %>Sql string


func (repo *<%= entity %>RepositoryImpl) <%= h.changeCase.pascal(method) %>(in *input.<%= h.changeCase.pascal(useCaseName) %>Input) ([]*entity.<%= h.changeCase.pascal(entity) %>, error) {

	//_, err := entity.New<%= h.changeCase.pascal(entity) %>(in.Name)
	//if err != nil {
	//	return nil, fmt.Errorf(err.Error())
	//}

	//cmd := fmt.Sprintf(<%= h.changeCase.camel(method) %><%= h.changeCase.pascal(entity) %>Sql, ***)

	//_, err = db.Conn().Exec(<%= h.changeCase.camel(method) %><%= h.changeCase.pascal(entity) %>Sql, u.Id())

	return nil, nil
}


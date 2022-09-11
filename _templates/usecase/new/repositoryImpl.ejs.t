---
to: src/interface/database/<%= h.changeCase.snake(entity) %>_repository.go
unless_exists: true
---
package database

import (
	_ "embed"
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/usecase/<%= entity %>usecase/input"
)

type <%= entity %>RepositoryImpl struct{}

//go:embed <%= h.changeCase.snake(entity) %>_repository_<%= h.changeCase.snake(method) %>.sql
var <%= h.changeCase.camel(method) %><%= h.changeCase.pascal(entity) %>Sql string

func New<%= entity %>RepositoryImpl() *<%= entity %>RepositoryImpl {
	return &<%= entity %>RepositoryImpl{}
}


func (repo *<%= entity %>RepositoryImpl) <%= h.changeCase.pascal(method) %>(in *input.<%= h.changeCase.pascal(useCaseName) %>Input) ([]*entity.<%= entity %>, error) {

	//_, err := entity.New<%= entity %>(in.Name)
	//if err != nil {
	//	return nil, fmt.Errorf(err.Error())
	//}

	//cmd := fmt.Sprintf(<%= h.changeCase.camel(method) %><%= h.changeCase.camel(entity) %>Sql, dbtable.TableNameUser)
	//_, err = db.Conn().Exec(cmd, u.Id())

	return nil, nil
}


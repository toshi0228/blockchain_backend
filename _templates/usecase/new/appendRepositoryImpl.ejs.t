---
to: src/interface/database/<%=h.changeCase.snake(entity) %>_repository.go
unless_exists: true
inject: true
append: dependencies
---

//===========================================================
// <%= h.changeCase.pascal(entity) %> <%=　h.changeCase.pascal(method) %>
//===========================================================

//go:embed <%= h.changeCase.snake(entity) %>_repository_<%= h.changeCase.snake(method) %>.sql
var <%= h.changeCase.camel(method) %><%= h.changeCase.pascal(entity) %>Sql string

func (repo *<%= entity %>RepositoryImpl) <%=h.changeCase.pascal(method) %>(in *input.<%= h.changeCase.pascal(useCaseName) %>Input) ([]*entity.<%= h.changeCase.pascal(entity) %>, error) {

	//_, err := entity.New<%= entity %>(in.Name)
	//if err != nil {
	//	return nil, fmt.Errorf(err.Error())
	//}

	//cmd := fmt.Sprintf(<%= h.changeCase.camel(method) %><%= h.changeCase.pascal(entity) %>Sql, ***)

	//_, err = db.Conn().Exec(<%= h.changeCase.camel(method) %><%= h.changeCase.pascal(entity) %>Sql, u.Id())

	return nil, nil
}

---
to: src/interface/database/<%=h.changeCase.snake(entity) %>_repository.go
unless_exists: true
inject: true
append: dependencies
---

func (repo *<%= entity %>RepositoryImpl) <%=h.changeCase.pascal(method) %>(in *input.<%= h.changeCase.pascal(useCaseName) %>Input) ([]*entity.<%= entity %>, error) {

	//_, err := entity.New<%= entity %>(in.Name)
	//if err != nil {
	//	return nil, fmt.Errorf(err.Error())
	//}

	//cmd := fmt.Sprintf(createUserSql, dbtable.TableNameUser)
	//_, err = db.Conn().Exec(cmd, u.Id())

	return nil, nil
}


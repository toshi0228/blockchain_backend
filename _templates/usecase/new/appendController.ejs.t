---
to: src/interface/controller/<%= h.changeCase.snake(entity) %>_controller.go
unless_exists: true
inject: true
append: dependencies
---


//===========================================================
//　<%= desc %>
//===========================================================

func (c *<%= h.changeCase.camel(entity) %>Controller) <%=　h.changeCase.pascal(useCaseName) %>(in *input.<%= h.changeCase.pascal(useCaseName) %>Input) error {
	usecase := <%= h.changeCase.lower(entity) %>usecase.New<%=　h.changeCase.pascal(useCaseName) %>(c.<%= h.changeCase.camel(entity) %>Repo)
	out, err := usecase.Exec(in)
	if err != nil {
		return err
	}

	return c.delivery.<%= h.changeCase.pascal(output) %>(out)
}

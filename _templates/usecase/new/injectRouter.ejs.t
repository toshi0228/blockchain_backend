---
to: src/infra/router/<%=　h.changeCase.snake(entity) %>_router.go
unless_exists: true
inject: true
after: func
---
    //===========================================================
    //　<%= desc %>
    //===========================================================
	//e.POST("/<%=　h.changeCase.lower(entity) %>/<%=　h.changeCase.camel(useCaseName) %>", func(c echo.Context) error {
	//
	//	in := &input.<%= h.changeCase.pascal(useCaseName) %>Input{}
	//	err := c.Bind(in)
	//	if err != nil {
	//		return err
	//	}
	//
	//	<%= h.changeCase.camel(entity) %>RepoImpl := database.New<%= h.changeCase.pascal(entity) %>RepositoryImpl()
	//	err = controller.New<%= h.changeCase.pascal(entity) %>Controller(presenter.New<%= h.changeCase.pascal(entity) %>Presenter(c), <%= h.changeCase.camel(entity) %>RepoImpl).<%=　h.changeCase.pascal(useCaseName) %>(in)
	//	if err != nil {
	//		return c.JSON(http.StatusInternalServerError, err.Error())
	//	}
	//
	//	return nil
	//})

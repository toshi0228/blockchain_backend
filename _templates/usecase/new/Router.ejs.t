---
to: src/infra/router/<%=　h.changeCase.snake(entity) %>_router.go
unless_exists: true
---
package router

import (
	"github.com/labstack/echo/v4"
	"github.com/toshi0228/blockchain/src/interface/database"
	"github.com/toshi0228/blockchain/src/usecase/<%=　h.changeCase.lower(entity) %>usecase"
	"net/http"
)

func New<%= h.changeCase.pascal(entity) %>Router(e *echo.Echo) {

	//e.POST("/<%=　h.changeCase.lower(entity) %>", func(c echo.Context) error {
	//
	//	in := &input.<%= h.changeCase.pascal(useCaseName) %>Input{}
	//	err := c.Bind(in)
	//	if err != nil {
	//		return fmt.Errorf("エラー")
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

}


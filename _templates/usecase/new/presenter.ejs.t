---
to: src/interface/presenter/<%= h.changeCase.snake(entity) %>_presenter.go
unless_exists: true
---
package presenter

import (
	"github.com/labstack/echo/v4"
	"github.com/toshi0228/blockchain/src/usecase/<%= h.changeCase.lower(entity) %>usecase/output"
	"net/http"
)

type <%= h.changeCase.camel(entity) %>Present struct {
	ctx echo.Context
}

func New<%= h.changeCase.pascal(entity) %>Presenter(p echo.Context) I<%= h.changeCase.pascal(entity) %>Presenter {
	return &<%= h.changeCase.camel(entity) %>Present{ctx: p}
}

type I<%= h.changeCase.pascal(entity) %>Presenter interface {
	<%= h.changeCase.pascal(output) %>(out *output.<%= h.changeCase.pascal(output) %>) error
}

func (p *<%= h.changeCase.camel(entity) %>Present) <%= h.changeCase.pascal(output) %>(out *output.<%= h.changeCase.pascal(output) %>) error {
	return p.ctx.JSON(http.StatusOK, out)
}

---
to: src/interface/presenter/<%= h.changeCase.snake(entity) %>_presenter.go
unless_exists: true
inject: true
append: dependencies
---


//===========================================================
//　<%= desc %>
//===========================================================

func (p *<%= h.changeCase.camel(entity) %>Present) <%= h.changeCase.pascal(output) %>(out *output.<%= h.changeCase.pascal(output) %>) error {
	return p.ctx.JSON(http.StatusOK, out)
}

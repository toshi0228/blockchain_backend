---
to: src/interface/controller/<%= h.changeCase.snake(entity) %>_controller.go
unless_exists: true
---
package controller

import (
	"github.com/toshi0228/blockchain/src/interface/presenter"
	"github.com/toshi0228/blockchain/src/usecase/<%= h.changeCase.lower(entity) %>usecase"
	"github.com/toshi0228/blockchain/src/usecase/<%= h.changeCase.lower(entity) %>usecase/input"
)

type <%= h.changeCase.camel(entity) %>Controller struct {
	<%= h.changeCase.camel(entity) %>Repo <%= h.changeCase.lower(entity) %>usecase.I<%= h.changeCase.pascal(entity) %>Repository
}

func New<%= h.changeCase.pascal(entity) %>Controller(<%= h.changeCase.camel(entity) %>Repo <%= h.changeCase.lower(entity) %>usecase.I<%= h.changeCase.pascal(entity) %>Repository) *<%= h.changeCase.camel(entity) %>Controller {
	return &<%= h.changeCase.camel(entity) %>Controller{
		<%= h.changeCase.camel(entity) %>Repo: <%= h.changeCase.camel(entity) %>Repo,
	}
}

func (c *<%= h.changeCase.camel(entity) %>Controller) <%=　h.changeCase.pascal(useCaseName) %>(in *input.<%= h.changeCase.pascal(useCaseName) %>Input) error {
	usecase := <%= h.changeCase.lower(entity) %>usecase.New<%=　h.changeCase.pascal(useCaseName) %>(c.<%= h.changeCase.camel(entity) %>Repo)
	out, err := usecase.Exec(in)
	if err != nil {
		return err
	}
	return nil
	// return c.delivery.UserList(out)

}

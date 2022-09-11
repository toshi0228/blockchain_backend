---
to: src/usecase/<%=　h.changeCase.lower(entity) %>usecase/<%=　h.changeCase.snake(useCaseName) %>.go
unless_exists: true
---
package <%=　h.changeCase.lower(entity) %>usecase

import (
	"github.com/toshi0228/blockchain/src/usecase/<%=　h.changeCase.lower(entity) %>usecase/input"
	"github.com/toshi0228/blockchain/src/usecase/<%=　h.changeCase.lower(entity) %>usecase/output"
)

type <%=　h.changeCase.pascal(useCaseName) %>usecase struct {
	<%= entity %>Repository I<%= h.changeCase.pascal(entity) %>Repository
}

func New<%=　h.changeCase.pascal(useCaseName) %>(<%= entity %>Repo I<%= h.changeCase.pascal(entity) %>Repository) *<%=　h.changeCase.pascal(useCaseName) %>usecase {
	return &<%=　h.changeCase.pascal(useCaseName) %>usecase{
		<%= entity %>Repository: <%= entity %>Repo,
	}
}

func (use *<%=　h.changeCase.pascal(useCaseName) %>usecase) Exec(in *input.<%= h.changeCase.pascal(useCaseName) %>Input) (*output.<%= h.changeCase.pascal(output) %>, error) {
	_, err := use.<%= entity %>Repository.<%=　h.changeCase.pascal(method) %>(in)
	if err != nil {
		return nil, err
	}

	return &output.<%= h.changeCase.pascal(output) %>{}, nil
}


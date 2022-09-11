---
to: src/usecase/<%= entity %>usecase/I_<%= h.changeCase.snake(entity) %>_repository.go
unless_exists: true
---
package <%=　h.changeCase.lower(entity) %>usecase

import "github.com/toshi0228/blockchain/src/entity"

type I<%= h.changeCase.pascal(entity) %>Repository interface {
	<%= h.changeCase.pascal(method) %>(in *input.<%= h.changeCase.pascal(useCaseName) %>Input) ([]*entity.<%= entity %>, error)
}


---
to: src/usecase/<%=　h.changeCase.lower(entity) %>usecase/I_<%=　h.changeCase.snake(entity) %>_repository.go
unless_exists: true
inject: true
after: interface {
---
    <%= h.changeCase.pascal(method) %>(in *input.<%= h.changeCase.pascal(useCaseName) %>Input) ([]*entity.<%= h.changeCase.pascal(entity) %>, error)

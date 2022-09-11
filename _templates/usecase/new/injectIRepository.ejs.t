---
to: src/usecase/<%= entity %>usecase/I_<%=　h.changeCase.snake(entity) %>_repository.go
unless_exists: true
inject: true
after: interface {
---
    <%= h.changeCase.pascal(method) %>(in *input.<%= h.changeCase.pascal(useCaseName) %>Input) ([]*entity.<%= entity %>, error)


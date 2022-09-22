---
to: src/interface/presenter/<%= h.changeCase.snake(entity) %>_presenter.go
unless_exists: true
inject: true
after: interface {
---




	<%= h.changeCase.pascal(output) %>(out *output.<%= h.changeCase.pascal(output) %>) error



---
to: src/usecase/<%=　h.changeCase.lower(entity) %>usecase/input/<%=　h.changeCase.snake(useCaseName) %>_input.go

unless_exists: true
---
package input

type <%= h.changeCase.pascal(useCaseName) %>Input struct {
	ID uint32
}


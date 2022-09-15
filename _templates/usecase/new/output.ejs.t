---
to: src/usecase/<%=　h.changeCase.lower(entity) %>usecase/output/<%=　h.changeCase.snake(useCaseName) %>.go

unless_exists: true
---
package output

import "time"

type <%= h.changeCase.pascal(output) %> struct {
	ID uint32 `json:"id"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

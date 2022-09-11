---
to: src/interface/database/<%= h.changeCase.snake(entity) %>_repository_<%= h.changeCase.snake(method) %>.sql
unless_exists: true
---
INSERT INTO %s
(
    id,
    created_at,
    updated_at
)
VALUES (?,?,?)


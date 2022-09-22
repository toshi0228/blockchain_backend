package output

import "time"

type CreateTransaction struct {
	ID uint32 `json:"id"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

package output

import "time"

type BlockList struct {
	ID uint32 `json:"id"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

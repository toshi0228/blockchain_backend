package output

import "time"

type UserList struct {
	Users []*UserItem `json:"users"`
}

type UserItem struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

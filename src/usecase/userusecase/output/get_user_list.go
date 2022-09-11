package output

import "time"

type UserList struct {
	Users []*User `json:"users"`
}

type User struct {
	ID        uint32 `json:"id"`
	Name      string
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

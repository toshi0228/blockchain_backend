package output

import (
	"time"
)

type UserList struct {
	Users []*User `json:"users"`
}

type User struct {
	UserItem   *UserItem   `json:"user"`
	WalletItem *WalletItem `json:"wallet"`
}

type UserItem struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type WalletItem struct {
	ID        uint32    `json:"iD"`
	UserID    uint32    `json:"userID"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

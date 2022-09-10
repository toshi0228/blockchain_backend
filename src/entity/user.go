package entity

import "github.com/toshi0228/blockchain/src/entity/vo"

type User struct {
	id        vo.ID        `json:"id"`
	name      string       `json:"name"`
	password  string       `json:"password"`
	createdAt vo.CreatedAt `json:"createdAt"`
	updatedAt vo.UpdatedAt `json:"updatedAt"`
}

func (u User) Id() vo.ID {
	return u.id
}

func (u User) Name() string {
	return u.name
}

func (u User) Password() string {
	return u.password
}

func (u User) CreatedAt() vo.CreatedAt {
	return u.createdAt
}

func (u User) UpdatedAt() vo.UpdatedAt {
	return u.updatedAt
}

func NewUser(name string, password string) (*User, error) {
	return &User{
		id:        vo.NewID(),
		name:      name,
		password:  password,
		createdAt: vo.NewCreatedAt(),
		updatedAt: vo.NewUpdatedAt(),
	}, nil
}

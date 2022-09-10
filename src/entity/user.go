package entity

type User struct {
	name   string  `json:"name"`
	wallet *Wallet `json:"wallet"`
}

func NewUser(name string) *User {
	return &User{
		name:   name,
		wallet: &Wallet{},
	}
}

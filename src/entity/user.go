package entity

import (
	"github.com/toshi0228/blockchain/src/entity/vo"
	"time"
)

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

func NewUser(id uint32, name string, password string, createdAt time.Time, updatedAt time.Time) (*User, error) {
	return &User{
		id:        vo.ID(id),
		name:      name,
		password:  password,
		createdAt: vo.CreatedAt(createdAt),
		updatedAt: vo.UpdatedAt(updatedAt),
	}, nil
}

// GenWhenCreateUser ユーザー作成の際のメソッド
func GenWhenCreateUser(name string, password string) (*User, error) {
	return &User{
		id:        vo.NewID(),
		name:      name,
		password:  password,
		createdAt: vo.NewCreatedAt(),
		updatedAt: vo.NewUpdatedAt(),
	}, nil
}

//==================================================
//ユーザー一覧ページのユーザー
//==================================================

type UserPager struct {
	id        vo.ID        `json:"id"`
	name      string       `json:"name"`
	password  string       `json:"password"`
	createdAt vo.CreatedAt `json:"createdAt"`
	updatedAt vo.UpdatedAt `json:"updatedAt"`
	wallet    *GetWallet   `json:"wallet"`
}

func (u UserPager) Id() vo.ID {
	return u.id
}

func (u UserPager) Name() string {
	return u.name
}

func (u UserPager) Password() string {
	return u.password
}

func (u UserPager) CreatedAt() vo.CreatedAt {
	return u.createdAt
}

func (u UserPager) UpdatedAt() vo.UpdatedAt {
	return u.updatedAt
}

func (u UserPager) Wallet() *GetWallet {
	return u.wallet
}

func NewUserPager(
	id uint32,
	name string,
	password string,
	createdAt time.Time,
	updatedAt time.Time,

	// relationされたwallet情報
	walletId uint32,
	userID uint32,
	blockchainAddress string,
	walletCreatedAt time.Time,
	walletUpdatedAt time.Time,
) (*UserPager, error) {

	getWallet, err := NewGetWallet(walletId, userID, blockchainAddress, walletCreatedAt, walletUpdatedAt)
	if err != nil {
		return nil, err
	}

	return &UserPager{
		id:        vo.ID(id),
		name:      name,
		password:  password,
		createdAt: vo.CreatedAt(createdAt),
		updatedAt: vo.UpdatedAt(updatedAt),
		wallet:    getWallet,
	}, nil
}

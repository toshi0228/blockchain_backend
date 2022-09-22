package database

import (
	_ "embed"
	"fmt"
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/infra/datamodel"
	"github.com/toshi0228/blockchain/src/infra/db"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

//===========================================================
// UserPager FindAll
//===========================================================

//go:embed user_repository_find_all.sql
var findAllUserSql string

// FindAll ユーザーの一覧取得
func (repo *UserRepositoryImpl) FindAll() ([]*entity.UserPager, error) {

	rows, err := db.Conn().Queryx(findAllUserSql)
	if err != nil {
		return nil, err
	}

	if rows == nil {
		return nil, nil
	}

	var users []*entity.UserPager

	for rows.Next() {

		u := &datamodel.UserPager{}
		err := rows.StructScan(u)

		if err != nil {
			fmt.Println("エラー文")
			fmt.Println(err)
		}

		userEntity, err := entity.NewUserPager(
			u.Id,
			u.Name,
			u.Password,
			u.CreatedAt,
			u.UpdatedAt,
			u.WalletId,
			u.UserID,
			u.BlockchainAddress,
			u.WalletCreatedAt,
			u.WalletUpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}

		users = append(users, userEntity)
	}

	return users, nil
}

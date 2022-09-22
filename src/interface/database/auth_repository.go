package database

import (
	_ "embed"
	"fmt"
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/infra/datamodel"
	"github.com/toshi0228/blockchain/src/infra/db"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
	"log"
)

type AuthRepositoryImpl struct{}

func NewAuthRepositoryImpl() *AuthRepositoryImpl {
	return &AuthRepositoryImpl{}
}

//===========================================================
// Auth Login
//===========================================================

//go:embed auth_repository_login.sql
var loginAuthSql string

func (repo *AuthRepositoryImpl) Login(in *input.LoginInput) (*entity.UserPager, error) {

	rows, err := db.Conn().Queryx(loginAuthSql, in.Name, in.Password)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer rows.Close()

	var users []*entity.UserPager

	for rows.Next() {
		u := &datamodel.UserPager{}
		err := rows.StructScan(u)

		if err != nil {
			fmt.Println("エラー文")
			fmt.Println(err)
			return nil, err
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

	if len(users) == 0 {
		return nil, fmt.Errorf("ログインに失敗しました。")
	}

	return users[0], nil
}

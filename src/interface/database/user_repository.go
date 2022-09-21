package database

import (
	_ "embed"
	"fmt"
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/infra/datamodel"
	"github.com/toshi0228/blockchain/src/infra/db"
	"github.com/toshi0228/blockchain/src/infra/dbtable"
	"github.com/toshi0228/blockchain/src/usecase/userusecase/input"
	"log"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

//go:embed user_repository_create.sql
var createUserSql string

func (repo UserRepositoryImpl) CreateUser(in *input.CreateUserInput) (*entity.User, error) {

	u, err := entity.GenWhenCreateUser(in.Name, in.Password)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	cmd := fmt.Sprintf(createUserSql, dbtable.TableNameUser)

	_, err = db.Conn().Exec(
		cmd,
		u.Id(),
		u.Name(),
		u.Password(),
		u.CreatedAt().Value(),
		u.UpdatedAt().Value(),
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return u, nil
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

//===========================================================
// User Login
//===========================================================

//go:embed user_repository_login.sql
var loginUserSql string

func (repo *UserRepositoryImpl) Login(in *input.LoginUserInput) (*entity.UserPager, error) {

	//cmd := fmt.Sprintf(loginUserSql, in.Name, in.Password)

	rows, err := db.Conn().Queryx(loginUserSql, in.Name, in.Password)
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

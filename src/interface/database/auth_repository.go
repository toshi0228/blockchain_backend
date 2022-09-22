package database

import (
	_ "embed"
	"fmt"
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/infra/datamodel"
	"github.com/toshi0228/blockchain/src/infra/db"
	"github.com/toshi0228/blockchain/src/infra/dbtable"
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

//===========================================================
// Auth SignUp
//===========================================================

//go:embed auth_repository_sign_up.sql
var signUpAuthSql string

func (repo *AuthRepositoryImpl) SignUp(in *input.SignUpInput) (*entity.User, error) {

	//_, err := entity.NewAuth(in.Name)
	//if err != nil {
	//	return nil, fmt.Errorf(err.Error())
	//}

	//cmd := fmt.Sprintf(signUpAuthSql, ***)

	//_, err = db.Conn().Exec(signUpAuthSql, u.Id())

	u, err := entity.GenWhenCreateUser(in.Name, in.Password)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	cmd := fmt.Sprintf(signUpAuthSql, dbtable.TableNameUser)

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

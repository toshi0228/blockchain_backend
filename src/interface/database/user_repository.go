package database

import (
	_ "embed"
	"fmt"
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/infra/db"
	"github.com/toshi0228/blockchain/src/infra/dbtable"
	"github.com/toshi0228/blockchain/src/usecase/userusecase/input"
	"log"
	"time"
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

type UserRow struct {
	Id        uint32    `db:"id"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	// relationされたwallet情報
	WalletId          uint32    `db:"wallet_id"`
	UserID            uint32    `db:"user_id"`
	BlockchainAddress string    `db:"blockchain_address"`
	WalletCreatedAt   time.Time `db:"wallet_created_at"`
	WalletUpdatedAt   time.Time `db:"wallet_updated_at"`
}

// FindAll ユーザーの一覧取得
func (repo *UserRepositoryImpl) FindAll() ([]*entity.UserPager, error) {

	//customers := make([]userRow, 0)

	rows, err := db.Conn().Queryx(findAllUserSql)
	if err != nil {
		return nil, err
	}
	//defer rows.C()

	if rows == nil {
		return nil, nil
	}

	var users []*entity.UserPager
	for rows.Next() {

		ur := &UserRow{}

		err := rows.StructScan(ur)

		if err != nil {
			fmt.Println("エラーぶん")
			fmt.Println(err)
		}

		u, err := entity.NewUserPager(
			ur.Id,
			ur.Name,
			ur.Password,
			ur.CreatedAt,
			ur.UpdatedAt,
			ur.WalletId,
			ur.UserID,
			ur.BlockchainAddress,
			ur.WalletCreatedAt,
			ur.WalletUpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}

		log.Println("=============")
		log.Println(u.Id())
		log.Println(u.Wallet().BlockchainAddress())
		log.Println(u.Wallet().Id())
		log.Println(u.Wallet().UpdatedAt())
		users = append(users, u)
	}

	return users, nil
	//return users, nil
}

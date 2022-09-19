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

// FindAll ユーザーの一覧取得
func (repo *UserRepositoryImpl) FindAll() ([]*entity.UserPager, error) {

	rows, _ := db.Conn().Query(findAllUserSql)
	defer rows.Close()

	if rows == nil {
		return nil, nil
	}

	type userRow struct {
		id        uint32
		name      string
		password  string
		createdAt time.Time
		updatedAt time.Time

		// relationされたwallet情報
		walletId          uint32
		userID            uint32
		blockchainAddress string
		walletCreatedAt   time.Time
		walletUpdatedAt   time.Time
	}

	var users []*entity.UserPager
	for rows.Next() {

		ur := &userRow{}

		err := rows.Scan(
			&ur.id,
			&ur.name,
			&ur.password,
			&ur.createdAt,
			&ur.updatedAt,
			&ur.walletId,
			&ur.userID,
			&ur.blockchainAddress,
			&ur.walletCreatedAt,
			&ur.walletUpdatedAt,
		)

		if err != nil {
			fmt.Println("エラーぶん")
			fmt.Println(err)
		}

		u, err := entity.NewUserPager(
			ur.id,
			ur.name,
			ur.password,
			ur.createdAt,
			ur.updatedAt,
			ur.walletId,
			ur.userID,
			ur.blockchainAddress,
			ur.walletCreatedAt,
			ur.walletUpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		users = append(users, u)
	}

	return users, nil
}

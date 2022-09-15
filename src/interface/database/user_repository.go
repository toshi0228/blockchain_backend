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

//go:embed user_repository_find_all.sql
var findAllUserSql string

func (repo *UserRepositoryImpl) FindAll() ([]*entity.User, error) {

	cmd := fmt.Sprintf(findAllUserSql, dbtable.TableNameUser)

	type userRow struct {
		id        uint32
		name      string
		password  string
		createdAt time.Time
		updatedAt time.Time
	}

	rows, _ := db.Conn().Query(cmd)
	defer rows.Close()

	if rows == nil {
		return nil, nil
	}

	var users []*entity.User
	for rows.Next() {

		ur := &userRow{}
		err := rows.Scan(&ur.id, &ur.name, &ur.password, &ur.createdAt, &ur.updatedAt)
		if err != nil {
			fmt.Println("エラーぶん")
			fmt.Println(err)
		}

		u, err := entity.NewUser(ur.id, ur.name, ur.password, ur.createdAt, ur.updatedAt)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		users = append(users, u)

	}

	return users, nil
}

package database

import (
	_ "embed"
	"fmt"
	"github.com/toshi0228/blockchain/src/entity"
	"github.com/toshi0228/blockchain/src/infra/db"
	"github.com/toshi0228/blockchain/src/infra/dbtable"
	"github.com/toshi0228/blockchain/src/usecase/userusecase/input"
	"log"
)

type UserRepositoryImpl struct{}

//go:embed user_repository_create.sql
var createUserSql string

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repo UserRepositoryImpl) CreateUser(in *input.CreateUserInput) (*entity.User, error) {

	u, err := entity.NewUser(in.Name, in.Password)
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
		log.Println(u.Name())
		log.Println(err)
		return nil, err
	}

	return u, nil
}

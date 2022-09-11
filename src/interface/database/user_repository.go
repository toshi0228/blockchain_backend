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
		log.Println(err)
		return nil, err
	}

	return u, nil
}

//go:embed user_repository_find_all.sql
var findAllUserSql string

func (repo *UserRepositoryImpl) FindAll() ([]*entity.User, error) {

	cmd := fmt.Sprintf(findAllUserSql, dbtable.TableNameUser)

	type user struct {
		id        uint32
		name      string
		password  string
		createdAt time.Time
		updatedAt time.Time
	}

	rows, _ := db.Conn().Query(cmd)
	defer rows.Close()

	var uu []*user
	for rows.Next() {

		u := &user{}
		err := rows.Scan(&u.id, &u.name, &u.password, &u.createdAt, &u.updatedAt)
		if err != nil {
			fmt.Println("エラーぶん")
			fmt.Println(err)
		}
		uu = append(uu, u)
	}

	for _, u := range uu {
		fmt.Println("---------------")
		fmt.Println(u.id, u.name)
	}

	//for _, user := range users {
	//	if user.ID() == id {
	//		return user, nil
	//	}
	//}

	//U, err := entity.NewUser()

	return nil, nil
}

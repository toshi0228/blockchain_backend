package authusecase_test

import (
	"github.com/toshi0228/blockchain/src/interface/database"
	"github.com/toshi0228/blockchain/src/test"
	"github.com/toshi0228/blockchain/src/usecase/authusecase"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
	"testing"
)

var (
	in = input.SignUpInput{Name: "test1", Password: "test1"}
)

func Test_SignUpExec(tT *testing.T) {

	t := test.NewTestLib(tT)

	err := t.DBInit()
	if err != nil {
		t.NG(err)
	}

	// repository
	authRepo := database.NewAuthRepositoryImpl()
	walletRepo := database.NewWalletRepositoryImpl()

	usecase := authusecase.NewSignUp(authRepo, walletRepo)
	out, err := usecase.Exec(&in)
	if err != nil {
		t.NG(err)
	}

	tT.Run("指定したユーザーが登録されているべき", func(t *testing.T) {
		if out.User.Name != "test1" {
			//t.NG(fmt.Errorf("作成された値と違う"))
			t.Errorf("作成された値と違う")
		}
	})

	tT.Run("指定したユーザーが登録されているべき1", func(t *testing.T) {
		if out.User.Name != "test1" {
			//t.NG(fmt.Errorf("作成された値と違う"))
			t.Errorf("作成された値と違う")
		}
	})

	//t.Title("指定したユーザーが登録されているべき")
	//if out.User.Name != "test1" {
	//	t.NG(fmt.Errorf("作成された値と違う"))
	//} else {
	//	t.OK()
	//}

}

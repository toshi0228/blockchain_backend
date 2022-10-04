package authusecase_test

import (
	"github.com/toshi0228/blockchain/src/interface/database"
	"github.com/toshi0228/blockchain/src/test"
	"github.com/toshi0228/blockchain/src/usecase/authusecase"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
	"log"
	"testing"
)

var (
	in = input.SignUpInput{Name: "test1", Password: "test1"}
)

func Test_SignUpExec(tt *testing.T) {

	t := test.NewTestLib(tt)

	//DBの初期化
	err := t.DBInit()
	if err != nil {
		t.NG(err)
	}

	// 必要なrepository
	authRepo := database.NewAuthRepositoryImpl()
	walletRepo := database.NewWalletRepositoryImpl()

	//case : ユーザーの新規登録を行う
	usecase := authusecase.NewSignUp(authRepo, walletRepo)
	out, err := usecase.Exec(&in)
	if err != nil {
		t.NG(err)
	}

	tt.Run("指定したユーザーが登録されているべき", func(t *testing.T) {
		if out.User.Name != "test1" {
			//t.NG(fmt.Errorf("作成された値と違う"))
			t.Errorf("作成された値と違う")
		}
	})

	tt.Run("ウォレットが作成されているべき", func(t *testing.T) {
		log.Println(len(out.Wallet.Address))
		log.Println(len(out.Wallet.Address))
		if out.Wallet.Address == "" {
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

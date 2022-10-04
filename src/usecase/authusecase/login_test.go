package authusecase_test

import (
	"github.com/toshi0228/blockchain/src/interface/database"
	"github.com/toshi0228/blockchain/src/test"
	"github.com/toshi0228/blockchain/src/usecase/authusecase"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
	"reflect"
	"testing"
)

func TestLoginusecase_Exec(t *testing.T) {

	//テスト用の便利関数
	ttt := test.NewTestLib(t)

	//DBの初期化
	err := ttt.DBInit()
	if err != nil {
		ttt.NG(err)
	}

	// 必要なrepository
	authRepo := database.NewAuthRepositoryImpl()

	//case : ユーザーの新規登録を行う
	usecase := authusecase.NewLogin(authRepo)

	tests := []struct {
		name  string
		input *input.LoginInput
		want  string
	}{
		{
			name:  "初めてのテスト1",
			input: &input.LoginInput{Name: "test1", Password: "test1"},
			want:  "test1",
		},
	}

	//out := &output.LoginUser{
	//	User: &output.LoginU{
	//		ID:        user.Id().Value(),
	//		Name:      user.Name(),
	//		CreatedAt: user.CreatedAt().Value(),
	//		UpdatedAt: user.UpdatedAt().Value(),
	//	},
	//
	//	Wallet: &output.LoginUserWallet{
	//		ID:        user.Wallet().Id().Value(),
	//		UserID:    user.Wallet().UserID().Value(),
	//		Address:   user.Wallet().BlockchainAddress(),
	//		CreatedAt: user.Wallet().CreatedAt().Value(),
	//		UpdatedAt: user.UpdatedAt().Value(),
	//	},
	//
	//	CryptKey: &output.LoginUserCryptKey{
	//		PrivateKey: key.PrivateKeyValue(),
	//		PublicKey:  key.PublicKeyValue(),
	//	},
	//}, nil
	//
	//

	//usecase := &Loginusecase{
	//	authRepository: database.NewAuthRepositoryImpl(),
	//}
	//
	//out, err := usecase.Exec(&in)
	//if err != nil {
	//	t.NG(err)
	//}
	//

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			out, _ := usecase.Exec(tt.input)
			if !reflect.DeepEqual(out, tt.want) {
				t.Errorf("Exec() got = %v, want %v", out, tt.want)
			}

			//got, err := usecase.Exec(tt.args.in)
			//if (err != nil) != tt.wantErr {
			//	t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			//	return
			//}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Exec() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

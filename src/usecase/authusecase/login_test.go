package authusecase

import (
	"github.com/toshi0228/blockchain/src/usecase/authusecase/input"
	"github.com/toshi0228/blockchain/src/usecase/authusecase/output"
	"reflect"
	"testing"
)

func TestLoginusecase_Exec(t *testing.T) {
	type fields struct {
		authRepository IAuthRepository
	}
	type args struct {
		in *input.LoginInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *output.LoginUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			use := &Loginusecase{
				authRepository: tt.fields.authRepository,
			}
			got, err := use.Exec(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exec() got = %v, want %v", got, tt.want)
			}
		})
	}
}

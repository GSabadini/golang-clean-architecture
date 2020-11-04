package presenter

import (
	"reflect"
	"testing"
	"time"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
	"github.com/GSabadini/go-challenge/usecase"
)

func Test_createUserPresenter_Output(t *testing.T) {
	type args struct {
		u entity.User
	}
	tests := []struct {
		name string
		args args
		want usecase.CreateUserOutput
	}{
		{
			name: "Create common user output",
			args: args{
				u: entity.NewCommonUser(
					vo.NewUuidStaticTest(),
					vo.NewFullName("Test testing"),
					vo.NewEmailTest("test@testing.com"),
					vo.NewPassword("passw"),
					vo.NewDocumentTest(vo.CNPJ, "98.521.079/0001-09"),
					vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
					time.Time{},
				),
			},
			want: usecase.CreateUserOutput{
				ID:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				FullName: "Test testing",
				Email:    "test@testing.com",
				Password: "passw",
				Document: usecase.CreateUserDocumentOutput{
					Type:  "CNPJ",
					Value: "98.521.079/0001-09",
				},
				Wallet: usecase.CreateUserWalletOutput{
					Currency: "BRL",
					Amount:   100,
				},
				Roles: usecase.CreateUserRolesOutput{
					CanTransfer: true,
				},
				Type:      "COMMON",
				CreatedAt: time.Time{}.Format(time.RFC3339),
			},
		},
		{
			name: "Create merchant user output",
			args: args{
				u: entity.NewMerchantUser(
					vo.NewUuidStaticTest(),
					vo.NewFullName("Test testing"),
					vo.NewEmailTest("test@testing.com"),
					vo.NewPassword("passw"),
					vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
					vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
					time.Time{},
				),
			},
			want: usecase.CreateUserOutput{
				ID:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				FullName: "Test testing",
				Email:    "test@testing.com",
				Password: "passw",
				Document: usecase.CreateUserDocumentOutput{
					Type:  "CNPJ",
					Value: "20.770.438/0001-66",
				},
				Wallet: usecase.CreateUserWalletOutput{
					Currency: "BRL",
					Amount:   100,
				},
				Roles: usecase.CreateUserRolesOutput{
					CanTransfer: false,
				},
				Type:      "MERCHANT",
				CreatedAt: time.Time{}.Format(time.RFC3339),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCreateUserPresenter()
			if got := c.Output(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%+v' \n| Want: '%+v'", tt.name, got, tt.want)
			}
		})
	}
}

package presenter

import (
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
	"github.com/GSabadini/go-challenge/usecase"
	"reflect"
	"testing"
	"time"
)

func Test_findUserByIDPresenter_Output(t *testing.T) {
	type args struct {
		u entity.User
	}
	tests := []struct {
		name string
		args args
		want usecase.FindUserByIDOutput
	}{
		{
			name: "Create find user by id output",
			args: args{
				u: entity.NewCustomUser(
					vo.NewUuidStaticTest(),
					vo.NewFullName("Test testing"),
					vo.NewEmailTest("test@testing.com"),
					vo.NewPassword("passw"),
					vo.NewDocumentTest(vo.CNPJ, "98.521.079/0001-09"),
					vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
					time.Time{},
				),
			},
			want: usecase.FindUserByIDOutput{
				ID:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				FullName: "Test testing",
				Email:    "test@testing.com",
				Document: usecase.FindUserByIDDocumentOutput{
					Type:  "CNPJ",
					Value: "98.521.079/0001-09",
				},
				Wallet: usecase.FindUserByIDWalletOutput{
					Currency: "BRL",
					Amount:   100,
				},
				Roles: usecase.FindUserByIDRolesOutput{
					CanTransfer: true,
				},
				Type:      "CUSTOM",
				CreatedAt: time.Time{}.String(),
			},
		},
		{
			name: "Create find user by id output",
			args: args{
				u: entity.NewMerchantUser(
					vo.NewUuidStaticTest(),
					vo.NewFullName("Test testing"),
					vo.NewEmailTest("test@testing.com"),
					vo.NewPassword("passw"),
					vo.NewDocumentTest(vo.CPF, "07091054965"),
					vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
					time.Time{},
				),
			},
			want: usecase.FindUserByIDOutput{
				ID:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				FullName: "Test testing",
				Email:    "test@testing.com",
				Document: usecase.FindUserByIDDocumentOutput{
					Type:  "CPF",
					Value: "07091054965",
				},
				Wallet: usecase.FindUserByIDWalletOutput{
					Currency: "BRL",
					Amount:   100,
				},
				Roles: usecase.FindUserByIDRolesOutput{
					CanTransfer: false,
				},
				Type:      "MERCHANT",
				CreatedAt: time.Time{}.String(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFindUserByIDPresenter()
			if got := f.Output(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%+v' \n| Want: '%+v'", tt.name, got, tt.want)
			}
		})
	}
}

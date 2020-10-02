package usecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
	"github.com/pkg/errors"
)

type findUserByIDRepoStub struct {
	result entity.User
	err    error
}

func (f findUserByIDRepoStub) FindByID(_ context.Context, _ vo.Uuid) (entity.User, error) {
	return f.result, f.err
}

type findUserByIDPresenterStub struct {
	result FindUserByIDOutput
}

func (f findUserByIDPresenterStub) Output(entity.User) FindUserByIDOutput {
	return f.result
}

func TestFindUserByIDInteractor_Execute(t *testing.T) {
	type fields struct {
		repo entity.FindUserByIDRepository
		pre  FindUserByIDPresenter
	}
	type args struct {
		ctx context.Context
		ID  vo.Uuid
	}
	var tests = []struct {
		name    string
		fields  fields
		args    args
		want    FindUserByIDOutput
		wantErr bool
	}{
		{
			name: "Find custom user by id success",
			fields: fields{
				repo: findUserByIDRepoStub{
					result: entity.NewCustomUser(
						vo.NewUuidStaticTest(),
						vo.NewFullName("Custom user"),
						vo.Email{},
						vo.NewPassword("passw"),
						vo.NewDocumentTest(vo.CPF, "07091054954"),
						nil,
						time.Now(),
					),
					err: nil,
				},
				pre: findUserByIDPresenterStub{
					result: FindUserByIDOutput{
						ID:       vo.NewUuidStaticTest().Value(),
						FullName: "Custom user",
						Document: FindUserByIDDocumentOutput{
							Type:  "CPF",
							Value: "07091054954",
						},
						Email:  vo.Email{}.Value(),
						Wallet: FindUserByIDWalletOutput{},
						Type:   "CUSTOM",
					},
				},
			},
			args: args{
				ctx: nil,
				ID:  vo.NewUuidStaticTest(),
			},
			want: FindUserByIDOutput{
				ID:       vo.NewUuidStaticTest().Value(),
				FullName: "Custom user",
				Document: FindUserByIDDocumentOutput{
					Type:  "CPF",
					Value: "07091054954",
				},
				Email:  vo.Email{}.Value(),
				Wallet: FindUserByIDWalletOutput{},
				Type:   "CUSTOM",
			},
			wantErr: false,
		},
		{
			name: "Find merchant user by id success",
			fields: fields{
				repo: findUserByIDRepoStub{
					result: entity.NewMerchantUser(
						vo.NewUuidStaticTest(),
						vo.NewFullName("Merchant user"),
						vo.Email{},
						vo.NewPassword("passw"),
						vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
						nil,
						time.Now(),
					),
					err: nil,
				},
				pre: findUserByIDPresenterStub{
					result: FindUserByIDOutput{
						ID:       vo.NewUuidStaticTest().Value(),
						FullName: "Merchant user",
						Document: FindUserByIDDocumentOutput{
							Type:  "CNPJ",
							Value: "20.770.438/0001-66",
						},
						Email:  vo.Email{}.Value(),
						Wallet: FindUserByIDWalletOutput{},
						Type:   "MERCHANT",
					},
				},
			},
			args: args{
				ctx: nil,
				ID:  vo.NewUuidStaticTest(),
			},
			want: FindUserByIDOutput{
				ID:       vo.NewUuidStaticTest().Value(),
				FullName: "Merchant user",
				Document: FindUserByIDDocumentOutput{
					Type:  "CNPJ",
					Value: "20.770.438/0001-66",
				},
				Email:  vo.Email{}.Value(),
				Wallet: FindUserByIDWalletOutput{},
				Type:   "MERCHANT",
			},
			wantErr: false,
		},
		{
			name: "Find merchant user by id db error",
			fields: fields{
				repo: findUserByIDRepoStub{
					result: entity.User{},
					err:    errors.New("fail db"),
				},
				pre: findUserByIDPresenterStub{},
			},
			args: args{
				ctx: nil,
				ID:  vo.NewUuidStaticTest(),
			},
			want:    FindUserByIDOutput{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFindUserByIDInteractor(
				tt.fields.repo,
				tt.fields.pre,
			)

			got, err := f.Execute(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("[TestCase '%s'] Err: '%v' | WantErr: '%v'", tt.name, err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%+v' | Want: '%+v'", tt.name, got, tt.want)
			}
		})
	}
}

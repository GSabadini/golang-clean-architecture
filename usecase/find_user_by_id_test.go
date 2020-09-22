package usecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
)

type FindUserByIDRepoStub struct {
	entity.UserRepository

	result entity.User
	err    error
}

func (f FindUserByIDRepoStub) FindByID(ctx context.Context, uuid vo.Uuid) (entity.User, error) {
	return f.result, f.err
}

type FindUserByIDPresenterStub struct {
	result UserOutput
}

func (f FindUserByIDPresenterStub) Output(entity.User) UserOutput {
	return f.result
}

func TestFindUserByIDInteractor_Execute(t *testing.T) {
	type fields struct {
		repo entity.UserRepository
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
		want    UserOutput
		wantErr bool
	}{
		{
			name: "Create user success",
			fields: fields{
				repo: FindUserByIDRepoStub{
					result: entity.NewCustomUser(
						"0db298eb-c8e7-4829-84b7-c1036b4f0791",
						"Test testing",
						vo.Email{},
						"passw",
						entity.Document{
							Type:   entity.CPF,
							Number: "07091054954",
						},
						nil,
						time.Now(),
					),
					err: nil,
				},
				pre: FindUserByIDPresenterStub{
					result: UserOutput{
						ID:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
						FullName: "Test testing",
						Document: entity.Document{
							Type:   entity.CPF,
							Number: "07091054954",
						},
						Email:    vo.Email{},
						Password: "passw",
						Wallet:   nil,
						Type:     entity.CUSTOM,
					},
				},
			},
			args: args{
				ctx: nil,
				ID:  "0db298eb-c8e7-4829-84b7-c1036b4f0791",
			},
			want: UserOutput{
				ID:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				FullName: "Test testing",
				Document: entity.Document{
					Type:   entity.CPF,
					Number: "07091054954",
				},
				Email:    vo.Email{},
				Password: "passw",
				Wallet:   nil,
				Type:     entity.CUSTOM,
			},
			wantErr: false,
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

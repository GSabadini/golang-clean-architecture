package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/GSabadini/golang-clean-architecture/domain/entity"
	"github.com/GSabadini/golang-clean-architecture/domain/vo"
)

type stubTransferRepoCreator struct {
	err error
}

func (s stubTransferRepoCreator) Create(_ context.Context, _ entity.Transfer) error {
	return s.err
}

func (s stubTransferRepoCreator) WithTransaction(_ context.Context, fn func(context.Context) error) error {
	if err := fn(context.Background()); err != nil {
		return err
	}

	return nil
}

type spyUserRepoUpdater struct {
	errUpdatePayer error
	errUpdatePayee error
	invoked        bool
}

func (s *spyUserRepoUpdater) UpdateWallet(_ context.Context, _ vo.Uuid, _ vo.Money) error {
	if s.invoked == true {
		return s.errUpdatePayee
	}
	s.invoked = true

	return s.errUpdatePayer
}

type spyUserRepoFinder struct {
	findPayer func() (entity.User, error)
	findPayee func() (entity.User, error)
	invoked   bool

	err error
}

func (f *spyUserRepoFinder) FindByID(_ context.Context, _ vo.Uuid) (entity.User, error) {
	if f.invoked == true {
		return f.findPayee()
	}
	f.invoked = true

	return f.findPayer()
}

type stubAuthorizer struct {
	result bool
	err    error
}

func (s stubAuthorizer) Authorized(_ context.Context, _ entity.Transfer) (bool, error) {
	return s.result, s.err
}

type stubNotifier struct{}

func (s stubNotifier) Notify(_ context.Context, _ entity.Transfer) {}

type stubCreateTransferPresenter struct {
	result CreateTransferOutput
}

func (s stubCreateTransferPresenter) Output(_ entity.Transfer) CreateTransferOutput {
	return s.result
}

func Test_createTransferInteractor_Execute(t *testing.T) {
	type fields struct {
		repoTransferCreator entity.TransferRepositoryCreator
		repoUserUpdater     entity.UserRepositoryUpdater
		repoUserFinder      entity.UserRepositoryFinder
		pre                 CreateTransferPresenter
		authorizer          Authorizer
		notifier            Notifier
	}
	type args struct {
		i CreateTransferInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    CreateTransferOutput
		wantErr bool
	}{
		{
			name: "Create transfer success",
			fields: fields{
				repoTransferCreator: stubTransferRepoCreator{
					err: nil,
				},
				repoUserUpdater: &spyUserRepoUpdater{
					errUpdatePayer: nil,
					errUpdatePayee: nil,
				},
				repoUserFinder: &spyUserRepoFinder{
					findPayer: func() (entity.User, error) {
						return entity.NewCommonUser(
							vo.NewUuidStaticTest(),
							vo.NewFullName("Test testing"),
							vo.NewEmailTest("test@testing.com"),
							vo.NewPassword("passw"),
							vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
							vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
							time.Now(),
						), nil
					},
					findPayee: func() (entity.User, error) {
						return entity.NewMerchantUser(
							vo.NewUuidStaticTest(),
							vo.NewFullName("Merchant user"),
							vo.NewEmailTest("test@testing.com"),
							vo.NewPassword("passw"),
							vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
							vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
							time.Now(),
						), nil
					},
				},
				pre: stubCreateTransferPresenter{
					result: CreateTransferOutput{
						ID:        vo.NewUuidStaticTest().Value(),
						PayerID:   vo.NewUuidStaticTest().Value(),
						PayeeID:   vo.NewUuidStaticTest().Value(),
						Value:     100,
						CreatedAt: time.Time{}.String(),
					},
				},
				authorizer: stubAuthorizer{
					result: true,
					err:    nil,
				},
				notifier: stubNotifier{},
			},
			args: args{
				i: CreateTransferInput{
					ID:        vo.NewUuidStaticTest(),
					PayerID:   vo.NewUuidStaticTest(),
					PayeeID:   vo.NewUuidStaticTest(),
					Value:     vo.NewMoneyBRL(vo.NewAmountTest(100)),
					CreatedAt: time.Time{},
				},
			},
			want: CreateTransferOutput{
				ID:        vo.NewUuidStaticTest().Value(),
				PayerID:   vo.NewUuidStaticTest().Value(),
				PayeeID:   vo.NewUuidStaticTest().Value(),
				Value:     100,
				CreatedAt: time.Time{}.String(),
			},
			wantErr: false,
		},
		{
			name: "Create transfer error",
			fields: fields{
				repoTransferCreator: stubTransferRepoCreator{
					err: errors.New("failed create transfer"),
				},
				repoUserUpdater: &spyUserRepoUpdater{
					errUpdatePayer: nil,
					errUpdatePayee: nil,
				},
				repoUserFinder: &spyUserRepoFinder{
					findPayer: func() (entity.User, error) {
						return entity.NewCommonUser(
							vo.NewUuidStaticTest(),
							vo.NewFullName("Test testing"),
							vo.NewEmailTest("test@testing.com"),
							vo.NewPassword("passw"),
							vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
							vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
							time.Now(),
						), nil
					},
					findPayee: func() (entity.User, error) {
						return entity.NewMerchantUser(
							vo.NewUuidStaticTest(),
							vo.NewFullName("Merchant user"),
							vo.NewEmailTest("test@testing.com"),
							vo.NewPassword("passw"),
							vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
							vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
							time.Now(),
						), nil
					},
				},
				pre: stubCreateTransferPresenter{
					result: CreateTransferOutput{},
				},
				authorizer: stubAuthorizer{
					result: true,
					err:    nil,
				},
				notifier: stubNotifier{},
			},
			args: args{
				i: CreateTransferInput{
					ID:        vo.NewUuidStaticTest(),
					PayerID:   vo.NewUuidStaticTest(),
					PayeeID:   vo.NewUuidStaticTest(),
					Value:     vo.NewMoneyBRL(vo.NewAmountTest(100)),
					CreatedAt: time.Time{},
				},
			},
			want:    CreateTransferOutput{},
			wantErr: true,
		},
		{
			name: "Create transfer authorization denied error",
			fields: fields{
				pre: stubCreateTransferPresenter{
					result: CreateTransferOutput{},
				},
				authorizer: stubAuthorizer{
					result: false,
					err:    errors.New("authorization denied"),
				},
				notifier: stubNotifier{},
			},
			args: args{
				i: CreateTransferInput{
					ID:        vo.NewUuidStaticTest(),
					PayerID:   vo.NewUuidStaticTest(),
					PayeeID:   vo.NewUuidStaticTest(),
					Value:     vo.NewMoneyBRL(vo.NewAmountTest(100)),
					CreatedAt: time.Time{},
				},
			},
			want:    CreateTransferOutput{},
			wantErr: true,
		},
		{
			name: "Create transfer unauthorized user type error",
			fields: fields{
				repoTransferCreator: stubTransferRepoCreator{
					err: nil,
				},
				repoUserUpdater: &spyUserRepoUpdater{
					errUpdatePayer: nil,
					errUpdatePayee: nil,
				},
				repoUserFinder: &spyUserRepoFinder{
					findPayer: func() (entity.User, error) {
						return entity.NewMerchantUser(
							vo.NewUuidStaticTest(),
							vo.NewFullName("Merchant user"),
							vo.NewEmailTest("test@testing.com"),
							vo.NewPassword("passw"),
							vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
							vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
							time.Now(),
						), nil
					},
					findPayee: func() (entity.User, error) {
						return entity.NewCommonUser(
							vo.NewUuidStaticTest(),
							vo.NewFullName("Test testing"),
							vo.NewEmailTest("test@testing.com"),
							vo.NewPassword("passw"),
							vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
							vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
							time.Now(),
						), nil
					},
				},
				pre: stubCreateTransferPresenter{
					result: CreateTransferOutput{},
				},
				authorizer: stubAuthorizer{
					result: true,
					err:    nil,
				},
				notifier: stubNotifier{},
			},
			args: args{
				i: CreateTransferInput{
					ID:        vo.NewUuidStaticTest(),
					PayerID:   vo.NewUuidStaticTest(),
					PayeeID:   vo.NewUuidStaticTest(),
					Value:     vo.NewMoneyBRL(vo.NewAmountTest(100)),
					CreatedAt: time.Time{},
				},
			},
			want:    CreateTransferOutput{},
			wantErr: true,
		},
		{
			name: "Create transfer find payer error",
			fields: fields{
				repoTransferCreator: stubTransferRepoCreator{
					err: nil,
				},
				repoUserUpdater: &spyUserRepoUpdater{
					errUpdatePayer: nil,
					errUpdatePayee: nil,
				},
				repoUserFinder: &spyUserRepoFinder{
					findPayer: func() (entity.User, error) {
						return entity.User{}, entity.ErrNotFoundUser
					},
					findPayee: func() (entity.User, error) {
						return entity.User{}, nil
					},
				},
				pre: stubCreateTransferPresenter{
					result: CreateTransferOutput{},
				},
				authorizer: stubAuthorizer{
					result: true,
					err:    nil,
				},
				notifier: stubNotifier{},
			},
			args: args{
				i: CreateTransferInput{
					ID:        vo.NewUuidStaticTest(),
					PayerID:   vo.NewUuidStaticTest(),
					PayeeID:   vo.NewUuidStaticTest(),
					Value:     vo.Money{},
					CreatedAt: time.Time{},
				},
			},
			want:    CreateTransferOutput{},
			wantErr: true,
		},
		{
			name: "Create transfer find payee error",
			fields: fields{
				repoTransferCreator: stubTransferRepoCreator{
					err: nil,
				},
				repoUserUpdater: &spyUserRepoUpdater{
					errUpdatePayer: nil,
					errUpdatePayee: nil,
				},
				repoUserFinder: &spyUserRepoFinder{
					findPayer: func() (entity.User, error) {
						return entity.NewCommonUser(
							vo.NewUuidStaticTest(),
							vo.NewFullName("Test testing"),
							vo.NewEmailTest("test@testing.com"),
							vo.NewPassword("passw"),
							vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
							vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
							time.Now(),
						), nil
					},
					findPayee: func() (entity.User, error) {
						return entity.User{}, entity.ErrNotFoundUser
					},
				},
				pre: stubCreateTransferPresenter{
					result: CreateTransferOutput{},
				},
				authorizer: stubAuthorizer{
					result: true,
					err:    nil,
				},
				notifier: stubNotifier{},
			},
			args: args{
				i: CreateTransferInput{
					ID:        vo.NewUuidStaticTest(),
					PayerID:   vo.NewUuidStaticTest(),
					PayeeID:   vo.NewUuidStaticTest(),
					Value:     vo.Money{},
					CreatedAt: time.Time{},
				},
			},
			want:    CreateTransferOutput{},
			wantErr: true,
		},
		{
			name: "Create transfer update payer error",
			fields: fields{
				repoTransferCreator: stubTransferRepoCreator{
					err: nil,
				},
				repoUserUpdater: &spyUserRepoUpdater{
					errUpdatePayer: errors.New("failed update user"),
					errUpdatePayee: nil,
				},
				repoUserFinder: &spyUserRepoFinder{
					findPayer: func() (entity.User, error) {
						return entity.NewCommonUser(
							vo.NewUuidStaticTest(),
							vo.NewFullName("Test testing"),
							vo.NewEmailTest("test@testing.com"),
							vo.NewPassword("passw"),
							vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
							vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
							time.Now(),
						), nil
					},
					findPayee: func() (entity.User, error) {
						return entity.NewMerchantUser(
							vo.NewUuidStaticTest(),
							vo.NewFullName("Merchant user"),
							vo.NewEmailTest("test@testing.com"),
							vo.NewPassword("passw"),
							vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
							vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
							time.Now(),
						), nil
					},
				},
				pre: stubCreateTransferPresenter{
					result: CreateTransferOutput{},
				},
				authorizer: stubAuthorizer{
					result: true,
					err:    nil,
				},
				notifier: stubNotifier{},
			},
			args: args{
				i: CreateTransferInput{
					ID:        vo.NewUuidStaticTest(),
					PayerID:   vo.NewUuidStaticTest(),
					PayeeID:   vo.NewUuidStaticTest(),
					Value:     vo.NewMoneyBRL(vo.NewAmountTest(100)),
					CreatedAt: time.Time{},
				},
			},
			want:    CreateTransferOutput{},
			wantErr: true,
		},
		{
			name: "Create transfer update payee error",
			fields: fields{
				repoTransferCreator: stubTransferRepoCreator{
					err: nil,
				},
				repoUserUpdater: &spyUserRepoUpdater{
					errUpdatePayer: nil,
					errUpdatePayee: errors.New("failed update user"),
				},
				repoUserFinder: &spyUserRepoFinder{
					findPayer: func() (entity.User, error) {
						return entity.NewCommonUser(
							vo.NewUuidStaticTest(),
							vo.NewFullName("Test testing"),
							vo.NewEmailTest("test@testing.com"),
							vo.NewPassword("passw"),
							vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
							vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
							time.Now(),
						), nil
					},
					findPayee: func() (entity.User, error) {
						return entity.NewMerchantUser(
							vo.NewUuidStaticTest(),
							vo.NewFullName("Merchant user"),
							vo.NewEmailTest("test@testing.com"),
							vo.NewPassword("passw"),
							vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
							vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
							time.Now(),
						), nil
					},
				},
				pre: stubCreateTransferPresenter{
					result: CreateTransferOutput{},
				},
				authorizer: stubAuthorizer{
					result: true,
					err:    nil,
				},
				notifier: stubNotifier{},
			},
			args: args{
				i: CreateTransferInput{
					ID:        vo.NewUuidStaticTest(),
					PayerID:   vo.NewUuidStaticTest(),
					PayeeID:   vo.NewUuidStaticTest(),
					Value:     vo.NewMoneyBRL(vo.NewAmountTest(100)),
					CreatedAt: time.Time{},
				},
			},
			want:    CreateTransferOutput{},
			wantErr: true,
		},
		{
			name: "Create transfer user does not have sufficient balance error",
			fields: fields{
				repoTransferCreator: stubTransferRepoCreator{
					err: nil,
				},
				repoUserUpdater: &spyUserRepoUpdater{
					errUpdatePayer: nil,
					errUpdatePayee: nil,
				},
				repoUserFinder: &spyUserRepoFinder{
					findPayer: func() (entity.User, error) {
						return entity.NewCommonUser(
							vo.NewUuidStaticTest(),
							vo.NewFullName("Test testing"),
							vo.NewEmailTest("test@testing.com"),
							vo.NewPassword("passw"),
							vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
							vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
							time.Now(),
						), nil
					},
					findPayee: func() (entity.User, error) {
						return entity.NewMerchantUser(
							vo.NewUuidStaticTest(),
							vo.NewFullName("Merchant user"),
							vo.NewEmailTest("test@testing.com"),
							vo.NewPassword("passw"),
							vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
							vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
							time.Now(),
						), nil
					},
				},
				pre: stubCreateTransferPresenter{
					result: CreateTransferOutput{},
				},
				authorizer: stubAuthorizer{
					result: true,
					err:    nil,
				},
				notifier: stubNotifier{},
			},
			args: args{
				i: CreateTransferInput{
					ID:        vo.NewUuidStaticTest(),
					PayerID:   vo.NewUuidStaticTest(),
					PayeeID:   vo.NewUuidStaticTest(),
					Value:     vo.NewMoneyBRL(vo.NewAmountTest(150)),
					CreatedAt: time.Time{},
				},
			},
			want:    CreateTransferOutput{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCreateTransferInteractor(
				tt.fields.repoTransferCreator,
				tt.fields.repoUserUpdater,
				tt.fields.repoUserFinder,
				tt.fields.authorizer,
				tt.fields.notifier,
				tt.fields.pre,
			)

			got, err := c.Execute(context.Background(), tt.args.i)
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

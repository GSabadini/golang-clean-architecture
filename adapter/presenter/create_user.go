package presenter

import (
	"github.com/GSabadini/golang-clean-architecture/domain/entity"
	"github.com/GSabadini/golang-clean-architecture/usecase"
	"time"
)

type createUserPresenter struct{}

// NewCreateUserPresenter creates new createUserPresenter
func NewCreateUserPresenter() usecase.CreateUserPresenter {
	return createUserPresenter{}
}

// Output returns the user creation response
func (c createUserPresenter) Output(u entity.User) usecase.CreateUserOutput {
	return usecase.CreateUserOutput{
		ID:       u.ID().Value(),
		FullName: u.FullName().Value(),
		Password: u.Password().Value(),
		Email:    u.Email().Value(),
		Document: usecase.CreateUserDocumentOutput{
			Type:  u.Document().Type().String(),
			Value: u.Document().Value(),
		},
		Wallet: usecase.CreateUserWalletOutput{
			Currency: u.Wallet().Money().Currency().String(),
			Amount:   u.Wallet().Money().Amount().Value(),
		},
		Roles: usecase.CreateUserRolesOutput{
			CanTransfer: u.Roles().CanTransfer,
		},
		Type:      u.TypeUser().String(),
		CreatedAt: u.CreatedAt().Format(time.RFC3339),
	}
}

package presenter

import (
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/usecase"
)

type FindUserByIDPresenter struct{}

func (f FindUserByIDPresenter) Output(u entity.User) usecase.FindUserByIDOutput {
	return usecase.FindUserByIDOutput{
		ID:       u.ID().Value(),
		FullName: u.FullName().Value(),
		Email:    u.Email().Value(),
		Document: usecase.FindUserByIDDocumentOutput{
			Type:  u.Document().Type().String(),
			Value: u.Document().Value(),
		},
		Wallet: usecase.FindUserByIDWalletOutput{
			Currency: u.Wallet().Money().Currency().String(),
			Amount:   u.Wallet().Money().Amount().Value(),
		},
		Roles: usecase.FindUserByIDRolesOutput{
			CanTransfer: u.Roles().CanTransfer,
		},
		Type: u.TypeUser().String(),
	}
}

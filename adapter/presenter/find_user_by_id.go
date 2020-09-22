package presenter

import (
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/usecase"
)

type FindUserByIDPresenter struct{}

func (f FindUserByIDPresenter) Output(u entity.User) usecase.FindUserByIDOutput {
	return usecase.FindUserByIDOutput{
		ID:       u.ID(),
		FullName: u.FullName(),
		Document: u.Document(),
		Email:    u.Email(),
		Password: u.Password(),
		Wallet:   u.Wallet(),
		Type:     u.TypeUser(),
	}
}

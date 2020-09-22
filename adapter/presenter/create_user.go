package presenter

import (
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/usecase"
)

type CreateUserPresenter struct{}

func (c CreateUserPresenter) Output(u entity.User) usecase.CreateUserOutput {
	return usecase.CreateUserOutput{
		ID:       u.ID(),
		FullName: u.FullName(),
		Document: u.Document(),
		Email:    u.Email(),
		Password: u.Password(),
		Wallet:   u.Wallet(),
		Type:     u.TypeUser(),
	}
}

package presenter

import (
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/usecase"
)

type CreateUserPresenter struct{}

func (c CreateUserPresenter) Output(u entity.User) usecase.CreateUserOutput {
	return usecase.CreateUserOutput{
		ID:       string(u.ID()),
		FullName: string(u.FullName()),
		Document: usecase.DocumentUserOutput{
			Type:   string(u.Document().Type),
			Number: u.Document().Number,
		},
		Email:    u.Email().String(),
		Password: string(u.Password()),
		Wallet:   u.Wallet(),
		Type:     string(u.TypeUser()),
	}
}

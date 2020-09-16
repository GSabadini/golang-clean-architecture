package usecase

import (
	"context"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
)

type UserInput struct {
	ID       vo.Uuid         `json:"id"`
	FullName vo.FullName     `json:"full_name"`
	Document entity.Document `json:"document"`
	Email    vo.Email        `json:"email"`
	Password vo.Password     `json:"password"`
	Wallet   *entity.Wallet  `json:"wallet"`
	Type     entity.TypeUser `json:"type"`
}

type CreateUserUseCase interface {
	Execute(UserInput) error
}

type createUserInteractor struct {
	repo entity.UserRepository
}

func NewCreateUserInteractor(repo entity.UserRepository) createUserInteractor {
	return createUserInteractor{
		repo: repo,
	}
}

func (c createUserInteractor) Execute(ctx context.Context, i UserInput) error {
	var u, err = entity.NewUser(
		i.ID,
		i.FullName,
		i.Email,
		i.Password,
		i.Document,
		i.Wallet,
		i.Type,
	)
	if err != nil {
		return err
	}

	return c.repo.Save(ctx, u)
}

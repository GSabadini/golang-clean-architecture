package usecase

import (
	"context"
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
)

type UserInput struct {
	ID       vo.Uuid         `json:"id"`
	FullName vo.FullName     `json:"full_name"`
	Document vo.Document     `json:"document"`
	Email    vo.Email        `json:"email"`
	Password vo.Password     `json:"password"`
	Wallet   vo.Money        `json:"wallet"`
	Type     entity.TypeUser `json:"type"`
	//CreatedAt time.Time       `json:"created_at"`
}

type CreateUserUseCase interface {
	Execute(input UserInput) error
}

type CreateUserInteractor struct {
	Repo entity.UserRepository
}

func (c CreateUserInteractor) Execute(ctx context.Context, i UserInput) error {
	return c.Repo.Save(ctx, entity.NewUserFactory(
		i.ID,
		i.FullName,
		i.Email,
		i.Password,
		i.Document,
		i.Wallet,
		i.Type,
	))
}

package usecase

import (
	"context"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
)

//Output port
type FindUserByIDPresenter interface {
	Output(entity.User) UserOutput
}

//Input port
type FindUserByID interface {
	Execute(context.Context, vo.Uuid) (UserOutput, error)
}

type UserOutput struct {
	ID       vo.Uuid         `json:"id"`
	FullName vo.FullName     `json:"full_name"`
	Document entity.Document `json:"document"`
	Email    vo.Email        `json:"email"`
	Password vo.Password     `json:"password"`
	Wallet   *entity.Wallet  `json:"wallet"`
	Type     entity.TypeUser `json:"type"`
}

type FindUserByIDInteractor struct {
	repo entity.UserRepository
	pre  FindUserByIDPresenter
}

func NewFindUserByIDInteractor(repo entity.UserRepository) FindUserByIDInteractor {
	return FindUserByIDInteractor{
		repo: repo,
	}
}

func (f FindUserByIDInteractor) Execute(ctx context.Context, ID vo.Uuid) (UserOutput, error) {
	u, err := f.repo.FindByID(ctx, ID)
	if err != nil {
		return f.pre.Output(entity.User{}), err
	}

	return f.pre.Output(u), nil
}

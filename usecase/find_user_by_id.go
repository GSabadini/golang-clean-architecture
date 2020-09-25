package usecase

import (
	"context"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
)

//Output port
type FindUserByIDPresenter interface {
	Output(entity.User) FindUserByIDOutput
}

//Input port
type FindUserByID interface {
	Execute(context.Context, vo.Uuid) (FindUserByIDOutput, error)
}

//Output data
type FindUserByIDOutput struct {
	ID       vo.Uuid         `json:"id"`
	FullName vo.FullName     `json:"full_name"`
	Document entity.Document `json:"document"`
	Email    vo.Email        `json:"email"`
	Password vo.Password     `json:"password"`
	Wallet   *entity.Wallet  `json:"wallet"`
	Type     entity.TypeUser `json:"type"`
}

type FindUserByIDInteractor struct {
	repo entity.FindUserByIDRepository
	pre  FindUserByIDPresenter
}

func NewFindUserByIDInteractor(repo entity.FindUserByIDRepository, pre FindUserByIDPresenter) FindUserByIDInteractor {
	return FindUserByIDInteractor{
		repo: repo,
		pre:  pre,
	}
}

func (f FindUserByIDInteractor) Execute(ctx context.Context, ID vo.Uuid) (FindUserByIDOutput, error) {
	user, err := f.repo.FindByID(ctx, ID)
	if err != nil {
		return f.pre.Output(entity.User{}), err
	}

	return f.pre.Output(user), nil
}

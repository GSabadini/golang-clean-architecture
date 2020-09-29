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
	ID       string                     `json:"id"`
	FullName string                     `json:"full_name"`
	Document FindUserByIDDocumentOutput `json:"document"`
	Email    string                     `json:"email"`
	Wallet   FindUserByIDWalletOutput   `json:"wallet"`
	Type     string                     `json:"type"`
}

//Output data
type FindUserByIDDocumentOutput struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

//Output data
type FindUserByIDWalletOutput struct {
	Currency string `json:"currency"`
	Amount   int64  `json:"amount"`
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

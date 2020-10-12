package usecase

import (
	"context"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
)

type (
	// Output port
	FindUserByIDPresenter interface {
		Output(entity.User) FindUserByIDOutput
	}

	// Input port
	FindUserByID interface {
		Execute(context.Context, vo.Uuid) (FindUserByIDOutput, error)
	}

	// Output data
	FindUserByIDOutput struct {
		ID        string                     `json:"id"`
		FullName  string                     `json:"full_name"`
		Email     string                     `json:"email"`
		Document  FindUserByIDDocumentOutput `json:"document"`
		Wallet    FindUserByIDWalletOutput   `json:"wallet"`
		Roles     FindUserByIDRolesOutput    `json:"roles"`
		Type      string                     `json:"type"`
		CreatedAt string                     `json:"created_at"`
	}

	// Output data
	FindUserByIDDocumentOutput struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}

	// Output data
	FindUserByIDWalletOutput struct {
		Currency string `json:"currency"`
		Amount   int64  `json:"amount"`
	}

	// Output data
	FindUserByIDRolesOutput struct {
		CanTransfer bool `json:"can_transfer"`
	}

	findUserByIDInteractor struct {
		repo entity.FindUserByIDRepository
		pre  FindUserByIDPresenter
	}
)

// NewFindUserByIDInteractor creates new findUserByIDInteractor with its dependencies
func NewFindUserByIDInteractor(repo entity.FindUserByIDRepository, pre FindUserByIDPresenter) FindUserByID {
	return findUserByIDInteractor{
		repo: repo,
		pre:  pre,
	}
}

// Execute orchestrates the use case
func (f findUserByIDInteractor) Execute(ctx context.Context, ID vo.Uuid) (FindUserByIDOutput, error) {
	user, err := f.repo.FindByID(ctx, ID)
	if err != nil {
		return f.pre.Output(entity.User{}), err
	}

	return f.pre.Output(user), nil
}

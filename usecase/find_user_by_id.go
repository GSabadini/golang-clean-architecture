package usecase

import (
	"context"
	"time"

	"github.com/GSabadini/golang-clean-architecture/domain/entity"
	"github.com/GSabadini/golang-clean-architecture/domain/vo"
)

type (
	// Input port
	FindUserByIDUseCase interface {
		Execute(context.Context, FindUserByIDInput) (FindUserByIDOutput, error)
	}

	// Input data
	FindUserByIDInput struct {
		ID vo.Uuid
	}

	// Output port
	FindUserByIDPresenter interface {
		Output(entity.User) FindUserByIDOutput
	}

	// Output data
	FindUserByIDOutput struct {
		ID        string                     `json:"id"`
		FullName  string                     `json:"fullname"`
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
		repo entity.UserRepositoryFinder
		pre  FindUserByIDPresenter
	}
)

// NewFindUserByIDInteractor creates new findUserByIDInteractor with its dependencies
func NewFindUserByIDInteractor(repo entity.UserRepositoryFinder, pre FindUserByIDPresenter) FindUserByIDUseCase {
	return findUserByIDInteractor{
		repo: repo,
		pre:  pre,
	}
}

// Execute orchestrates the use case
func (f findUserByIDInteractor) Execute(ctx context.Context, i FindUserByIDInput) (FindUserByIDOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := f.repo.FindByID(ctx, i.ID)
	if err != nil {
		return f.pre.Output(entity.User{}), err
	}

	return f.pre.Output(user), nil
}

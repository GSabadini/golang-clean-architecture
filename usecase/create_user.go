package usecase

import (
	"context"
	"github.com/GSabadini/go-challenge/domain/vo"
	"time"

	"github.com/GSabadini/go-challenge/domain/entity"
)

type (
	// Input port
	CreateUserUseCase interface {
		Execute(context.Context, CreateUserInput) (CreateUserOutput, error)
	}

	// Input data
	CreateUserInput struct {
		ID        vo.Uuid
		FullName  vo.FullName
		Document  vo.Document
		Email     vo.Email
		Password  vo.Password
		Wallet    *vo.Wallet
		Type      vo.TypeUser
		CreatedAt time.Time
	}

	// Output port
	CreateUserPresenter interface {
		Output(entity.User) CreateUserOutput
	}

	// Output data
	CreateUserOutput struct {
		ID        string                   `json:"id"`
		FullName  string                   `json:"full_name"`
		Email     string                   `json:"email"`
		Password  string                   `json:"password"`
		Document  CreateUserDocumentOutput `json:"document"`
		Wallet    CreateUserWalletOutput   `json:"wallet"`
		Roles     CreateUserRolesOutput    `json:"roles"`
		Type      string                   `json:"type"`
		CreatedAt string                   `json:"created_at"`
	}

	// Output data
	CreateUserDocumentOutput struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}

	// Output data
	CreateUserWalletOutput struct {
		Currency string `json:"currency"`
		Amount   int64  `json:"amount"`
	}

	// Output data
	CreateUserRolesOutput struct {
		CanTransfer bool `json:"can_transfer"`
	}

	createUserInteractor struct {
		repo entity.UserRepositoryCreator
		pre  CreateUserPresenter
	}
)

// NewCreateUserInteractor creates new createUserInteractor with its dependencies
func NewCreateUserInteractor(repo entity.UserRepositoryCreator, pre CreateUserPresenter) CreateUserUseCase {
	return createUserInteractor{
		repo: repo,
		pre:  pre,
	}
}

// Execute orchestrates the use case
func (c createUserInteractor) Execute(ctx context.Context, i CreateUserInput) (CreateUserOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	u, err := entity.NewUser(
		i.ID,
		i.FullName,
		i.Email,
		i.Password,
		i.Document,
		i.Wallet,
		i.Type,
		i.CreatedAt,
	)
	if err != nil {
		return c.pre.Output(entity.User{}), err
	}

	user, err := c.repo.Create(ctx, u)
	if err != nil {
		return c.pre.Output(entity.User{}), err
	}

	return c.pre.Output(user), nil
}

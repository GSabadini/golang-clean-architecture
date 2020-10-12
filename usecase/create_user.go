package usecase

import (
	"context"
	"time"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
)

type (
	// Input port
	CreateUserUseCase interface {
		Execute(context.Context, CreateUserInput) (CreateUserOutput, error)
	}

	// Output port
	CreateUserPresenter interface {
		Output(entity.User) CreateUserOutput
	}

	// Input data
	CreateUserInput struct {
		FullName vo.FullName
		Document vo.Document
		Email    vo.Email
		Password vo.Password
		Wallet   *vo.Wallet
		Type     entity.TypeUser
	}

	// Output data
	CreateUserOutput struct {
		ID        string                   `json:"id"`
		FullName  string                   `json:"full_name"`
		Email     string                   `json:"email"`
		Password  string                   `json:"password"`
		Document  CreateUserDocumentOutput `json:"document"`
		Wallet    CreateUserWalletOutput   `json:"wallet"`
		Roles     CreateUserRolesOutput
		Type      string `json:"type"`
		CreatedAt string `json:"created_at"`
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
		repo entity.CreateUserRepository
		pre  CreateUserPresenter
	}
)

func NewCreateUserInput(fullName vo.FullName, document vo.Document, email vo.Email, password vo.Password, wallet *vo.Wallet, t entity.TypeUser) CreateUserInput {
	return CreateUserInput{FullName: fullName, Document: document, Email: email, Password: password, Wallet: wallet, Type: t}
}

// NewCreateUserInteractor creates new createUserInteractor with its dependencies
func NewCreateUserInteractor(repo entity.CreateUserRepository, pre CreateUserPresenter) CreateUserUseCase {
	return createUserInteractor{
		repo: repo,
		pre:  pre,
	}
}

// Execute orchestrates the use case
func (c createUserInteractor) Execute(ctx context.Context, i CreateUserInput) (CreateUserOutput, error) {
	uuid, err := vo.NewUuid(vo.CreateUuid())
	if err != nil {
		return c.pre.Output(entity.User{}), err
	}

	u, err := entity.NewUser(
		uuid,
		i.FullName,
		i.Email,
		i.Password,
		i.Document,
		i.Wallet,
		i.Type,
		time.Now(),
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

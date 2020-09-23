package usecase

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
)

//Input port
type CreateUserUseCase interface {
	Execute(CreateUserInput) (CreateUserOutput, error)
}

//Output port
type CreateUserPresenter interface {
	Output(entity.User) CreateUserOutput
}

type CreateUserInput struct {
	FullName vo.FullName
	Document entity.Document
	Email    vo.Email
	Password vo.Password
	Wallet   *entity.Wallet
	Type     entity.TypeUser
}

func NewCreateUserInput(fullName vo.FullName, document entity.Document, email vo.Email, password vo.Password, wallet *entity.Wallet, t entity.TypeUser) CreateUserInput {
	return CreateUserInput{FullName: fullName, Document: document, Email: email, Password: password, Wallet: wallet, Type: t}
}

type CreateUserOutput struct {
	ID       string             `json:"id"`
	FullName string             `json:"full_name"`
	Document DocumentUserOutput `json:"document"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Wallet   *entity.Wallet     `json:"wallet"`
	Type     string             `json:"type"`
}

type DocumentUserOutput struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type CreateUserInteractor struct {
	repo entity.CreateUserRepository
	pre  CreateUserPresenter
}

func NewCreateUserInteractor(repo entity.CreateUserRepository, pre CreateUserPresenter) CreateUserInteractor {
	return CreateUserInteractor{
		repo: repo,
		pre:  pre,
	}
}

func (c CreateUserInteractor) Execute(ctx context.Context, i CreateUserInput) (CreateUserOutput, error) {
	var u, err = entity.NewUser(
		vo.Uuid(fmt.Sprintf("i.ID-%d", rand.Intn(20))),
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

package entity

import (
	"context"
	"errors"
	"time"

	"github.com/GSabadini/go-challenge/domain/vo"
)

var (
	ErrInsufficientBalance = errors.New("user does not have sufficient balance")
)

type (
	CreateUserRepository interface {
		Create(context.Context, User) (User, error)
	}

	FindUserByIDRepository interface {
		FindByID(context.Context, vo.Uuid) (User, error)
	}

	UpdateUserWalletRepository interface {
		UpdateWallet(context.Context, vo.Uuid, vo.Money) error
	}

	User struct {
		id        vo.Uuid
		fullName  vo.FullName
		email     vo.Email
		password  vo.Password
		document  vo.Document
		wallet    *vo.Wallet
		typeUser  vo.TypeUser
		roles     vo.Roles
		createdAt time.Time
	}
)

func NewUser(
	ID vo.Uuid,
	fullName vo.FullName,
	email vo.Email,
	password vo.Password,
	document vo.Document,
	wallet *vo.Wallet,
	typeUser vo.TypeUser,
	createdAt time.Time,
) (User, error) {
	switch typeUser.ToUpper() {
	case vo.CUSTOM:
		return NewCustomUser(
			ID,
			fullName,
			email,
			password,
			document,
			wallet,
			createdAt,
		), nil
	case vo.MERCHANT:
		return NewMerchantUser(
			ID,
			fullName,
			email,
			password,
			document,
			wallet,
			createdAt,
		), nil
	}

	return User{}, vo.ErrInvalidTypeUser
}

func NewCustomUser(
	ID vo.Uuid,
	fullName vo.FullName,
	email vo.Email,
	password vo.Password,
	document vo.Document,
	wallet *vo.Wallet,
	createdAt time.Time,
) User {
	return User{
		id:       ID,
		fullName: fullName,
		document: document,
		email:    email,
		password: password,
		wallet:   wallet,
		typeUser: vo.CUSTOM,
		roles: vo.Roles{
			CanTransfer: true,
		},
		createdAt: createdAt,
	}
}

func NewMerchantUser(
	ID vo.Uuid,
	fullName vo.FullName,
	email vo.Email,
	password vo.Password,
	document vo.Document,
	wallet *vo.Wallet,
	createdAt time.Time,
) User {
	return User{
		id:       ID,
		fullName: fullName,
		document: document,
		email:    email,
		password: password,
		wallet:   wallet,
		typeUser: vo.MERCHANT,
		roles: vo.Roles{
			CanTransfer: false,
		},
		createdAt: createdAt,
	}
}

func (u User) Withdraw(money vo.Money) error {
	if u.Wallet().Money().Amount().Value() < money.Amount().Value() {
		return ErrInsufficientBalance
	}

	u.Wallet().Sub(money.Amount())

	return nil
}

func (u User) Deposit(money vo.Money) {
	u.Wallet().Add(money.Amount())
}

func (u User) CanTransfer() bool {
	return u.Roles().CanTransfer
}

func (u User) ID() vo.Uuid {
	return u.id
}

func (u User) FullName() vo.FullName {
	return u.fullName
}

func (u User) Password() vo.Password {
	return u.password
}

func (u User) Email() vo.Email {
	return u.email
}

func (u User) Roles() vo.Roles {
	return u.roles
}

func (u User) TypeUser() vo.TypeUser {
	return u.typeUser
}

func (u User) Wallet() *vo.Wallet {
	return u.wallet
}

func (u User) Document() vo.Document {
	return u.document
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

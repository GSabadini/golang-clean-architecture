package entity

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/GSabadini/go-challenge/domain/vo"
)

const (
	CUSTOM   TypeUser = "CUSTOM"
	MERCHANT TypeUser = "MERCHANT"
)

var (
	ErrInvalidTypeUser = errors.New("invalid type user")

	ErrInsufficientBalance = errors.New("user does not have sufficient balance")
)

type (
	// TypeUser define user types
	TypeUser string
)

func (t TypeUser) String() string {
	return string(t)
}

func (t TypeUser) ToUpper() TypeUser {
	return TypeUser(strings.ToUpper(string(t)))
}

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
		typeUser  TypeUser
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
	typeUser TypeUser,
	createdAt time.Time,
) (User, error) {
	switch typeUser.ToUpper() {
	case CUSTOM:
		return NewCustomUser(
			ID,
			fullName,
			email,
			password,
			document,
			wallet,
			createdAt,
		), nil
	case MERCHANT:
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

	return User{}, ErrInvalidTypeUser
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
		email:    email,
		password: password,
		document: document,
		wallet:   wallet,
		roles: vo.Roles{
			CanTransfer: true,
		},
		typeUser:  CUSTOM,
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
		email:    email,
		password: password,
		document: document,
		wallet:   wallet,
		roles: vo.Roles{
			CanTransfer: false,
		},
		typeUser:  MERCHANT,
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

func (u User) TypeUser() TypeUser {
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

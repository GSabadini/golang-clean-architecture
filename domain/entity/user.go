package entity

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/GSabadini/go-challenge/domain/vo"
)

var (
	ErrInsufficientBalance = errors.New("user does not have sufficient balance")

	ErrInvalidTypeUser = errors.New("invalid type user")
)

type TypeUser string

func (t TypeUser) toUpper() TypeUser {
	return TypeUser(strings.ToUpper(string(t)))
}

const (
	CUSTOM   TypeUser = "CUSTOM"
	MERCHANT TypeUser = "MERCHANT"
)

type CreateUserRepository interface {
	Create(context.Context, User) (User, error)
}

type FindUserByIDRepository interface {
	FindByID(context.Context, vo.Uuid) (User, error)
}

type UpdateUserWalletRepository interface {
	UpdateWallet(context.Context, vo.Uuid, vo.Money) error
}

type User struct {
	id       vo.Uuid
	fullName vo.FullName
	email    vo.Email
	password vo.Password

	document Document
	wallet   *Wallet
	typeUser TypeUser
	roles    Roles

	createdAt time.Time
}

func NewUser(
	ID vo.Uuid,
	fullName vo.FullName,
	email vo.Email,
	password vo.Password,
	document Document,
	wallet *Wallet,
	typeUser TypeUser,
	createdAt time.Time,
) (User, error) {
	switch typeUser.toUpper() {
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
	document Document,
	wallet *Wallet,
	createdAt time.Time,
) User {
	return User{
		id:       ID,
		fullName: fullName,
		document: document,
		email:    email,
		password: password,
		wallet:   wallet,
		typeUser: CUSTOM,
		roles: Roles{
			canTransfer: true,
		},
		createdAt: createdAt,
	}
}

func NewMerchantUser(
	ID vo.Uuid,
	fullName vo.FullName,
	email vo.Email,
	password vo.Password,
	document Document,
	wallet *Wallet,
	createdAt time.Time,
) User {
	return User{
		id:       ID,
		fullName: fullName,
		document: document,
		email:    email,
		password: password,
		wallet:   wallet,
		typeUser: MERCHANT,
		roles: Roles{
			canTransfer: false,
		},
		createdAt: createdAt,
	}
}

func (u User) Withdraw(money vo.Money) error {
	if u.Wallet().Money().Amount() < money.Amount() {
		return ErrInsufficientBalance
	}

	u.Wallet().Sub(money.Amount())

	return nil
}

func (u User) Deposit(money vo.Money) {
	u.Wallet().Add(money.Amount())
}

func (u User) CanTransfer() bool {
	return u.Roles().CanTransfer()
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

func (u User) Roles() Roles {
	return u.roles
}

func (u User) TypeUser() TypeUser {
	return u.typeUser
}

func (u User) Wallet() *Wallet {
	return u.wallet
}

func (u User) Document() Document {
	return u.document
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

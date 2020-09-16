package entity

import (
	"context"
	"errors"
	"time"

	"github.com/GSabadini/go-challenge/domain/vo"
)

var (
	ErrInsufficientBalance = errors.New("origin account does not have sufficient balance")
)

type UserRepository interface {
	Save(context.Context, User) error
	FindByID(context.Context, vo.Uuid) (User, error)
	UpdateWallet(context.Context, vo.Uuid, vo.Money) error
}

type Roles struct {
	canTransfer bool
}

type TypeUser string

const (
	Custom   TypeUser = "CUSTOM"
	Merchant TypeUser = "MERCHANT"
)

type User struct {
	ID       vo.Uuid
	FullName vo.FullName
	Document vo.Document
	Email    vo.Email
	Password vo.Password
	Wallet   vo.Money

	Type  TypeUser
	Roles Roles

	CreatedAt time.Time
}

func NewUserFactory(
	ID vo.Uuid,
	fullName vo.FullName,
	email vo.Email,
	password vo.Password,
	document vo.Document,
	wallet vo.Money,
	typeUser TypeUser,
) User {
	switch typeUser {
	case Custom:
		return NewCustomUser(
			ID,
			fullName,
			email,
			password,
			document,
			wallet,
		)
	case Merchant:
		return NewMerchantUser(
			ID,
			fullName,
			email,
			password,
			document,
			wallet,
		)
	}

	return User{}
}

func NewCustomUser(
	ID vo.Uuid,
	fullName vo.FullName,
	email vo.Email,
	password vo.Password,
	document vo.Document,
	wallet vo.Money,
) User {
	return User{
		ID:        ID,
		FullName:  fullName,
		Document:  document,
		Email:     email,
		Password:  password,
		Wallet:    wallet,
		Type:      Custom,
		Roles:     Roles{canTransfer: true},
		CreatedAt: time.Now(),
	}
}

func NewMerchantUser(
	ID vo.Uuid,
	fullName vo.FullName,
	email vo.Email,
	password vo.Password,
	document vo.Document,
	wallet vo.Money,
) User {
	return User{
		ID:        ID,
		FullName:  fullName,
		Document:  document,
		Email:     email,
		Password:  password,
		Wallet:    wallet,
		Type:      Merchant,
		Roles:     Roles{canTransfer: false},
		CreatedAt: time.Now(),
	}
}

func (u *User) Withdraw(amount vo.Money) error {
	if u.Wallet.Value < amount.Value {
		return ErrInsufficientBalance
	}

	u.Wallet.Value -= amount.Value

	return nil
}

func (u *User) Deposit(amount vo.Money) {
	u.Wallet.Value += amount.Value
}

func (u User) IsCanTransfer() bool {
	return u.Roles.canTransfer
}

package entity

import (
	"context"
	"errors"
	"time"

	"github.com/GSabadini/go-challenge/domain/vo"
)

type UserRepository interface {
	Save(context.Context, User) error
	FindByID(context.Context, vo.Uuid) (User, error)
	UpdateWallet(context.Context, vo.Uuid, vo.Money) error
}

type TyperUser string

const (
	Custom1   TyperUser = "CUSTOM"
	Merchant1 TyperUser = "MERCHANT"
)

type TypeUser interface {
	isCanTransfer() bool
}

type Custom struct{}

func (c Custom) isCanTransfer() bool {
	return true
}

type Merchant struct{}

func (m Merchant) isCanTransfer() bool {
	return false
}

type User struct {
	ID       vo.Uuid
	FullName vo.FullName
	Document vo.Document
	Email    vo.Email
	Password vo.Password
	Wallet   vo.Money

	Type TypeUser

	CreatedAt time.Time
}

func NewUser(
	ID vo.Uuid,
	fullName vo.FullName,
	email vo.Email,
	password vo.Password,
	document vo.Document,
	wallet vo.Money,
	t TypeUser,
	createdAt time.Time,
) User {
	return User{
		ID:        ID,
		FullName:  fullName,
		Document:  document,
		Email:     email,
		Password:  password,
		Wallet:    wallet,
		Type:      t,
		CreatedAt: createdAt,
	}
}

var (
	ErrInsufficientBalance = errors.New("origin account does not have sufficient balance")
)

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
	return u.Type.isCanTransfer()
}

package entity

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/GSabadini/go-challenge/domain/vo"
)

var (
	ErrInsufficientBalance = errors.New("origin account does not have sufficient balance")

	ErrInvalidTypeUser = errors.New("invalid type User")
)

type UserRepository interface {
	Save(context.Context, *User) error
	FindByID(context.Context, vo.Uuid) (User, error)
	UpdateWallet(context.Context, vo.Uuid, vo.Money) error
}

type roles struct {
	canTransfer bool
}

func NewRoles(canTransfer bool) roles {
	return roles{canTransfer: canTransfer}
}

func (r roles) CanTransfer() bool {
	return r.canTransfer
}

type TypeUser string

func (t TypeUser) toUpper() TypeUser {
	return TypeUser(strings.ToUpper(string(t)))
}

const (
	Custom   TypeUser = "CUSTOM"
	Merchant TypeUser = "MERCHANT"
)

type TypeDocument string

const (
	CPF  TypeDocument = "CPF"
	CNPJ TypeDocument = "CNPJ"
)

type Document struct {
	Type   TypeDocument
	Number string
}

type User struct {
	id       vo.Uuid
	fullName vo.FullName
	email    vo.Email
	password vo.Password

	document Document
	wallet  *Wallet
	typeUser TypeUser
	roles    roles

	createdAt time.Time
}

func NewUserFactory(
	ID vo.Uuid,
	fullName vo.FullName,
	email vo.Email,
	password vo.Password,
	document Document,
	wallet *Wallet,
	typeUser TypeUser,
) (*User, error) {
	switch typeUser.toUpper() {
	case Custom:
		return NewCustomUser(
			ID,
			fullName,
			email,
			password,
			document,
			wallet,
		), nil
	case Merchant:
		return NewMerchantUser(
			ID,
			fullName,
			email,
			password,
			document,
			wallet,
		), nil
	}

	return nil, ErrInvalidTypeUser
}

func NewCustomUser(
	ID vo.Uuid,
	fullName vo.FullName,
	email vo.Email,
	password vo.Password,
	document Document,
	wallet *Wallet,
) *User {
	return &User{
		id:       ID,
		fullName: fullName,
		document: document,
		email:    email,
		password: password,
		wallet:   wallet,
		typeUser: Custom,
		roles: roles{
			canTransfer: true,
		},
		createdAt: time.Now(),
	}
}

func NewMerchantUser(
	ID vo.Uuid,
	fullName vo.FullName,
	email vo.Email,
	password vo.Password,
	document Document,
	wallet *Wallet,
) *User {
	return &User{
		id:       ID,
		fullName: fullName,
		document: document,
		email:    email,
		password: password,
		wallet:   wallet,
		typeUser: Merchant,
		roles: roles{
			canTransfer: false,
		},
		createdAt: time.Now(),
	}
}

func (u *User) Withdraw(amount vo.Money) error {
	if u.wallet.money.Amount() < amount.Amount() {
		return ErrInsufficientBalance
	}

	u.wallet.NewMoney(u.wallet.SubMoney(amount.Amount()))

	return nil
}

func (u *User) Deposit(amount vo.Money) {
	u.wallet.NewMoney(u.wallet.AddMoney(amount.Amount()))
}

func (u User) CanTransfer() bool {
	return u.roles.CanTransfer()
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

func (u User) Roles() roles {
	return u.roles
}

func (u User) TypeUser() TypeUser {
	return u.typeUser
}

func (u *User) Wallet() *Wallet {
	return u.wallet
}

func (u User) Document() Document {
	return u.document
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

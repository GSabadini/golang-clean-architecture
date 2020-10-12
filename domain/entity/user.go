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

	ErrNotFoundUser = errors.New("not found user")
)

type (
	// TypeUser define user types
	TypeUser string
)

// String returns string representation of the TypeUser
func (t TypeUser) String() string {
	return string(t)
}

// ToUpper
func (t TypeUser) ToUpper() TypeUser {
	return TypeUser(strings.ToUpper(string(t)))
}

type (
	// CreateUserRepository defines the operation of creating a transfer entity
	CreateUserRepository interface {
		Create(context.Context, User) (User, error)
	}

	// FindUserByIDRepository defines the search operation for a user entity
	FindUserByIDRepository interface {
		FindByID(context.Context, vo.Uuid) (User, error)
	}

	// UpdateUserWalletRepository defines the update operation of a user entity wallet
	UpdateUserWalletRepository interface {
		UpdateWallet(context.Context, vo.Uuid, vo.Money) error
	}

	// User defines the user entity
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

// NewUser is a factory for users
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

// NewCustomUser creates new custom user
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

// NewMerchant creates new merchant user
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

// Withdraw remove value of money of wallet
func (u User) Withdraw(money vo.Money) error {
	if u.Wallet().Money().Amount().Value() < money.Amount().Value() {
		return ErrInsufficientBalance
	}

	u.Wallet().Sub(money.Amount())

	return nil
}

// Deposit add value of money of wallet
func (u User) Deposit(money vo.Money) {
	u.Wallet().Add(money.Amount())
}

// CanTransfer returns whether it is possible to transfer
func (u User) CanTransfer() bool {
	return u.Roles().CanTransfer
}

// ID returns the id property
func (u User) ID() vo.Uuid {
	return u.id
}

// FullName returns the fullName property
func (u User) FullName() vo.FullName {
	return u.fullName
}

// Password returns the password property
func (u User) Password() vo.Password {
	return u.password
}

// Email returns the email property
func (u User) Email() vo.Email {
	return u.email
}

// Roles returns the roles property
func (u User) Roles() vo.Roles {
	return u.roles
}

// TypeUser returns the typeUser property
func (u User) TypeUser() TypeUser {
	return u.typeUser
}

// Wallet returns the wallet property
func (u User) Wallet() *vo.Wallet {
	return u.wallet
}

// Document returns the document property
func (u User) Document() vo.Document {
	return u.document
}

// CreatedAt returns the createdAt property
func (u User) CreatedAt() time.Time {
	return u.createdAt
}

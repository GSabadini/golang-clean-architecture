package entity

import (
	"context"
	"errors"
	"time"

	"github.com/GSabadini/go-challenge/domain/vo"
)

var (
	ErrUserInsufficientBalance = errors.New("user does not have sufficient balance")

	ErrNotFoundUser = errors.New("not found user")

	ErrUpdateUserWallet = errors.New("error updating the value of the wallet")

	ErrCreateUser = errors.New("error creating user")

	ErrFindUserByID = errors.New("error fetching user by ID")
)

type (
	// UserRepositoryCreator defines the operation of creating a transfer entity
	UserRepositoryCreator interface {
		Create(context.Context, User) (User, error)
	}

	// UserRepositoryFinder defines the search operation for a user entity
	UserRepositoryFinder interface {
		FindByID(context.Context, vo.Uuid) (User, error)
	}

	// UserRepositoryUpdater defines the update operation of a user entity wallet
	UserRepositoryUpdater interface {
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
		typeUser  vo.TypeUser
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
	typeUser vo.TypeUser,
	createdAt time.Time,
) (User, error) {
	switch typeUser.ToUpper() {
	case vo.COMMON:
		return NewCommonUser(
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

// NewCommonUser creates new common user
func NewCommonUser(
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
		typeUser:  vo.COMMON,
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
		typeUser:  vo.MERCHANT,
		createdAt: createdAt,
	}
}

// Withdraw remove value of money of wallet
func (u User) Withdraw(money vo.Money) error {
	if u.Wallet().Money().Amount().Value() < money.Amount().Value() {
		return ErrUserInsufficientBalance
	}

	u.Wallet().Sub(money.Amount())

	return nil
}

// Deposit add value of money of wallet
func (u User) Deposit(money vo.Money) {
	u.Wallet().Add(money.Amount())
}

// CanTransfer returns whether it is possible to transfer
func (u User) CanTransfer() error {
	if u.Roles().CanTransfer {
		return nil
	}

	return vo.ErrNotAllowedTypeUser
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
func (u User) TypeUser() vo.TypeUser {
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

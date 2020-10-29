package entity

import (
	"context"
	"errors"
	"time"

	"github.com/GSabadini/go-challenge/domain/vo"
)

var (
	ErrCreateTransfer = errors.New("error creating transfer")

	ErrUnauthorizedTransfer = errors.New("unauthorized transfer")
)

type (
	// TransferRepositoryCreator defines the operation of creating a transfer entity
	TransferRepositoryCreator interface {
		Create(context.Context, Transfer) (Transfer, error)
		WithTransaction(context.Context, func(context.Context) error) error
	}

	// Transfer define the transfer entity
	Transfer struct {
		id        vo.Uuid
		payer     vo.Uuid
		payee     vo.Uuid
		value     vo.Money
		createdAt time.Time
	}
)

// NewTransfer creates new transfer
func NewTransfer(
	ID vo.Uuid,
	payerID vo.Uuid,
	payeeID vo.Uuid,
	value vo.Money,
	createdAt time.Time,
) Transfer {
	return Transfer{
		id:        ID,
		payer:     payerID,
		payee:     payeeID,
		value:     value,
		createdAt: createdAt,
	}
}

// ID returns the id property
func (t Transfer) ID() vo.Uuid {
	return t.id
}

// Payer returns the payer property
func (t Transfer) Payer() vo.Uuid {
	return t.payer
}

// Payee returns the payee property
func (t Transfer) Payee() vo.Uuid {
	return t.payee
}

// Value returns the value property
func (t Transfer) Value() vo.Money {
	return t.value
}

// CreatedAt returns the createdAt property
func (t Transfer) CreatedAt() time.Time {
	return t.createdAt
}

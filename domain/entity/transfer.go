package entity

import (
	"context"
	"time"

	"github.com/GSabadini/go-challenge/domain/vo"
)

type CreateTransferRepository interface {
	Create(context.Context, Transfer) (Transfer, error)
}

type Transfer struct {
	id        vo.Uuid
	payer     vo.Uuid
	payee     vo.Uuid
	value     vo.Money
	createdAt time.Time
}

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

func (t Transfer) ID() vo.Uuid {
	return t.id
}

func (t Transfer) Payer() vo.Uuid {
	return t.payer
}

func (t Transfer) Payee() vo.Uuid {
	return t.payee
}

func (t Transfer) Value() vo.Money {
	return t.value
}

func (t Transfer) CreatedAt() time.Time {
	return t.createdAt
}

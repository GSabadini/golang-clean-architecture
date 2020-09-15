package entity

import (
	"context"
	"time"

	"github.com/GSabadini/go-challenge/domain/vo"
)

type TransferRepository interface {
	Save(context.Context, Transfer) error
}

type Transfer struct {
	ID        vo.Uuid
	Payer     vo.Uuid
	Payee     vo.Uuid
	Value     vo.Money
	CreatedAt time.Time
}

func NewTransfer(ID vo.Uuid, payerID vo.Uuid, payeeID vo.Uuid, value vo.Money, createdAt time.Time) Transfer {
	return Transfer{
		ID:        ID,
		Payer:     payerID,
		Payee:     payeeID,
		Value:     value,
		CreatedAt: createdAt,
	}
}

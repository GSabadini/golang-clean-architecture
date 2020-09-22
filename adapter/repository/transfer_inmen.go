package repository

import (
	"context"

	"github.com/GSabadini/go-challenge/domain/entity"
)

type TransferInMen struct {
	Transfer []*entity.Transfer
}

func (t *TransferInMen) Save(_ context.Context, transfer entity.Transfer) error {
	t.Transfer = append(t.Transfer, &transfer)

	return nil
}

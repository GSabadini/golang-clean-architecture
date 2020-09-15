package repository

import (
	"context"
	"database/sql"
	"github.com/GSabadini/go-challenge/domain/entity"
)

type Transfer struct {
	db *sql.DB
}

func NewTransfer(db *sql.DB) *Transfer {
	return &Transfer{db: db}
}

func (t Transfer) Save(ctx context.Context, transfer entity.Transfer) error {
	panic("implement me")
}

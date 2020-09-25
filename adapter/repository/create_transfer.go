package repository

import (
	"context"
	"database/sql"

	"github.com/GSabadini/go-challenge/domain/entity"
)

type CreateTransferRepository struct {
	db *sql.DB
}

func NewCreateTransferRepository(db *sql.DB) CreateTransferRepository {
	return CreateTransferRepository{
		db: db,
	}
}

func (c CreateTransferRepository) Create(ctx context.Context, t entity.Transfer) (entity.Transfer, error) {
	panic("implement me")
}

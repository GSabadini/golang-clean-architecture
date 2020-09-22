package repository

import (
	"context"
	"database/sql"

	"github.com/GSabadini/go-challenge/domain/entity"
)

type UpdateUserWalletRepository struct {
	db *sql.DB
}

func NewUpdateUserWalletRepository(db *sql.DB) UpdateUserWalletRepository {
	return UpdateUserWalletRepository{
		db: db,
	}
}

func (u UpdateUserWalletRepository) UpdateWallet(ctx context.Context, transfer entity.Transfer) error {
	panic("implement me")
}

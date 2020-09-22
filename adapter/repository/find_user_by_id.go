package repository

import (
	"context"
	"database/sql"

	"github.com/GSabadini/go-challenge/domain/entity"
)

type FindUserByIDRepository struct {
	db *sql.DB
}

func NewFindUserByIDUser(db *sql.DB) FindUserByIDRepository {
	return FindUserByIDRepository{
		db: db,
	}
}
func (f FindUserByIDRepository) FindByID(ctx context.Context, u entity.User) (entity.User, error) {
	panic("implement me")
}

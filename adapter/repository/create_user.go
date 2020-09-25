package repository

import (
	"context"
	"database/sql"

	"github.com/GSabadini/go-challenge/domain/entity"
)

type CreateUserRepository struct {
	db *sql.DB
}

func NewCreateUserRepository(db *sql.DB) CreateUserRepository {
	return CreateUserRepository{
		db: db,
	}
}

func (c CreateUserRepository) Create(ctx context.Context, u entity.User) (entity.User, error) {
	panic("implement me")
}

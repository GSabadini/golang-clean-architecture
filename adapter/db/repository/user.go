package repository

import (
	"context"
	"database/sql"
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (u User) Save(ctx context.Context, user entity.User) error {
	panic("implement me")
}

func (u User) FindByID(ctx context.Context, uuid vo.Uuid) (entity.User, error) {
	panic("implement me")
}

func (u User) UpdateWallet(ctx context.Context, uuid vo.Uuid) error {
	panic("implement me")
}

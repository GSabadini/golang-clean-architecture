package usecase

import (
	"context"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
)

type FindUserByID interface {
	Execute(context.Context, vo.Uuid) (entity.User, error)
}

type FindUserByIDInteractor struct {
	Repo entity.UserRepository
}

func (f FindUserByIDInteractor) Execute(ctx context.Context, ID vo.Uuid) (entity.User, error) {
	return f.Repo.FindByID(ctx, ID)
}

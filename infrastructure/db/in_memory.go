package db

import (
	"context"
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
)

type UserInMen struct {
	users []*entity.User
}

func (u *UserInMen) Create(_ context.Context, user entity.User) (entity.User, error) {
	u.users = append(u.users, &user)

	return user, nil
}

func (u *UserInMen) FindByID(_ context.Context, ID vo.Uuid) (entity.User, error) {
	for _, user := range u.users {
		if user.ID() == ID {
			return *user, nil
		}
	}

	return entity.User{}, nil
}

func (u *UserInMen) UpdateWallet(_ context.Context, ID vo.Uuid, money vo.Money) error {
	for _, user := range u.users {
		if user.ID() == ID {
			user.Wallet().NewMoney(vo.NewMoneyBRL(money.Amount()))
		}
	}

	return nil
}

type TransferInMen struct {
	Transfer []*entity.Transfer
}

func (t *TransferInMen) Create(_ context.Context, transfer entity.Transfer) (entity.Transfer, error) {
	t.Transfer = append(t.Transfer, &transfer)

	return transfer, nil
}

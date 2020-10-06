package repository

import (
	"context"
	"github.com/GSabadini/go-challenge/domain/vo"
	"github.com/GSabadini/go-challenge/infrastructure/db"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/GSabadini/go-challenge/domain/entity"
)

type (
	updateUserWalletRepository struct {
		handler    *db.MongoHandler
		collection string
	}
)

// NewUpdateUserWalletRepository creates new updateUserWalletRepository with its dependencies
func NewUpdateUserWalletRepository(handler *db.MongoHandler) entity.UpdateUserWalletRepository {
	return updateUserWalletRepository{
		handler:    handler,
		collection: "users",
	}
}

// UpdateWallet performs updateOne into the database
func (u updateUserWalletRepository) UpdateWallet(ctx context.Context, ID vo.Uuid, money vo.Money) error {
	var (
		query  = bson.M{"id": ID.Value()}
		update = bson.M{"$set": bson.M{"wallet.amount": money.Amount().Value()}}
	)

	if _, err := u.handler.Db().Collection(u.collection).UpdateOne(ctx, query, update); err != nil {
		switch err {
		case mongo.ErrNilDocument:
			//return errors.Wrap(domain.ErrAccountNotFound, "error updating account balance")
			return errors.Wrap(err, "error updating the value of the wallet")
		default:
			return errors.Wrap(err, "error updating the value of the wallet")
		}
	}

	return nil
}

package repository

import (
	"context"

	"github.com/GSabadini/golang-clean-architecture/domain/entity"
	"github.com/GSabadini/golang-clean-architecture/domain/vo"
	"github.com/GSabadini/golang-clean-architecture/infrastructure/database"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type updateUserWalletRepository struct {
	handler    *database.MongoHandler
	collection string
}

// NewUpdateUserWalletRepository creates new updateUserWalletRepository with its dependencies
func NewUpdateUserWalletRepository(handler *database.MongoHandler) entity.UserRepositoryUpdater {
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
			return errors.Wrap(entity.ErrNotFoundUser, entity.ErrUpdateUserWallet.Error())
		default:
			return errors.Wrap(err, entity.ErrUpdateUserWallet.Error())
		}
	}

	return nil
}

package repository

import (
	"context"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/infrastructure/db"
	"github.com/pkg/errors"
)

type (
	// Bson Data
	createTransferBSON struct {
		ID        string `bson:"id"`
		PayerID   string `bson:"payer"`
		PayeeID   string `bson:"payee"`
		Value     int64  `bson:"value"`
		CreatedAt string `bson:"created_at"`
	}

	createTransferRepository struct {
		handler    *db.MongoHandler
		collection string
	}
)

// NewCreateTransferRepository creates new createTransferRepository with its dependencies
func NewCreateTransferRepository(handler *db.MongoHandler) entity.CreateTransferRepository {
	return createTransferRepository{
		handler:    handler,
		collection: "transfers",
	}
}

// Create performs insertOne into the database
func (c createTransferRepository) Create(ctx context.Context, t entity.Transfer) (entity.Transfer, error) {
	var bson = createTransferBSON{
		ID:        t.ID().Value(),
		PayerID:   t.Payer().Value(),
		PayeeID:   t.Payee().Value(),
		Value:     t.Value().Amount().Value(),
		CreatedAt: t.CreatedAt().String(),
	}

	if _, err := c.handler.Db().Collection(c.collection).InsertOne(ctx, bson); err != nil {
		return entity.Transfer{}, errors.Wrap(err, "error creating transfer")
	}

	return t, nil
}

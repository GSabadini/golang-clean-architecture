package repository

import (
	"context"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/infrastructure/db"
	"github.com/pkg/errors"
)

type (
	// Bson data
	createUserBSON struct {
		ID        string                 `bson:"id"`
		FullName  string                 `bson:"full_name"`
		Email     string                 `bson:"email"`
		Password  string                 `bson:"password"`
		Document  createUserDocumentBSON `bson:"document"`
		Wallet    createUserWalletBSON   `bson:"wallet"`
		Roles     createUserRolesBSON    `bson:"roles"`
		Type      string                 `bson:"type"`
		CreatedAt string                 `bson:"created_at"`
	}

	// Bson data
	createUserDocumentBSON struct {
		Type  string `bson:"type"`
		Value string `bson:"value"`
	}

	// Bson data
	createUserWalletBSON struct {
		Currency string `bson:"currency"`
		Amount   int64  `bson:"amount"`
	}

	// Bson data
	createUserRolesBSON struct {
		CanTransfer bool `bson:"can_transfer"`
	}

	createUserRepository struct {
		handler    *db.MongoHandler
		collection string
	}
)

// NewCreateUserRepository creates new createUserRepository with its dependencies
func NewCreateUserRepository(handler *db.MongoHandler) entity.CreateUserRepository {
	return createUserRepository{
		handler:    handler,
		collection: "users",
	}
}

// Create performs insertOne into the database
func (c createUserRepository) Create(ctx context.Context, u entity.User) (entity.User, error) {
	var bson = createUserBSON{
		ID:       u.ID().Value(),
		FullName: u.FullName().Value(),
		Document: createUserDocumentBSON{
			Type:  u.Document().Type().String(),
			Value: u.Document().Value(),
		},
		Email:    u.Email().Value(),
		Password: u.Password().Value(),
		Wallet: createUserWalletBSON{
			Currency: u.Wallet().Money().Currency().String(),
			Amount:   u.Wallet().Money().Amount().Value(),
		},
		Type: u.TypeUser().String(),
	}

	if _, err := c.handler.Db().Collection(c.collection).InsertOne(ctx, bson); err != nil {
		return entity.User{}, errors.Wrap(err, "error creating user")
	}

	return u, nil
}

package repository

import (
	"context"
	"time"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
	"github.com/GSabadini/go-challenge/infrastructure/database"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	// Bson data
	findUserByIDBSON struct {
		ID        string                   `bson:"id"`
		FullName  string                   `bson:"full_name"`
		Email     string                   `bson:"email"`
		Password  string                   `bson:"password"`
		Document  findUserByIDDocumentBSON `bson:"document"`
		Wallet    findUserByIDWalletBSON   `bson:"wallet"`
		Roles     findUserByIDRolesBSON    `bson:"roles"`
		Type      string                   `bson:"type"`
		CreatedAt time.Time                `bson:"created_at"`
	}

	// Bson data
	findUserByIDDocumentBSON struct {
		Type  string `bson:"type"`
		Value string `bson:"value"`
	}

	// Bson data
	findUserByIDWalletBSON struct {
		Currency string `bson:"currency"`
		Amount   int64  `bson:"amount"`
	}

	// Bson data
	findUserByIDRolesBSON struct {
		CanTransfer bool `bson:"can_transfer"`
	}

	findUserByIDRepository struct {
		handler    *database.MongoHandler
		collection string
	}
)

// NewFindUserByIDUserRepository creates new findUserByIDRepository with its dependencies
func NewFindUserByIDUserRepository(handler *database.MongoHandler) entity.UserRepositoryFinder {
	return findUserByIDRepository{
		handler:    handler,
		collection: "users",
	}
}

// FindByID performs findOne into the database
func (f findUserByIDRepository) FindByID(ctx context.Context, ID vo.Uuid) (entity.User, error) {
	var (
		userBSON = &findUserByIDBSON{}
		query    = bson.M{"id": ID.Value()}
	)

	var err = f.handler.Db().Collection(f.collection).
		FindOne(
			ctx,
			query,
		).Decode(userBSON)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return entity.User{}, entity.ErrNotFoundUser
		default:
			return entity.User{}, errors.Wrap(err, entity.ErrFindUserByID.Error())
		}
	}

	uuid, err := vo.NewUuid(userBSON.ID)
	if err != nil {
		return entity.User{}, err
	}

	email, err := vo.NewEmail(userBSON.Email)
	if err != nil {
		return entity.User{}, err
	}

	doc, err := vo.NewDocument(vo.TypeDocument(userBSON.Document.Type), userBSON.Document.Value)
	if err != nil {
		return entity.User{}, err
	}

	currency, err := vo.NewCurrency(userBSON.Wallet.Currency)
	if err != nil {
		return entity.User{}, err
	}

	amount, err := vo.NewAmount(userBSON.Wallet.Amount)
	if err != nil {
		return entity.User{}, err
	}

	wallet := vo.NewWallet(vo.NewMoney(currency, amount))

	u, err := entity.NewUser(
		uuid,
		vo.NewFullName(userBSON.FullName),
		email,
		vo.NewPassword(userBSON.Password),
		doc,
		wallet,
		vo.TypeUser(userBSON.Type),
		userBSON.CreatedAt,
	)
	if err != nil {
		return entity.User{}, err
	}

	return u, nil
}

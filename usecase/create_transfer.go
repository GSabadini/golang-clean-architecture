package usecase

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"time"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
)

type (
	// Authorizer port
	Authorizer interface {
		Authorized(entity.Transfer) (bool, error)
	}

	// Notifier port
	Notifier interface {
		Notify(entity.Transfer) error
	}

	// Output port
	CreateTransferPresenter interface {
		Output(entity.Transfer) CreateTransferOutput
	}

	// Input port
	CreateTransferUseCase interface {
		Execute(context.Context, CreateTransferInput) (CreateTransferOutput, error)
	}

	// Input data
	CreateTransferInput struct {
		ID        vo.Uuid
		PayerID   vo.Uuid
		PayeeID   vo.Uuid
		Value     vo.Money
		CreatedAt time.Time
	}

	// Output data
	CreateTransferOutput struct {
		ID        string `json:"id"`
		PayerID   string `json:"payer"`
		PayeeID   string `json:"payee"`
		Value     int64  `json:"value"`
		CreatedAt string `json:"created_at"`
	}

	createTransferInteractor struct {
		createTransferRepo   entity.CreateTransferRepository
		updateUserWalletRepo entity.UpdateUserWalletRepository
		findUserByIDRepo     entity.FindUserByIDRepository
		pre                  CreateTransferPresenter
		authorizer           Authorizer
		notifier             Notifier
	}
)

func NewCreateTransferInteractor(
	createTransferRepo entity.CreateTransferRepository,
	updateUserWalletRepo entity.UpdateUserWalletRepository,
	findUserByIDRepo entity.FindUserByIDRepository,
	pre CreateTransferPresenter,
	authorizer Authorizer,
	notifier Notifier,
) CreateTransferUseCase {
	return createTransferInteractor{
		createTransferRepo:   createTransferRepo,
		updateUserWalletRepo: updateUserWalletRepo,
		findUserByIDRepo:     findUserByIDRepo,
		pre:                  pre,
		authorizer:           authorizer,
		notifier:             notifier,
	}
}

func (c createTransferInteractor) Execute(ctx context.Context, i CreateTransferInput) (CreateTransferOutput, error) {
	var transfer entity.Transfer

	err := c.createTransferRepo.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		if err := c.process(sessCtx, i.PayerID, i.PayeeID, i.Value); err != nil {
			return err
		}

		uuid, err := vo.NewUuid(vo.CreateUuid())
		if err != nil {
			return err
		}

		transfer, err = c.createTransferRepo.Create(sessCtx, entity.NewTransfer(
			uuid,
			i.PayerID,
			i.PayeeID,
			i.Value,
			time.Now(),
		))
		if err != nil {
			return err
		}

		ok, err := c.authorizer.Authorized(transfer)
		if err != nil || !ok {
			return err
		}

		err = c.notifier.Notify(transfer)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return c.pre.Output(entity.Transfer{}), err
	}

	return c.pre.Output(transfer), nil
}

func (c createTransferInteractor) process(ctx context.Context, payerID vo.Uuid, payeeID vo.Uuid, value vo.Money) error {
	payer, err := c.findUserByIDRepo.FindByID(ctx, payerID)
	if err != nil {
		return err
	}

	if !payer.CanTransfer() {
		return errors.New("!authorized")
	}

	payee, err := c.findUserByIDRepo.FindByID(ctx, payeeID)
	if err != nil {
		return err
	}

	err = payer.Withdraw(value)
	if err != nil {
		return err
	}

	payee.Deposit(value)

	err = c.updateUserWalletRepo.UpdateWallet(ctx, payerID, payer.Wallet().Money())
	if err != nil {
		return err
	}

	err = c.updateUserWalletRepo.UpdateWallet(ctx, payeeID, payee.Wallet().Money())
	if err != nil {
		return err
	}

	return nil
}

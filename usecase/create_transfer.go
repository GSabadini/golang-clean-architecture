package usecase

import (
	"context"
	"time"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"

	"github.com/pkg/errors"
)

type (
	// Authorizer port
	Authorizer interface {
		Authorized(context.Context, entity.Transfer) (bool, error)
	}

	// Notifier port
	Notifier interface {
		Notify(context.Context, entity.Transfer)
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

	// Output port
	CreateTransferPresenter interface {
		Output(entity.Transfer) CreateTransferOutput
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
		repoTransferCreator entity.TransferRepositoryCreator
		repoUserUpdater     entity.UserRepositoryUpdater
		repoUserFinder      entity.UserRepositoryFinder
		pre                 CreateTransferPresenter
		authorizer          Authorizer
		notifier            Notifier
	}
)

// NewCreateTransferInteractor creates new createTransferInteractor with its dependencies
func NewCreateTransferInteractor(
	repoTransferCreator entity.TransferRepositoryCreator,
	repoUserUpdater entity.UserRepositoryUpdater,
	repoUserFinder entity.UserRepositoryFinder,
	authorizer Authorizer,
	notifier Notifier,
	pre CreateTransferPresenter,
) CreateTransferUseCase {
	return createTransferInteractor{
		repoTransferCreator: repoTransferCreator,
		repoUserUpdater:     repoUserUpdater,
		repoUserFinder:      repoUserFinder,
		authorizer:          authorizer,
		notifier:            notifier,
		pre:                 pre,
	}
}

// Execute orchestrates the use case
func (c createTransferInteractor) Execute(ctx context.Context, i CreateTransferInput) (CreateTransferOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var (
		transfer entity.Transfer
		err      error
	)

	err = c.repoTransferCreator.WithTransaction(ctx, func(sessCtx context.Context) error {
		if err = c.process(sessCtx, i.PayerID, i.PayeeID, i.Value); err != nil {
			return err
		}

		transfer, err = c.repoTransferCreator.Create(sessCtx, entity.NewTransfer(
			i.ID,
			i.PayerID,
			i.PayeeID,
			i.Value,
			i.CreatedAt,
		))
		if err != nil {
			return err
		}

		ok, err := c.authorizer.Authorized(sessCtx, transfer)
		if err != nil || !ok {
			return err
		}

		return nil
	})
	if err != nil {
		return c.pre.Output(entity.Transfer{}), err
	}

	c.notifier.Notify(ctx, transfer)

	return c.pre.Output(transfer), nil
}

func (c createTransferInteractor) process(ctx context.Context, payerID vo.Uuid, payeeID vo.Uuid, value vo.Money) error {
	payer, err := c.repoUserFinder.FindByID(ctx, payerID)
	if err != nil {
		return err
	}

	if err := payer.CanTransfer(); err != nil {
		return errors.Wrap(err, entity.ErrUnauthorizedTransfer.Error())
	}

	payee, err := c.repoUserFinder.FindByID(ctx, payeeID)
	if err != nil {
		return err
	}

	err = payer.Withdraw(value)
	if err != nil {
		return err
	}

	payee.Deposit(value)

	err = c.repoUserUpdater.UpdateWallet(ctx, payerID, payer.Wallet().Money())
	if err != nil {
		return err
	}

	err = c.repoUserUpdater.UpdateWallet(ctx, payeeID, payee.Wallet().Money())
	if err != nil {
		return err
	}

	return nil
}

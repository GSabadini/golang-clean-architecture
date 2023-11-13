package usecase

import (
	"context"
	"time"

	"github.com/GSabadini/golang-clean-architecture/domain/entity"
	"github.com/GSabadini/golang-clean-architecture/domain/vo"

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

	var transfer = entity.NewTransfer(i.ID, i.PayerID, i.PayeeID, i.Value, i.CreatedAt)

	ok, err := c.authorizer.Authorized(ctx, transfer)
	if err != nil || !ok {
		return c.pre.Output(entity.Transfer{}), err
	}

	err = c.repoTransferCreator.WithTransaction(ctx, func(sessCtx context.Context) error {
		err = c.performTransactionalFlow(sessCtx, transfer)
		if err != nil {
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

func (c createTransferInteractor) performTransactionalFlow(ctx context.Context, transfer entity.Transfer) error {
	payer, err := c.repoUserFinder.FindByID(ctx, transfer.Payer())
	if err != nil {
		return err
	}

	if err = payer.CanTransfer(); err != nil {
		return errors.Wrap(err, entity.ErrUnauthorizedTransfer.Error())
	}

	payee, err := c.repoUserFinder.FindByID(ctx, transfer.Payer())
	if err != nil {
		return err
	}

	err = payer.Withdraw(transfer.Value())
	if err != nil {
		return err
	}

	payee.Deposit(transfer.Value())

	err = c.repoUserUpdater.UpdateWallet(ctx, transfer.Payer(), payer.Wallet().Money())
	if err != nil {
		return err
	}

	err = c.repoUserUpdater.UpdateWallet(ctx, transfer.Payee(), payee.Wallet().Money())
	if err != nil {
		return err
	}

	err = c.repoTransferCreator.Create(ctx, transfer)
	if err != nil {
		return err
	}

	return nil
}

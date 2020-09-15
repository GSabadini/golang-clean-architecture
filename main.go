package main

import (
	"context"
	"fmt"
	"github.com/GSabadini/go-challenge/adapter/db/repository"
	"github.com/GSabadini/go-challenge/adapter/http"
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
	"github.com/GSabadini/go-challenge/usecase"
	"time"
)

func main() {
	email, err := vo.NewEmail("gfacina@hotmail.com")
	if err != nil {
		fmt.Println(err)
	}

	payer := entity.User{
		ID:       "0db298eb-c8e7-4829-84b7-c1036b4f0792",
		FullName: "Gabriel Sabadini Facina",
		Document: vo.Document{
			Type:   vo.RG,
			Number: "102476239",
		},
		Email:    email,
		Password: "facina123",
		Type:     entity.Custom{},
		Wallet: vo.Money{
			Currency: vo.BRL,
			Value:    100,
		},
		CreatedAt: time.Now(),
	}

	payee := entity.User{
		ID:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
		FullName: "Gabriel Facina",
		Document: vo.Document{
			Type:   vo.RG,
			Number: "102476239",
		},
		Email:    email,
		Password: "facina123",
		Type:     entity.Merchant{},
		Wallet: vo.Money{
			Currency: vo.BRL,
			Value:    100,
		},
		CreatedAt: time.Now(),
	}

	userRepo := &repository.UserInMen{}
	_ = userRepo.Save(
		context.TODO(),
		payer,
	)
	_ = userRepo.Save(
		context.TODO(),
		payee,
	)

	transferRepo := &repository.TransferInMen{}

	createTransfer := usecase.CreateTransferInteractor{
		TransferRepo: transferRepo,
		UserRepo:     userRepo,
		ExternalAuthorizer:   http.Authorizer{},
		Notifier:     http.Notifier{},
	}

	transfer, err := createTransfer.Execute(
		context.TODO(),
		payer.ID,
		payee.ID,
		vo.Money{
			Currency: vo.BRL,
			Value:    100,
		})
	if err != nil {
		fmt.Println(err)
	}

	payerR, _ := userRepo.FindByID(context.TODO(), payer.ID)
	fmt.Println(payerR, "payer \n\n")

	payeeR, _ := userRepo.FindByID(context.TODO(), payee.ID)
	fmt.Println(payeeR, "payee \n\n")

	fmt.Println(transfer, "transfer")
}

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/GSabadini/go-challenge/domain/entity"
	infrahttp "github.com/GSabadini/go-challenge/infrastructure/http"
	"io/ioutil"
	"net/http"
	"time"

	adapterhttp "github.com/GSabadini/go-challenge/adapter/http"
	"github.com/GSabadini/go-challenge/adapter/presenter"
	"github.com/GSabadini/go-challenge/domain/vo"
	"github.com/GSabadini/go-challenge/infrastructure/db"
	"github.com/GSabadini/go-challenge/usecase"
)

func main() {
	email, err := vo.NewEmail("gfacina@hotmail.com")
	if err != nil {
		panic(err)
	}

	uuid, err := vo.NewUuid(vo.CreateUuid())
	if err != nil {
		//fmt.Println(err)
		panic(err)
	}
	fmt.Print(uuid)

	payer := usecase.NewCreateUserInput(
		vo.NewFullName("Gabriel Facina"),
		vo.NewDocumentTest("CPF", "1231231231"),
		email,
		vo.NewPassword("passw"),
		vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
		entity.CUSTOM,
	)
	if err != nil {
		fmt.Println(err)
	}

	payee := usecase.NewCreateUserInput(
		vo.NewFullName("Gabriel Facina"),
		vo.NewDocumentTest("CPF", "1231231231"),
		email,
		vo.NewPassword("passw"),
		vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
		entity.MERCHANT,
	)
	if err != nil {
		fmt.Println(err)
	}

	createUserRepo := &db.UserInMen{}
	createUserUC := usecase.NewCreateUserInteractor(createUserRepo, presenter.CreateUserPresenter{})
	u1, _ := createUserUC.Execute(
		context.TODO(),
		payer,
	)

	b2, err := json.Marshal(u1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b2), "b222222222")

	u2, _ := createUserUC.Execute(
		context.TODO(),
		payee,
	)

	payerID, err := vo.NewUuid(u1.ID)
	if err != nil {
		fmt.Println(err)
	}

	payeeID, err := vo.NewUuid(u2.ID)
	if err != nil {
		fmt.Println(err)
	}

	createTransferRepo := &db.TransferInMen{}
	createTransfer := usecase.NewCreateTransferInteractor(
		createTransferRepo,
		createUserRepo,
		createUserRepo,
		presenter.CreateTransferPresenter{},
		adapterhttp.NewAuthorizer(adapterhttp.NewHTTPGetterStub(
			&http.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte(`{"message":"Autorizado"}`)))},
			nil,
		)),
		adapterhttp.NewNotifier(adapterhttp.NewHTTPGetterStub(
			&http.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte(`{"message":"Enviado"}`)))},
			nil,
		)),
	)

	transfer, err := createTransfer.Execute(
		context.TODO(),
		usecase.CreateTransferInput{
			ID:        uuid,
			PayerID:   payerID,
			PayeeID:   payeeID,
			Value:     vo.NewMoneyBRL(vo.NewAmountTest(100)),
			CreatedAt: time.Time{},
		})
	if err != nil {
		fmt.Println(err)
	}

	payerR, _ := createUserRepo.FindByID(context.TODO(), payerID)
	fmt.Println(" \n\npayer")
	fmt.Printf("%+v: ", payerR.Wallet())

	payeeR, _ := createUserRepo.FindByID(context.TODO(), payeeID)
	fmt.Println(" \n\npayee")
	fmt.Printf("%+v: ", payeeR.Wallet())

	fmt.Println("\n\ntransfer")
	fmt.Printf("%+v: ", transfer)
	b, _ := json.Marshal(transfer)
	fmt.Println(string(b))

	fmt.Println("--------------------------------------------------")
	//transfer1, err := createTransfer.Execute(
	//	context.TODO(),
	//	usecase.CreateTransferInput{
	//		ID:        uuid,
	//		PayerID:   payeeID,
	//		PayeeID:   payerID,
	//		Value:     vo.NewMoneyBRL(vo.NewAmountTest(100)),
	//		CreatedAt: time.Time{},
	//	})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//payerR1, _ := createUserRepo.FindByID(context.TODO(), payerID)
	//fmt.Println(" \n\npayer")
	//fmt.Printf("%+v: ", payerR1.Wallet())
	//
	//payeeR1, _ := createUserRepo.FindByID(context.TODO(), payeeID)
	//fmt.Println(" \n\npayee")
	//fmt.Printf("%+v: ", payeeR1.Wallet())
	//
	//fmt.Println("\n\ntransfer")
	//fmt.Printf("%+v: ", transfer1)
	//b1, _ := json.Marshal(transfer1)
	//fmt.Println(string(b1))

	auth := adapterhttp.NewAuthorizer(
		infrahttp.NewClient(
			infrahttp.NewRequest(
				infrahttp.WithRetry(3, 400*time.Millisecond, []int{http.StatusInternalServerError}),
				infrahttp.WithTimeout(5*time.Second),
			),
		),
	)

	r, err := auth.Authorized(entity.Transfer{})
	fmt.Println(r, err)

	//conn, err := db.NewMongoHandler()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//repoTr := repository.NewUpdateUserWalletRepository(conn)
	//uuid, _ = vo.NewUuid("0db298eb-c8e7-4829-84b7-c1036b4f0791")
	//err = repoTr.UpdateWallet(context.TODO(), uuid, vo.NewMoneyBRL(vo.NewAmountTest(99999)))
	//if err != nil {
	//	fmt.Println(err)
	//}
}

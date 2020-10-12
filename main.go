package main

import (
	"bytes"
	"context"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	adapterhttp "github.com/GSabadini/go-challenge/adapter/http"
	"github.com/GSabadini/go-challenge/adapter/logger"
	"github.com/GSabadini/go-challenge/adapter/presenter"
	"github.com/GSabadini/go-challenge/adapter/repository"
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
	"github.com/GSabadini/go-challenge/infrastructure/db"
	infralogger "github.com/GSabadini/go-challenge/infrastructure/logger"
	"github.com/GSabadini/go-challenge/usecase"
	"io/ioutil"
	"net/http"
	"time"
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
		vo.NewDocumentTest("CPF", "07091054954"),
		email,
		vo.NewPassword("passw"),
		vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
		entity.CUSTOM,
	)

	payee := usecase.NewCreateUserInput(
		vo.NewFullName("Gabriel Facina"),
		vo.NewDocumentTest("CPF", "07091054954"),
		email,
		vo.NewPassword("passw"),
		vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
		entity.MERCHANT,
	)

	conn, err := db.NewMongoHandler()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(conn.Db().Name())

	createUserRepo := repository.NewCreateUserRepository(conn)
	createUserUC := usecase.NewCreateUserInteractor(createUserRepo, presenter.NewCreateUserPresenter())
	u1, err := createUserUC.Execute(
		context.TODO(),
		payer,
	)
	if err != nil {
		fmt.Println(err)
	}

	b2, err := json.Marshal(u1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b2), "b222222222")

	u2, _ := createUserUC.Execute(
		context.TODO(),
		payee,
	)

	b3, err := json.Marshal(u2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b3), "b3333")

	payerID, err := vo.NewUuid(u1.ID)
	if err != nil {
		fmt.Println(err)
	}

	payeeID, err := vo.NewUuid(u2.ID)
	if err != nil {
		fmt.Println(err)
	}

	createTransferRepo := repository.NewCreateTransferRepository(conn)
	updateWalletRepo := repository.NewUpdateUserWalletRepository(conn)
	findUser := repository.NewFindUserByIDUserRepository(conn)
	createTransfer := usecase.NewCreateTransferInteractor(
		createTransferRepo,
		updateWalletRepo,
		findUser,
		presenter.NewCreateTransferPresenter(),
		adapterhttp.NewAuthorizer(adapterhttp.NewHTTPGetterStub(
			&http.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte(`{"message":"!Autorizado"}`)))},
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
		fmt.Println(err, "EEEEEER")
	}

	payerR, _ := findUser.FindByID(context.TODO(), payerID)
	fmt.Println(" \n\npayer")
	fmt.Printf("%+v: ", payerR.Wallet())

	payeeR, _ := findUser.FindByID(context.TODO(), payeeID)
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

	//auth := adapterhttp.NewAuthorizer(
	//	infrahttp.NewClient(
	//		infrahttp.NewRequest(
	//			infrahttp.WithRetry(3, 400*time.Millisecond, []int{http.StatusInternalServerError}),
	//			infrahttp.WithTimeout(5*time.Second),
	//		),
	//	),
	//)
	//
	//r, err := auth.Authorized(entity.Transfer{})
	//fmt.Println(r, err)

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

	logrus := logger.NewLoggerAdapter(infralogger.NewLogrus())

	logrus.Log().Infof("HAHAHAHA")
	logrus.Log().WithFields(logger.Fields{
		"key":         "i.key",
		"http_status": "i.httpStatus",
	}).Infof("HAUHUAHU")
}

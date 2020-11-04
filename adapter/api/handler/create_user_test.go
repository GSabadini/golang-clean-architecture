package handler

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/GSabadini/go-challenge/adapter/logger"
	"github.com/GSabadini/go-challenge/adapter/presenter"
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
	infralogger "github.com/GSabadini/go-challenge/infrastructure/logger"
	"github.com/GSabadini/go-challenge/usecase"
)

type stubCreateUserUseCase struct {
	result usecase.CreateUserOutput
	err    error
}

func (s stubCreateUserUseCase) Execute(_ context.Context, _ usecase.CreateUserInput) (usecase.CreateUserOutput, error) {
	return s.result, s.err
}

func TestCreateUserHandler_Handle(t *testing.T) {
	type fields struct {
		uc     usecase.CreateUserUseCase
		log    logger.Logger
		logKey string
	}
	type args struct {
		rawPayload []byte
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		expectedBody       string
		expectedStatusCode int
	}{
		{
			name: "Success create custom user",
			fields: fields{
				uc: stubCreateUserUseCase{
					result: presenter.NewCreateUserPresenter().Output(
						entity.NewCustomUser(
							vo.NewUuidStaticTest(),
							vo.NewFullName("Custom user"),
							vo.NewEmailTest("test@testing.com"),
							vo.NewPassword("passw"),
							vo.NewDocumentTest(vo.CPF, "07091054954"),
							vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
							time.Time{},
						)),
					err: nil,
				},
				log: infralogger.Dummy{},
			},
			args: args{
				rawPayload: []byte(`
					{
						"fullname": "Gabriel Gabriel",
						"email": "gabriel@hotmail.com",
						"password": "passw123",
						"document": {
							"type": "CPF",
							"value": "070.910.549-54"
						},
						"wallet": {
							"currency": "BRL",
							"amount": 100
						},
						"type": "custom"
					}`,
				),
			},
			expectedBody:       `{"id":"0db298eb-c8e7-4829-84b7-c1036b4f0791","full_name":"Custom user","email":"test@testing.com","password":"passw","document":{"type":"CPF","value":"07091054954"},"wallet":{"currency":"BRL","amount":100},"Roles":{"can_transfer":true},"type":"CUSTOM","created_at":"0001-01-01T00:00:00Z"}`,
			expectedStatusCode: http.StatusCreated,
		},
		{
			name: "Success create custom user",
			fields: fields{
				uc: stubCreateUserUseCase{
					result: presenter.NewCreateUserPresenter().Output(
						entity.NewMerchantUser(
							vo.NewUuidStaticTest(),
							vo.NewFullName("Custom user"),
							vo.NewEmailTest("test@testing.com"),
							vo.NewPassword("passw"),
							vo.NewDocumentTest(vo.CPF, "07091054954"),
							vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
							time.Time{},
						)),
					err: nil,
				},
				log: infralogger.Dummy{},
			},
			args: args{
				rawPayload: []byte(`
					{
						"fullname": "Gabriel Gabriel",
						"email": "gabriel@hotmail.com",
						"password": "passw123",
						"document": {
							"type": "CPF",
							"value": "070.910.549-54"
						},
						"wallet": {
							"currency": "BRL",
							"amount": 100
						},
						"type": "merchant"
					}`,
				),
			},
			expectedBody:       `{"id":"0db298eb-c8e7-4829-84b7-c1036b4f0791","full_name":"Custom user","email":"test@testing.com","password":"passw","document":{"type":"CPF","value":"07091054954"},"wallet":{"currency":"BRL","amount":100},"Roles":{"can_transfer":false},"type":"MERCHANT","created_at":"0001-01-01T00:00:00Z"}`,
			expectedStatusCode: http.StatusCreated,
		},
		{
			name: "Error create user invalid inout",
			fields: fields{
				uc:  stubCreateUserUseCase{},
				log: infralogger.Dummy{},
			},
			args: args{
				rawPayload: []byte(`
					{
						"fullname": "Gabriel Gabriel",
						"email": "gabriel",
						"password": "passw123",
						"document": {
							"type": "not exists",
							"value": "070.910.549-54"
						},
						"wallet": {
							"currency": "BRL",
							"amount": 100
						},
						"type": "not exists"
					}`,
				),
			},
			expectedBody:       `{"errors":["invalid type document","invalid email","invalid type user"]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Error create user database failed",
			fields: fields{
				uc: stubCreateUserUseCase{
					result: usecase.CreateUserOutput{},
					err:    errors.New("db_error"),
				},
				log: infralogger.Dummy{},
			},
			args: args{
				rawPayload: []byte(`
				{
						"fullname": "Gabriel Gabriel",
						"email": "gabriel@hotmail.com",
						"password": "passw123",
						"document": {
							"type": "CPF",
							"value": "070.910.549-54"
						},
						"wallet": {
							"currency": "BRL",
							"amount": 100
						},
						"type": "merchant"
					}`,
				),
			},
			expectedBody:       `{"errors":["db_error"]}`,
			expectedStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(
				http.MethodPost,
				"/users",
				bytes.NewReader(tt.args.rawPayload),
			)

			var (
				w       = httptest.NewRecorder()
				handler = NewCreateUserHandler(tt.fields.uc, tt.fields.log)
			)

			handler.Handle(w, req)

			if w.Code != tt.expectedStatusCode {
				t.Errorf(
					"[TestCase '%s'] O handler retornou um HTTP status code inesperado: retornado '%v' esperado '%v'",
					tt.name,
					w.Code,
					tt.expectedStatusCode,
				)
			}

			var result = strings.TrimSpace(w.Body.String())
			if !strings.EqualFold(result, tt.expectedBody) {
				t.Errorf(
					"[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tt.name,
					result,
					tt.expectedBody,
				)
			}
		})
	}
}

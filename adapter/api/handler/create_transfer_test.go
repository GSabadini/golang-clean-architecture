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

	"github.com/GSabadini/golang-clean-architecture/adapter/logger"
	"github.com/GSabadini/golang-clean-architecture/adapter/presenter"
	"github.com/GSabadini/golang-clean-architecture/domain/entity"
	"github.com/GSabadini/golang-clean-architecture/domain/vo"
	infralogger "github.com/GSabadini/golang-clean-architecture/infrastructure/logger"
	"github.com/GSabadini/golang-clean-architecture/usecase"
)

type stubCreateTransferUseCase struct {
	result usecase.CreateTransferOutput
	err    error
}

func (s stubCreateTransferUseCase) Execute(_ context.Context, _ usecase.CreateTransferInput) (usecase.CreateTransferOutput, error) {
	return s.result, s.err
}

func TestCreateTransferHandler_Handle(t *testing.T) {
	type fields struct {
		uc  usecase.CreateTransferUseCase
		log logger.Logger
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
			name: "Success create transfer",
			fields: fields{
				uc: stubCreateTransferUseCase{
					result: presenter.NewCreateTransferPresenter().Output(
						entity.NewTransfer(
							vo.NewUuidStaticTest(),
							vo.NewUuidStaticTest(),
							vo.NewUuidStaticTest(),
							vo.NewMoneyBRL(vo.NewAmountTest(100)),
							time.Time{},
						)),
					err: nil,
				},
				log: infralogger.Dummy{},
			},
			args: args{
				rawPayload: []byte(`
					{
						"payer_id": "0db298eb-c8e7-4829-84b7-c1036b4f0791",
						"payee_id": "0db298eb-c8e7-4829-84b7-c1036b4f0791",
						"value": 100
					}`,
				),
			},
			expectedBody:       `{"id":"0db298eb-c8e7-4829-84b7-c1036b4f0791","payer":"0db298eb-c8e7-4829-84b7-c1036b4f0791","payee":"0db298eb-c8e7-4829-84b7-c1036b4f0791","value":100,"created_at":"0001-01-01T00:00:00Z"}`,
			expectedStatusCode: http.StatusCreated,
		},
		{
			name: "Error create transfer invalid input",
			fields: fields{
				uc:  stubCreateTransferUseCase{},
				log: infralogger.Dummy{},
			},
			args: args{
				rawPayload: []byte(`
					{
						"payer_id": "0db298eb-c8e7-4829-84b7",
						"payee_id": "0db298eb-c8e7-4829-84b7",
						"value": -100
					}`,
				),
			},
			expectedBody:       `{"errors":["invalid uuid","invalid uuid","invalid amount"]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Error create transfer database failed",
			fields: fields{
				uc: stubCreateTransferUseCase{
					result: usecase.CreateTransferOutput{},
					err:    errors.New("db_error"),
				},
				log: infralogger.Dummy{},
			},
			args: args{
				rawPayload: []byte(`
					{
						"payer_id": "0db298eb-c8e7-4829-84b7-c1036b4f0791",
						"payee_id": "0db298eb-c8e7-4829-84b7-c1036b4f0791",
						"value": 100
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
				"/transfers",
				bytes.NewReader(tt.args.rawPayload),
			)

			var (
				w       = httptest.NewRecorder()
				handler = NewCreateTransferHandler(tt.fields.uc, tt.fields.log)
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

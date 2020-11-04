package handler

import (
	"context"
	"errors"
	"fmt"
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
	"github.com/gorilla/mux"
)

type stubFindUserByIDUseCase struct {
	result usecase.FindUserByIDOutput
	err    error
}

func (s stubFindUserByIDUseCase) Execute(_ context.Context, _ usecase.FindUserByIDInput) (usecase.FindUserByIDOutput, error) {
	return s.result, s.err
}

func TestFindUserByIDHandler_Execute(t *testing.T) {
	type fields struct {
		uc  usecase.FindUserByIDUseCase
		log logger.Logger
	}
	type args struct {
		ID string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		expectedBody       string
		expectedStatusCode int
	}{
		{
			name: "Success find user by id",
			fields: fields{
				uc: stubFindUserByIDUseCase{
					result: presenter.NewFindUserByIDPresenter().Output(
						entity.NewCommonUser(
							vo.NewUuidStaticTest(),
							vo.NewFullName("Common user"),
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
				ID: vo.NewUuidStaticTest().Value(),
			},
			expectedBody:       `{"id":"0db298eb-c8e7-4829-84b7-c1036b4f0791","fullname":"Common user","email":"test@testing.com","document":{"type":"CPF","value":"07091054954"},"wallet":{"currency":"BRL","amount":100},"roles":{"can_transfer":true},"type":"COMMON","created_at":"0001-01-01T00:00:00Z"}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Error find user by id invalid parameter",
			fields: fields{
				uc:  stubFindUserByIDUseCase{},
				log: infralogger.Dummy{},
			},
			args:               args{},
			expectedBody:       `{"errors":["invalid parameter"]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Error find user by id database failed",
			fields: fields{
				uc: stubFindUserByIDUseCase{
					result: usecase.FindUserByIDOutput{},
					err:    errors.New("db_error"),
				},
				log: infralogger.Dummy{},
			},
			args: args{
				ID: vo.NewUuidStaticTest().Value(),
			},
			expectedBody:       `{"errors":["db_error"]}`,
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: "Error find user by id not found",
			fields: fields{
				uc: stubFindUserByIDUseCase{
					result: usecase.FindUserByIDOutput{},
					err:    entity.ErrNotFoundUser,
				},
				log: infralogger.Dummy{},
			},
			args: args{
				ID: vo.NewUuidStaticTest().Value(),
			},
			expectedBody:       `{"errors":["not found user"]}`,
			expectedStatusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uri := fmt.Sprintf("/users/%s", tt.args.ID)
			req, _ := http.NewRequest(http.MethodGet, uri, nil)

			req = mux.SetURLVars(req, map[string]string{"user_id": tt.args.ID})

			var (
				w       = httptest.NewRecorder()
				handler = NewFindUserByIDHandler(tt.fields.uc, tt.fields.log)
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

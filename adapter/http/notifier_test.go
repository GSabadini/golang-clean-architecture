package http

import (
	"bytes"
	"errors"
	"github.com/GSabadini/go-challenge/domain/entity"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestNotifier_Notify(t *testing.T) {
	type fields struct {
		client HTTPGetter
	}
	type args struct {
		t entity.Transfer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test notify success",
			fields: fields{
				client: HTTPGetterStub{
					res: &http.Response{
						Body: ioutil.NopCloser(
							bytes.NewReader([]byte(`{"message":"Enviado"}`)),
						),
					},
					err: nil,
				},
			},
			args: args{
				t: entity.Transfer{},
			},
			wantErr: false,
		},
		{
			name: "Test notify error response",
			fields: fields{
				client: HTTPGetterStub{
					res: &http.Response{
						Body: ioutil.NopCloser(
							bytes.NewReader([]byte(`{"message":"Enviad1o"}`)),
						),
					},
					err: nil,
				},
			},
			args: args{
				t: entity.Transfer{},
			},
			wantErr: true,
		},
		{
			name: "Test notify error",
			fields: fields{
				client: HTTPGetterStub{
					res: &http.Response{},
					err: errors.New("failure client"),
				},
			},
			args: args{
				t: entity.Transfer{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNotifier(tt.fields.client)
			if err := n.Notify(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("[TestCase '%s'] Err: '%v' | WantErr: '%v'", tt.name, err, tt.wantErr)
			}
		})
	}
}

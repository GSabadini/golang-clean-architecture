package http

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/GSabadini/go-challenge/domain/entity"
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
				client: httpGetterStub{
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
				client: httpGetterStub{
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
				client: httpGetterStub{
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
			if err := n.Notify(context.TODO(), tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("[TestCase '%s'] Err: '%v' | WantErr: '%v'", tt.name, err, tt.wantErr)
			}
		})
	}
}

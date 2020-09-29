package vo

import (
	"reflect"
	"testing"
)

func TestNewCurrency(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    Currency
		wantErr bool
	}{
		{
			name: "Test new valid currency",
			args: args{
				value: "BRL",
			},
			want: Currency{
				value: BRL,
			},
			wantErr: false,
		},
		{
			name: "Test new valid currency",
			args: args{
				value: "USD",
			},
			want: Currency{
				value: USD,
			},
			wantErr: false,
		},
		{
			name: "Test new invalid currency",
			args: args{
				value: "FAKE",
			},
			want:    Currency{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCurrency(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("[TestCase '%s'] Err: '%v' | WantErr: '%v'", tt.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestCurrency_Equals(t *testing.T) {
	type fields struct {
		value string
	}
	type args struct {
		value Value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Test currency value equals",
			fields: fields{
				value: "BRL",
			},
			args: args{
				value: Currency{value: BRL},
			},
			want: true,
		},
		{
			name: "Test currency value equals",
			fields: fields{
				value: "USD",
			},
			args: args{
				value: Currency{value: USD},
			},
			want: true,
		},
		{
			name: "Test currency not value equals",
			fields: fields{
				value: "USD",
			},
			args: args{
				value: Currency{value: BRL},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currency, err := NewCurrency(tt.fields.value)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v'", tt.name, err)
				return
			}

			if got := currency.Equals(tt.args.value); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

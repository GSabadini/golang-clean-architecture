package vo

import (
	"reflect"
	"testing"
)

func TestNewAmount(t *testing.T) {
	type args struct {
		value int64
	}
	tests := []struct {
		name    string
		args    args
		want    Amount
		wantErr bool
	}{
		{
			name: "Test new valid amount",
			args: args{
				value: 100,
			},
			want: Amount{
				value: 100,
			},
			wantErr: false,
		},
		{
			name: "Test new valid amount",
			args: args{
				value: 10020,
			},
			want: Amount{
				value: 10020,
			},
			wantErr: false,
		},
		{
			name: "Test new invalid amount",
			args: args{
				value: -9999,
			},
			want:    Amount{},
			wantErr: true,
		},
		{
			name: "Test new invalid amount",
			args: args{
				value: -100,
			},
			want:    Amount{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAmount(tt.args.value)
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

func TestAmount_Equals(t *testing.T) {
	type fields struct {
		value int64
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
			name: "Amount value equals",
			fields: fields{
				value: 1050,
			},
			args: args{
				value: Amount{1050},
			},
			want: true,
		},
		{
			name: "Amount value equals",
			fields: fields{
				value: 100,
			},
			args: args{
				value: Amount{100},
			},
			want: true,
		},
		{
			name: "Amount not value equals",
			fields: fields{
				value: 100,
			},
			args: args{
				value: Amount{1},
			},
			want: false,
		},
		{
			name: "Amount not value equals",
			fields: fields{
				value: 100,
			},
			args: args{
				value: Amount{10},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amount, err := NewAmount(tt.fields.value)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v'", tt.name, err)
				return
			}

			if got := amount.Equals(tt.args.value); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestAmount_Value(t *testing.T) {
	type fields struct {
		value int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "Test get value",
			fields: fields{
				value: 1,
			},
			want: 1,
		},
		{
			name: "Test get value",
			fields: fields{
				value: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amount, err := NewAmount(tt.fields.value)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v'", tt.name, err)
				return
			}

			if got := amount.Value(); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

package vo

import (
	"reflect"
	"testing"
)

func TestNewEmail(t *testing.T) {
	type args struct {
		value string
	}

	tests := []struct {
		name    string
		args    args
		want    Email
		wantErr bool
	}{
		{
			name:    "Test new valid email",
			args:    args{value: "test@test.com"},
			want:    Email{value: "test@test.com"},
			wantErr: false,
		},
		{
			name:    "Test new valid email",
			args:    args{value: "yyyyyy@xxxxxx.com"},
			want:    Email{value: "yyyyyy@xxxxxx.com"},
			wantErr: false,
		},
		{
			name:    "Test new invalid email",
			args:    args{value: "xxxxxxxx.com"},
			wantErr: true,
		},
		{
			name:    "Test new invalid email",
			args:    args{value: "test#test.com"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEmail(tt.args.value)
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

func TestEmail_Equals(t *testing.T) {
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
			name: "email value equals",
			fields: fields{
				value: "test@test.com",
			},
			args: args{
				value: Email{"test@test.com"},
			},
			want: true,
		},
		{
			name: "email value equals",
			fields: fields{
				value: "xxxxxx@xxxxxx.com",
			},
			args: args{
				value: Email{"xxxxxx@xxxxxx.com"},
			},
			want: true,
		},
		{
			name: "email value not equals",
			fields: fields{
				value: "test@test.com",
			},
			args: args{
				value: Email{"test1@test.com"},
			},
			want: false,
		},
		{
			name: "email value not equals",
			fields: fields{
				value: "xxxx@yyyy.com",
			},
			args: args{
				value: Email{"yyyy@xxxx.com"},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := NewEmail(tt.fields.value)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v'", tt.name, err)
				return
			}

			if got := e.Equals(tt.args.value); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestEmail_Value(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get value",
			fields: fields{
				value: "test@test.com",
			},
			want: "test@test.com",
		},
		{
			name: "get value",
			fields: fields{
				value: "test@example.com",
			},
			want: "test@example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := NewEmail(tt.fields.value)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v'", tt.name, err)
				return
			}

			if got := e.Value(); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

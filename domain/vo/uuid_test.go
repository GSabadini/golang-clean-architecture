package vo

import (
	"reflect"
	"testing"
)

func TestNewUuid(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    Uuid
		wantErr bool
	}{
		{
			name: "Test new valid uuid",
			args: args{
				value: "0db298eb-c8e7-4829-84b7-c1036b4f0792",
			},
			want:    Uuid{"0db298eb-c8e7-4829-84b7-c1036b4f0792"},
			wantErr: false,
		},
		{
			name: "Test new valid uuid",
			args: args{
				value: "554cb83f-2c4f-41f6-b074-d00d21adc220",
			},
			want:    Uuid{"554cb83f-2c4f-41f6-b074-d00d21adc220"},
			wantErr: false,
		},
		{
			name: "Test new invalid uuid",
			args: args{
				value: "554cb83f-2c4f-41f6-b074",
			},
			wantErr: true,
		},
		{
			name: "Test new invalid uuid",
			args: args{
				value: "3b9f811d-5b22",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUuid(tt.args.value)
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

func TestUuid_Equals(t *testing.T) {
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
			name: "Uuid value equals",
			fields: fields{
				value: "3b9f811d-5b22-47b9-aa21-5003abb98d8b",
			},
			args: args{
				value: Uuid{
					value: "3b9f811d-5b22-47b9-aa21-5003abb98d8b",
				},
			},
			want: true,
		},
		{
			name: "Uuid value equals",
			fields: fields{
				value: "d21c2aab-3aba-46a4-9500-2140578bda92",
			},
			args: args{
				value: Uuid{
					value: "d21c2aab-3aba-46a4-9500-2140578bda92",
				},
			},
			want: true,
		},
		{
			name: "Uuid value not equals",
			fields: fields{
				value: "3b9f811d-5b22-47b9-aa21-5003abb98d8b",
			},
			args: args{
				value: Uuid{
					value: "d21c2aab-3aba-46a4-9500-2140578bda92",
				},
			},
			want: false,
		},
		{
			name: "Uuid value not equals",
			fields: fields{
				value: "9622887c-c484-4857-9ae8-29001178933e",
			},
			args: args{
				value: Uuid{
					value: "3b9f811d-5b22-47b9-aa21-5003abb98d8b",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uuid, err := NewUuid(tt.fields.value)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v'", tt.name, err)
				return
			}

			if got := uuid.Equals(tt.args.value); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestUuid_Value(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get value",
			fields: fields{
				value: "d21c2aab-3aba-46a4-9500-2140578bda92",
			},
			want: "d21c2aab-3aba-46a4-9500-2140578bda92",
		},
		{
			name: "Get value",
			fields: fields{
				value: "6bb32a18-0b08-4e05-8828-49869140959a",
			},
			want: "6bb32a18-0b08-4e05-8828-49869140959a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uuid, err := NewUuid(tt.fields.value)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v'", tt.name, err)
				return
			}

			if got := uuid.Value(); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

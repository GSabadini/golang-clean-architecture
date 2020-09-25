package vo

import (
	"reflect"
	"testing"
)

func TestNewCPF(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    CPF
		wantErr bool
	}{
		{
			name: "Test new valid cpf",
			args: args{
				value: "070.910.549-45",
			},
			want:    CPF{"070.910.549-45"},
			wantErr: false,
		},
		{
			name: "Test new valid cpf",
			args: args{
				value: "876.066.350-21",
			},
			want:    CPF{"876.066.350-21"},
			wantErr: false,
		},
		{
			name: "Test new invalid cpf",
			args: args{
				value: "070.910",
			},
			wantErr: true,
		},
		{
			name: "Test new invalid cpf",
			args: args{
				value: "549-45",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCPF(tt.args.value)
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

func TestCPF_Equals(t *testing.T) {
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
			name: "CPF value equals",
			fields: fields{
				value: "876.066.350-21",
			},
			args: args{
				value: CPF{"876.066.350-21"},
			},
			want: true,
		},
		{
			name: "CPF value equals",
			fields: fields{
				value: "664.789.720-89",
			},
			args: args{
				value: CPF{"664.789.720-89"},
			},
			want: true,
		},
		{
			name: "CPF value not equals",
			fields: fields{
				value: "876.066.350-21",
			},
			args: args{
				value: CPF{"426.423.030-63"},
			},
			want: false,
		},
		{
			name: "CPF value not equals",
			fields: fields{
				value: "876.066.350-21",
			},
			args: args{
				value: CPF{"572.398.610-40"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpf, err := NewCPF(tt.fields.value)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v'", tt.name, err)
				return
			}

			if got := cpf.Equals(tt.args.value); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestCPF_Value(t *testing.T) {
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
				value: "664.789.720-89",
			},
			want: "664.789.720-89",
		},
		{
			name: "Get value",
			fields: fields{
				value: "398.473.760-26",
			},
			want: "398.473.760-26",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpf, err := NewCPF(tt.fields.value)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v'", tt.name, err)
				return
			}

			if got := cpf.Value(); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

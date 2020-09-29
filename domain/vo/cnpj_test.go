package vo

import (
	"reflect"
	"testing"
)

func TestNewCNPJ(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    Cnpj
		wantErr bool
	}{
		{
			name: "Test new valid cnpj",
			args: args{
				value: "20.770.438/0001-66",
			},
			want:    Cnpj{"20.770.438/0001-66"},
			wantErr: false,
		},
		{
			name: "Test new valid cnpj",
			args: args{
				value: "15.412.832/0001-92",
			},
			want:    Cnpj{"15.412.832/0001-92"},
			wantErr: false,
		},
		{
			name: "Test new valid cnpj",
			args: args{
				value: "15412832000192",
			},
			want:    Cnpj{"15412832000192"},
			wantErr: false,
		},
		{
			name: "Test new invalid cnpj",
			args: args{
				value: "070910/004-45",
			},
			wantErr: true,
		},
		{
			name: "Test new invalid cnpj",
			args: args{
				value: "00000/549-45",
			},
			wantErr: true,
		},
		{
			name: "Test new invalid cnpj",
			args: args{
				value: "549454554454545454545",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCNPJ(tt.args.value)
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

func TestCNPJ_Equals(t *testing.T) {
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
			name: "Cnpj value equals",
			fields: fields{
				value: "20.770.438/0001-66",
			},
			args: args{
				value: Cnpj{"20.770.438/0001-66"},
			},
			want: true,
		},
		{
			name: "Cnpj value equals",
			fields: fields{
				value: "20770438000166",
			},
			args: args{
				value: Cnpj{"20770438000166"},
			},
			want: true,
		},
		{
			name: "Cnpj value equals",
			fields: fields{
				value: "98.521.079/0001-09",
			},
			args: args{
				value: Cnpj{"98.521.079/0001-09"},
			},
			want: true,
		},
		{
			name: "Cnpj value not equals",
			fields: fields{
				value: "15.412.832/0001-92",
			},
			args: args{
				value: Cnpj{"90.691.635/0001-75"},
			},
			want: false,
		},
		{
			name: "Cnpj value not equals",
			fields: fields{
				value: "90.691.635/0001-75",
			},
			args: args{
				value: Cnpj{"15.412.832/0001-92"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cnpj, err := NewCNPJ(tt.fields.value)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v'", tt.name, err)
				return
			}

			if got := cnpj.Equals(tt.args.value); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestCNPJ_Value(t *testing.T) {
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
				value: "90.691.635/0001-75",
			},
			want: "90.691.635/0001-75",
		},
		{
			name: "Get value",
			fields: fields{
				value: "90691635000175",
			},
			want: "90691635000175",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cnpj, err := NewCNPJ(tt.fields.value)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v'", tt.name, err)
				return
			}

			if got := cnpj.Value(); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

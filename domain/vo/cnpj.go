package vo

import (
	"errors"
	"regexp"
)

var (
	// ErrInvalidCNPJ return invalid Cnpj
	ErrInvalidCNPJ = errors.New("invalid cnpj")

	rxCNPJ = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`)
)

//Cnpj structure
type Cnpj struct {
	value string
}

// NewCNPJ create new Cnpj
func NewCNPJ(value string) (Cnpj, error) {
	var cnpj = Cnpj{value: value}

	if !cnpj.validate() {
		return Cnpj{}, ErrInvalidCNPJ
	}

	return cnpj, nil
}

func (c Cnpj) validate() bool {
	return rxCNPJ.MatchString(c.value)
}

// Value return value Cnpj
func (c Cnpj) Value() string {
	return c.value
}

// String returns string representation of the Cnpj
func (c Cnpj) String() string {
	return c.value
}

// Equals checks that two Cnpj are the same
func (c Cnpj) Equals(value Value) bool {
	o, ok := value.(Cnpj)
	return ok && c.value == o.value
}

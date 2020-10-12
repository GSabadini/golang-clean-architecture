package vo

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidCPF = errors.New("invalid cpf")

	rxCPF = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
)

// Cpf structure
type Cpf struct {
	value string
}

// NewCPF create new Cpf
func NewCPF(value string) (Cpf, error) {
	var cpf = Cpf{value: value}

	if !cpf.validate() {
		return Cpf{}, ErrInvalidCPF
	}

	return cpf, nil
}

func (c Cpf) validate() bool {
	return rxCPF.MatchString(c.value)
}

// Value return value Cpf
func (c Cpf) Value() string {
	return c.value
}

// String returns string representation of the Cpf
func (c Cpf) String() string {
	return c.value
}

// Equals checks that two Cpf are the same
func (c Cpf) Equals(value Value) bool {
	o, ok := value.(Cpf)
	return ok && c.value == o.value
}

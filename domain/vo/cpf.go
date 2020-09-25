package vo

import (
	"errors"
	"regexp"
)

var (
	// ErrInvalidCPF return invalid CPF
	ErrInvalidCPF = errors.New("invalid cpf")

	rxCPFClear = regexp.MustCompile(`[^0-9]`)

	rxCPF = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
)

//CPF structure
type CPF struct {
	value string
}

// NewCPF create new CPF
func NewCPF(value string) (CPF, error) {
	var cpf = CPF{value: value}

	if !cpf.validate() {
		return CPF{}, ErrInvalidCPF
	}

	return cpf, nil
}

func (c CPF) validate() bool {
	//cpf := c.clear()

	return rxCPF.MatchString(c.value)
}

func (c CPF) clear() string {
	return rxCPFClear.ReplaceAllString(c.value, "")
}

// Value return value CPF
func (c CPF) Value() string {
	return c.value
}

// String returns string representation of the CPF
func (c CPF) String() string {
	return c.value
}

// Equals checks that two CPF are the same
func (c CPF) Equals(value Value) bool {
	o, ok := value.(CPF)
	return ok && c.value == o.value
}

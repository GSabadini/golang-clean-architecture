package vo

import "errors"

var (
	ErrInvalidCurrency = errors.New("invalid currency")
)

type TypeCurrency string

const (
	BRL TypeCurrency = "BRL"
	USD TypeCurrency = "USD"
)

func (tc TypeCurrency) String() string {
	return string(tc)
}

type Currency struct {
	value TypeCurrency
}

func NewCurrency(value string) (Currency, error) {
	var currency = Currency{value: TypeCurrency(value)}

	if !currency.validate() {
		return Currency{}, ErrInvalidCurrency
	}

	return currency, nil
}

func (c Currency) validate() bool {
	switch c.value {
	case BRL, USD:
		return true
	}

	return false
}

// Value return value Currency
func (c Currency) Value() TypeCurrency {
	return c.value
}

// String returns string representation of the Currency
func (c Currency) String() string {
	return string(c.value)
}

// Equals checks that two Currency are the same
func (c Currency) Equals(value Value) bool {
	o, ok := value.(Currency)
	return ok && c.value == o.value
}

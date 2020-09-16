package vo

type Currency string

const (
	BRL Currency = "BRL"
)

type Money struct {
	currency Currency
	amount   int64
}

// NewMoney create new Money
func NewMoney(currency Currency, amount int64) Money {
	return Money{
		currency: currency,
		amount:   amount,
	}
}

// NewMoneyBRL create new Money with currency BRL
func NewMoneyBRL(amount int64) Money {
	return Money{
		currency: BRL,
		amount:   amount,
	}
}

// Equals checks that two Money are the same
func (m Money) Equals(value Value) bool {
	o, ok := value.(Money)
	return ok && m.amount == o.amount && m.currency == o.currency
}

// Amount return value amount
func (m Money) Amount() int64 {
	return m.amount
}

// Currency return value currency
func (m Money) Currency() Currency {
	return m.currency
}

// Add value in amount
func (m Money) Add(amount int64) Money {
	return Money{
		currency: m.currency,
		amount:   m.amount + amount,
	}
}

// Sub value in amount
func (m Money) Sub(amount int64) Money {
	return Money{
		currency: m.currency,
		amount:   m.amount - amount,
	}
}

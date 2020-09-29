package vo

// Money structure
type Money struct {
	currency Currency
	amount   Amount
}

// NewMoney create new Money
func NewMoney(currency Currency, amount Amount) Money {
	return Money{
		currency: currency,
		amount:   amount,
	}
}

// NewMoneyBRL create new Money with currency BRL
func NewMoneyBRL(amount Amount) Money {
	return Money{
		currency: Currency{value: BRL},
		amount:   amount,
	}
}

// Equals checks that two Money are the same
func (m Money) Equals(value Value) bool {
	o, ok := value.(Money)
	return ok && m.amount == o.amount && m.currency == o.currency
}

// Amount return value Amount
func (m Money) Amount() Amount {
	return m.amount
}

// Currency return value currency
func (m Money) Currency() Currency {
	return m.currency
}

// Add value in Amount
func (m Money) Add(amount Amount) Money {
	return Money{
		currency: m.currency,
		amount:   Amount{value: m.amount.Value() + amount.Value()},
	}
}

// Sub value in Amount
func (m Money) Sub(amount Amount) Money {
	return Money{
		currency: m.currency,
		amount:   Amount{value: m.amount.Value() - amount.Value()},
	}
}

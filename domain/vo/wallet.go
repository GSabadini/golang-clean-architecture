package vo

//  Wallet structure
type Wallet struct {
	money Money
}

// Currency return value currency
func NewWallet(money Money) *Wallet {
	return &Wallet{money: money}
}

// Money return value money
func (w *Wallet) Money() Money {
	return w.money
}

// Add value in money value amount
func (w *Wallet) Add(amount Amount) Money {
	w.money = w.money.Add(amount)
	return w.money
}

// Sub value in money value amount
func (w *Wallet) Sub(amount Amount) Money {
	w.money = w.money.Sub(amount)
	return w.money
}

// Equals checks that two Wallet are the same
func (w *Wallet) Equals(value Value) bool {
	o, ok := value.(*Wallet)
	return ok && w.money == o.money
}

func (w *Wallet) NewMoney(money Money) {
	w.money = money
}

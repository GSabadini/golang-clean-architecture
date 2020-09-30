package vo

//Wallet structure
type Wallet struct {
	money Money
}

func NewWallet(money Money) *Wallet {
	return &Wallet{money: money}
}

func (w *Wallet) NewMoney(money Money) {
	w.money = money
}

func (w *Wallet) Add(amount Amount) Money {
	w.money = w.money.Add(amount)
	return w.money
}

func (w *Wallet) Sub(amount Amount) Money {
	w.money = w.money.Sub(amount)
	return w.money.Sub(amount)
}

func (w *Wallet) Money() Money {
	return w.money
}

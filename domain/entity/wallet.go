package entity

import "github.com/GSabadini/go-challenge/domain/vo"

type Wallet struct {
	money vo.Money
}

func NewWallet(money vo.Money) *Wallet {
	return &Wallet{money: money}
}

func (w *Wallet) NewMoney(money vo.Money) {
	w.money = money
}

func (w *Wallet) Add(amount vo.Amount) vo.Money {
	w.money = w.money.Add(amount)
	return w.money
}

func (w *Wallet) Sub(amount vo.Amount) vo.Money {
	w.money = w.money.Sub(amount)
	return w.money.Sub(amount)
}

func (w *Wallet) Money() vo.Money {
	return w.money
}

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

func (w *Wallet) AddMoney(amount int64) vo.Money {
	return w.money.Add(amount)
}

func (w *Wallet) SubMoney(amount int64) vo.Money {
	return w.money.Add(amount)
}

func (w *Wallet) Money() vo.Money {
	return w.money
}

package ptr

import (
	"errors"
	"fmt"
)

const (
	ErrInsufficientFunds = "withdrawal error: insufficient funds"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(coins Bitcoin) {
	w.balance += coins
}

func (w *Wallet) Withdraw(coins Bitcoin) error {
	if coins > w.balance {
		return errors.New(ErrInsufficientFunds)
	}
	w.balance -= coins
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

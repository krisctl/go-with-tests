package ptr

import "fmt"

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

func (w *Wallet) Withdraw(coins Bitcoin) {
	w.balance -= coins
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

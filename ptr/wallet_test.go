package ptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	assertBalance := func(want, got Bitcoin) {
		assert.Equal(t, want, got, "Wallet balance doesn't match the expected amount, want: %s, got: %s", want, got)
	}
	t.Run("Deposit", func(t *testing.T) {
		w := Wallet{balance: Bitcoin(90)}
		w.Deposit(Bitcoin(10))
		assertBalance(Bitcoin(100), w.Balance())
	})
	t.Run("Withdraw", func(t *testing.T) {
		w := Wallet{balance: Bitcoin(100)}
		w.Withdraw(Bitcoin(10))
		assertBalance(Bitcoin(90), w.Balance())
	})
}

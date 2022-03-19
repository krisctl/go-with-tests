package ptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		err := w.Withdraw(Bitcoin(10))
		assert.NoError(t, err, "Got a non-nil error in the response")
		assertBalance(Bitcoin(90), w.Balance())
	})
	t.Run("Overdraft", func(t *testing.T) {
		startingBalance := Bitcoin(100)
		w := Wallet{balance: startingBalance}
		drawCoins := Bitcoin(110)
		err := w.Withdraw(drawCoins)
		currbalance := w.Balance()
		assertBalance(startingBalance, currbalance)
		require.Error(t, err, "Withdrawing more than the available balance did not return error, drawing: %s, available: %s", drawCoins, currbalance)
		assert.Equal(t, err.Error(), ErrInsufficientFunds)
	})
}

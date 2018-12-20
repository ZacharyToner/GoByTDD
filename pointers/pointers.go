//Package pointers is for exploring pointers and errors using a 'Bitcoin' wallet
package pointers

import (
	"errors"
	"fmt"
)

//Bitcoin is an int behind the scenes
type Bitcoin int

//Bitcoin stringer
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//ErrInsufficientFunds is our custom error type for insufficient funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Wallet struct to represent a Bitcoin Wallet
type Wallet struct {
	balance Bitcoin
}

//Deposit adds provided number of Bitcoins to the Wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Withdraw removes the provided number of Bitcoins from the Wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

//Balance returns the current amount of Bitcoins in the Wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

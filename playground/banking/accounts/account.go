package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}

func (a *account) Deposit(amount int) {
	a.balance += amount
}

func (a account) Balance() int {
	return a.balance
}

func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a account) Owner() string {
	return a.owner
}

func (a account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}

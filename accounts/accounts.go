package accounts

import "errors"

// Account struct
type account struct {
	owner   string
	balance int
}

// NewAccount constructor
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}

// Deposit
func (a *account) Deposit(amount int) {
	a.balance += amount
}

// Balance
func (a account) Balance() int {
	return a.balance
}

var errNoMoney = errors.New("can't withdraw")

// WithDraw
func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner
func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner
func (a account) Owner() string {
	return a.owner
}

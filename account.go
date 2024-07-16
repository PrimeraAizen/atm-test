package ATMtesttask

import "errors"

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	balance float64
}

func NewAccount() *Account {
	return &Account{
		balance: 0,
	}
}

func (a *Account) Deposit(amount float64) error {
	a.balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if a.balance < amount {
		return errors.New("Not enough money on the account")
	}
	a.balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

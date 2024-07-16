package main

import (
	"fmt"
	"sync"
)

type Operation struct {
	AccountID int
	Amount    float64
	Type      string
	Response  chan error
}

var (
	Accounts         = make(map[int]*Account)
	AccountIDcounter = 1
)

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	ID      int
	Balance float64
	mu      sync.Mutex
}

func NewAccount() *Account {
	account := &Account{ID: AccountIDcounter}
	Accounts[AccountIDcounter] = account
	AccountIDcounter++
	return account
}

func (a *Account) Deposit(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amount > a.Balance {
		return fmt.Errorf("insufficient funds")
	}
	a.Balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.Balance
}

func HandleOperations(ops chan Operation) {
	for op := range ops {
		account, exists := Accounts[op.AccountID]
		if !exists {
			op.Response <- fmt.Errorf("account not found")
			continue
		}
		switch op.Type {
		case "deposit":
			err := account.Deposit(op.Amount)
			op.Response <- err
		case "withdraw":
			err := account.Withdraw(op.Amount)
			op.Response <- err
		case "balance":
			op.Response <- nil
			fmt.Printf("Account ID: %d, Balance: %.2f\n", op.AccountID, account.GetBalance())
		}
	}
}

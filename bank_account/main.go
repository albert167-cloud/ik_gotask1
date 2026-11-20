package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// Sentinel errors. Callers (and tests) can compare against them with errors.Is.
var (
	ErrNilAccount         = errors.New("account is nil")
	ErrEmptyAccountNumber = errors.New("account number must not be empty")
	ErrEmptyHolder        = errors.New("account holder must not be empty")
	ErrInvalidBalance     = errors.New("initial balance must be a non-negative number")
	ErrInvalidAmount      = errors.New("amount must be a positive number")
	ErrInsufficientFunds  = errors.New("insufficient funds")
	ErrSameAccount        = errors.New("sender and receiver cannot be the same account")
)

type BankAccount struct {
	AccountNumber string
	AccountHolder string
	Balance       float64
}

// isFinite reports whether x is a real number (not NaN and not +/-Inf).
func isFinite(x float64) bool {
	return !math.IsNaN(x) && !math.IsInf(x, 0)
}

// validateAmount checks that a deposit/withdrawal/transfer amount is a finite, positive number.
func validateAmount(amount float64) error {
	if !isFinite(amount) || amount <= 0 {
		return ErrInvalidAmount
	}
	return nil
}

// createBankAccount validates its inputs and returns a new account, or an error.
func createBankAccount(accountNumber string, accountHolder string, initialBalance float64) (*BankAccount, error) {
	if strings.TrimSpace(accountNumber) == "" {
		return nil, ErrEmptyAccountNumber
	}
	if strings.TrimSpace(accountHolder) == "" {
		return nil, ErrEmptyHolder
	}
	if !isFinite(initialBalance) || initialBalance < 0 {
		return nil, ErrInvalidBalance
	}
	return &BankAccount{
		AccountNumber: accountNumber,
		AccountHolder: accountHolder,
		Balance:       initialBalance,
	}, nil
}

// deposit adds amount to the account balance.
func deposit(account *BankAccount, amount float64) error {
	if account == nil {
		return ErrNilAccount
	}
	if err := validateAmount(amount); err != nil {
		return err
	}
	account.Balance += amount
	return nil
}

// withdraw subtracts amount from the account balance.
func withdraw(account *BankAccount, amount float64) error {
	if account == nil {
		return ErrNilAccount
	}
	if err := validateAmount(amount); err != nil {
		return err
	}
	if amount > account.Balance {
		return ErrInsufficientFunds
	}
	account.Balance -= amount
	return nil
}

// checkBalance prints the account details.
func checkBalance(account *BankAccount) error {
	if account == nil {
		return ErrNilAccount
	}
	fmt.Printf("Account Number: %s\n", account.AccountNumber)
	fmt.Printf("Account Holder: %s\n", account.AccountHolder)
	fmt.Printf("Current Balance: %.2f\n", account.Balance)
	fmt.Println("====================================================")
	return nil
}

// transfer moves amount from one account to another. Nothing changes if it returns an error.
func transfer(fromAccount *BankAccount, toAccount *BankAccount, amount float64) error {
	if fromAccount == nil || toAccount == nil {
		return ErrNilAccount
	}
	if fromAccount == toAccount {
		return ErrSameAccount
	}
	if err := validateAmount(amount); err != nil {
		return err
	}
	if amount > fromAccount.Balance {
		return ErrInsufficientFunds
	}
	fromAccount.Balance -= amount
	toAccount.Balance += amount
	return nil
}

// report prints the outcome of an operation.
func report(action string, err error) {
	if err != nil {
		fmt.Printf("%s failed: %v\n", action, err)
		fmt.Println("====================================================")
		return
	}
	fmt.Printf("%s succeeded\n", action)
}

func main() {
	account1, err := createBankAccount("123456", "Alice", 1000.0)
	if err != nil {
		fmt.Println("could not create account 1:", err)
		return
	}
	account2, err := createBankAccount("654321", "Bob", 500.0)
	if err != nil {
		fmt.Println("could not create account 2:", err)
		return
	}
	account3, err := createBankAccount("987654", "Charlie", 0.0)
	if err != nil {
		fmt.Println("could not create account 3:", err)
		return
	}

	report("deposit 200 to Alice", deposit(account1, 200.0))
	report("withdraw 100 from Bob", withdraw(account2, 100.0))
	report("withdraw 100 from Charlie", withdraw(account3, 100.0))
	checkBalance(account1)
	checkBalance(account2)
	checkBalance(account3)

	report("transfer 300 Alice -> Bob", transfer(account1, account2, 300.0))
	checkBalance(account1)
	checkBalance(account2)

	// Examples of errors being reported instead of silently ignored:
	report("withdraw 10000 from Bob", withdraw(account2, 10000.0))
	report("transfer Alice -> Alice", transfer(account1, account1, 10.0))

	// Tests live in main_test.go. Run them with:  go test -v
}
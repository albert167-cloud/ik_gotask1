package main 


import (
	"fmt"
	"testing"
)

type BankAccount struct {
	AccountNumber string
	AccountHolder string
	Balance       float64
}

func createBankAccount(accountNumber string, accountHolder string, initialBalance float64) *BankAccount {
	// Create a new bank account with the given details
	return &BankAccount{
		AccountNumber: accountNumber,
		AccountHolder: accountHolder,
		Balance: initialBalance,
	}
}

func deposit(account *BankAccount, amount float64) {
	// check if amount is positive
	if amount <= 0 {
		fmt.Println("Deposit amount must be positive.")
		return
	}	
	// Add the deposit amount to the account balance
	account.Balance += amount
}

func withdraw(account *BankAccount, amount float64) {
	// check if amount is positive
	if amount <= 0 {
		fmt.Println("Withdrawal amount must be positive.")
		return
	}
	// check if there are sufficient funds
	if amount > account.Balance {
		fmt.Println("Insufficient funds.")
		return
	}
	// Subtract the withdrawal amount from the account balance
	account.Balance -= amount
}	

func checkBalance(account *BankAccount) {
	// Print the current balance of the account
	fmt.Printf("Account Number: %s\n", account.AccountNumber)
	fmt.Printf("Account Holder: %s\n", account.AccountHolder)
	fmt.Printf("Current Balance: %.2f\n", account.Balance)
}	

func transfer(fromAccount *BankAccount, toAccount *BankAccount, amount float64) {
	// check if sender is not same as receiver
	if fromAccount == toAccount {
		fmt.Println("Sender and receiver cannot be the same.")
		return
	}
	// check if amount is positive and if there are sufficient funds
	if amount <= 0 {
		fmt.Println("Transfer amount must be positive.")
		return
	}	
	// check if there are sufficient funds in the sender's account
	if amount > fromAccount.Balance {
		fmt.Println("Insufficient funds for transfer.")
		return
	}	
	// Subtract the transfer amount from the sender's account and add it to the receiver's account
	fromAccount.Balance -= amount
	toAccount.Balance += amount
}

func TestBankAccount(t *testing.T) {
	// fill in the test cases for the bank account functions following comments below
	// Test case 1: valid account creation

	// Test case 2: empty account holder name

	// Test case 3: Successful deposit

	// Test case 4: Zero deposit

	// Test case 5: Negative deposit

	// Test case 6: Successful withdrawal

	// Test case 7: Withdrawal with insufficient funds

	// Test case 8: Zero withdrawal

	// Test case 9: Negative withdrawal

	// Test case 10: Account with existing balance

	// Test case 11: Successful transfer

	// Test case 12: Transfer with insufficient funds

	// Test case 13: Zero or negative transfer

	// Test case 14: Transfer to the same account
}

func main() {
	// // Create two bank accounts
	// account1 := createBankAccount("123456", "Alice", 1000.0)
	// account2 := createBankAccount("654321", "Bob", 500.0)

	// // Perform some operations on the accounts
	// deposit(account1, 200.0)
	// withdraw(account2, 100.0)
	// checkBalance(account1)
	// checkBalance(account2)

	// // Transfer funds from account1 to account2
	// transfer(account1, account2, 300.0)
	// checkBalance(account1)
	// checkBalance(account2)


	// Run the test cases for the bank account functions follow the testcase.txt file and use the testing package to run the tests
	testBankAccount := testing.T{}
	TestBankAccount(&testBankAccount)	
	
}
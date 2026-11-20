package main

// ============================================================================
// HOW EVERY TEST IN THIS FILE WORKS (the "3 A's")
//
//   1. ARRANGE - set things up (for example, make an account with 100 coins)
//   2. ACT     - do the one thing we want to test (for example, deposit 50)
//   3. ASSERT  - check the result is what we expected. If it is NOT,
//                we tell Go the test failed, using t.Errorf or t.Fatalf.
//
// If a test finishes without calling t.Errorf / t.Fatalf, it PASSES.
//
// Run all tests with:   go test -v
// ============================================================================

import (
	"errors" // lets us ask: "is this the exact error we expected?"
	"math"   // gives us strange numbers: NaN ("not a number") and Inf (infinity)
	"testing"
)

// ----------------------------------------------------------------------------
// HELPER: newAccount
// Many tests need a ready-made account. This helper builds one so we do not
// repeat the same lines everywhere. If building fails, the test stops at once
// (t.Fatalf), because there is no point continuing without an account.
// ----------------------------------------------------------------------------
func newAccount(t *testing.T, number, holder string, balance float64) *BankAccount {
	t.Helper() // if this fails, Go reports the line in the TEST, not in this helper

	account, err := createBankAccount(number, holder, balance)
	if err != nil {
		t.Fatalf("could not set up the test account: %v", err)
	}
	return account
}

// ============================================================================
// PART 1: CREATING AN ACCOUNT
// ============================================================================

// A normal account should be created with exactly the details we gave.
func TestCreateAccount_Valid(t *testing.T) {
	// ACT: create an account with good details
	account, err := createBankAccount("123456", "Alice", 100.0)

	// ASSERT: there must be no error
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	// ASSERT: every field must match what we passed in
	if account.AccountNumber != "123456" {
		t.Errorf("expected account number 123456, got %s", account.AccountNumber)
	}
	if account.AccountHolder != "Alice" {
		t.Errorf("expected holder Alice, got %s", account.AccountHolder)
	}
	if account.Balance != 100.0 {
		t.Errorf("expected balance 100, got %.2f", account.Balance)
	}
}

// An account can start with any positive balance, for example 250.
func TestCreateAccount_ExistingBalance(t *testing.T) {
	// ARRANGE + ACT: create an account that starts with 250
	account := newAccount(t, "123456", "Alice", 250.0)

	// ASSERT: the starting balance is kept
	if account.Balance != 250.0 {
		t.Errorf("expected balance 250, got %.2f", account.Balance)
	}
}

// Starting with zero coins is allowed (an empty account is still an account).
func TestCreateAccount_ZeroBalanceAllowed(t *testing.T) {
	account := newAccount(t, "123456", "Alice", 0)

	if account.Balance != 0 {
		t.Errorf("expected balance 0, got %.2f", account.Balance)
	}
}

// An account with no holder name must be refused.
func TestCreateAccount_EmptyHolder(t *testing.T) {
	// ACT: try to create an account with an empty name
	account, err := createBankAccount("123456", "", 0)

	// ASSERT: we expect the "empty holder" error...
	if !errors.Is(err, ErrEmptyHolder) {
		t.Errorf("expected ErrEmptyHolder, got: %v", err)
	}
	// ASSERT: ...and no account should be returned
	if account != nil {
		t.Errorf("expected no account, but got: %+v", account)
	}
}

// A name made only of spaces is also "empty" and must be refused.
func TestCreateAccount_HolderOnlySpaces(t *testing.T) {
	account, err := createBankAccount("123456", "   ", 0)

	if !errors.Is(err, ErrEmptyHolder) {
		t.Errorf("expected ErrEmptyHolder, got: %v", err)
	}
	if account != nil {
		t.Errorf("expected no account, but got: %+v", account)
	}
}

// An account with no account number must be refused.
func TestCreateAccount_EmptyAccountNumber(t *testing.T) {
	account, err := createBankAccount("", "Alice", 0)

	if !errors.Is(err, ErrEmptyAccountNumber) {
		t.Errorf("expected ErrEmptyAccountNumber, got: %v", err)
	}
	if account != nil {
		t.Errorf("expected no account, but got: %+v", account)
	}
}

// You cannot start an account with negative money.
func TestCreateAccount_NegativeBalance(t *testing.T) {
	account, err := createBankAccount("123456", "Alice", -1)

	if !errors.Is(err, ErrInvalidBalance) {
		t.Errorf("expected ErrInvalidBalance, got: %v", err)
	}
	if account != nil {
		t.Errorf("expected no account, but got: %+v", account)
	}
}

// NaN means "not a number". It is not a real amount of money, so refuse it.
func TestCreateAccount_NaNBalance(t *testing.T) {
	account, err := createBankAccount("123456", "Alice", math.NaN())

	if !errors.Is(err, ErrInvalidBalance) {
		t.Errorf("expected ErrInvalidBalance, got: %v", err)
	}
	if account != nil {
		t.Errorf("expected no account, but got: %+v", account)
	}
}

// Infinity is not a real amount of money either, so refuse it.
func TestCreateAccount_InfiniteBalance(t *testing.T) {
	account, err := createBankAccount("123456", "Alice", math.Inf(1))

	if !errors.Is(err, ErrInvalidBalance) {
		t.Errorf("expected ErrInvalidBalance, got: %v", err)
	}
	if account != nil {
		t.Errorf("expected no account, but got: %+v", account)
	}
}

// ============================================================================
// PART 2: DEPOSITING MONEY
// ============================================================================

// Depositing 50 into an account with 100 should give 150.
func TestDeposit_Success(t *testing.T) {
	// ARRANGE: Alice has 100
	account := newAccount(t, "123456", "Alice", 100.0)

	// ACT: deposit 50
	err := deposit(account, 50.0)

	// ASSERT: no error, and the balance went up to 150
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if account.Balance != 150.0 {
		t.Errorf("expected balance 150, got %.2f", account.Balance)
	}
}

// Depositing zero makes no sense, so it is refused and the balance is unchanged.
func TestDeposit_Zero(t *testing.T) {
	account := newAccount(t, "123456", "Alice", 100.0)

	err := deposit(account, 0)

	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("expected ErrInvalidAmount, got: %v", err)
	}
	if account.Balance != 100.0 {
		t.Errorf("balance should stay 100, got %.2f", account.Balance)
	}
}

// Depositing a negative amount would secretly be a withdrawal, so refuse it.
func TestDeposit_Negative(t *testing.T) {
	account := newAccount(t, "123456", "Alice", 100.0)

	err := deposit(account, -10.0)

	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("expected ErrInvalidAmount, got: %v", err)
	}
	if account.Balance != 100.0 {
		t.Errorf("balance should stay 100, got %.2f", account.Balance)
	}
}

// "Not a number" must never be accepted; it would ruin the balance.
func TestDeposit_NaN(t *testing.T) {
	account := newAccount(t, "123456", "Alice", 100.0)

	err := deposit(account, math.NaN())

	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("expected ErrInvalidAmount, got: %v", err)
	}
	if account.Balance != 100.0 {
		t.Errorf("balance should stay 100, got %.2f", account.Balance)
	}
}

// Infinity must never be accepted either.
func TestDeposit_Infinite(t *testing.T) {
	account := newAccount(t, "123456", "Alice", 100.0)

	err := deposit(account, math.Inf(1))

	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("expected ErrInvalidAmount, got: %v", err)
	}
	if account.Balance != 100.0 {
		t.Errorf("balance should stay 100, got %.2f", account.Balance)
	}
}

// Depositing into an account that does not exist (nil) must not crash.
// It should just return an error.
func TestDeposit_NilAccount(t *testing.T) {
	err := deposit(nil, 10)

	if !errors.Is(err, ErrNilAccount) {
		t.Errorf("expected ErrNilAccount, got: %v", err)
	}
}

// ============================================================================
// PART 3: WITHDRAWING MONEY
// ============================================================================

// Withdrawing 40 from 100 should leave 60.
func TestWithdraw_Success(t *testing.T) {
	account := newAccount(t, "123456", "Alice", 100.0)

	err := withdraw(account, 40.0)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if account.Balance != 60.0 {
		t.Errorf("expected balance 60, got %.2f", account.Balance)
	}
}

// Taking out EVERYTHING is allowed. The balance becomes exactly 0.
// (This checks the edge: "amount > balance" is refused, but "amount == balance" is fine.)
func TestWithdraw_ExactBalance(t *testing.T) {
	account := newAccount(t, "123456", "Alice", 100.0)

	err := withdraw(account, 100.0)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if account.Balance != 0 {
		t.Errorf("expected balance 0, got %.2f", account.Balance)
	}
}

// Trying to take out more than you have must fail, and the balance stays safe.
func TestWithdraw_InsufficientFunds(t *testing.T) {
	// ARRANGE: Alice has only 50
	account := newAccount(t, "123456", "Alice", 50.0)

	// ACT: she tries to take out 60
	err := withdraw(account, 60.0)

	// ASSERT: the bank says "insufficient funds" and her 50 is untouched
	if !errors.Is(err, ErrInsufficientFunds) {
		t.Errorf("expected ErrInsufficientFunds, got: %v", err)
	}
	if account.Balance != 50.0 {
		t.Errorf("balance should stay 50, got %.2f", account.Balance)
	}
}

// Withdrawing zero makes no sense, so it is refused.
func TestWithdraw_Zero(t *testing.T) {
	account := newAccount(t, "123456", "Alice", 100.0)

	err := withdraw(account, 0)

	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("expected ErrInvalidAmount, got: %v", err)
	}
	if account.Balance != 100.0 {
		t.Errorf("balance should stay 100, got %.2f", account.Balance)
	}
}

// Withdrawing a negative amount would secretly be a deposit, so refuse it.
func TestWithdraw_Negative(t *testing.T) {
	account := newAccount(t, "123456", "Alice", 100.0)

	err := withdraw(account, -10.0)

	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("expected ErrInvalidAmount, got: %v", err)
	}
	if account.Balance != 100.0 {
		t.Errorf("balance should stay 100, got %.2f", account.Balance)
	}
}

// "Not a number" must be refused.
func TestWithdraw_NaN(t *testing.T) {
	account := newAccount(t, "123456", "Alice", 100.0)

	err := withdraw(account, math.NaN())

	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("expected ErrInvalidAmount, got: %v", err)
	}
	if account.Balance != 100.0 {
		t.Errorf("balance should stay 100, got %.2f", account.Balance)
	}
}

// Infinity must be refused.
func TestWithdraw_Infinite(t *testing.T) {
	account := newAccount(t, "123456", "Alice", 100.0)

	err := withdraw(account, math.Inf(1))

	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("expected ErrInvalidAmount, got: %v", err)
	}
	if account.Balance != 100.0 {
		t.Errorf("balance should stay 100, got %.2f", account.Balance)
	}
}

// Withdrawing from an account that does not exist (nil) must not crash.
func TestWithdraw_NilAccount(t *testing.T) {
	err := withdraw(nil, 10)

	if !errors.Is(err, ErrNilAccount) {
		t.Errorf("expected ErrNilAccount, got: %v", err)
	}
}

// ============================================================================
// PART 4: TRANSFERRING MONEY BETWEEN TWO ACCOUNTS
// ============================================================================

// Alice (100) sends 40 to Bob (25). Alice should have 60 and Bob should have 65.
func TestTransfer_Success(t *testing.T) {
	// ARRANGE: two accounts
	alice := newAccount(t, "123456", "Alice", 100.0)
	bob := newAccount(t, "654321", "Bob", 25.0)

	// ACT: Alice sends 40 to Bob
	err := transfer(alice, bob, 40.0)

	// ASSERT: no error; Alice lost 40; Bob gained 40
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if alice.Balance != 60.0 {
		t.Errorf("expected Alice to have 60, got %.2f", alice.Balance)
	}
	if bob.Balance != 65.0 {
		t.Errorf("expected Bob to have 65, got %.2f", bob.Balance)
	}
}

// Alice has 20 but tries to send 30. The transfer fails and NOBODY's money changes.
func TestTransfer_InsufficientFunds(t *testing.T) {
	alice := newAccount(t, "123456", "Alice", 20.0)
	bob := newAccount(t, "654321", "Bob", 25.0)

	err := transfer(alice, bob, 30.0)

	if !errors.Is(err, ErrInsufficientFunds) {
		t.Errorf("expected ErrInsufficientFunds, got: %v", err)
	}
	if alice.Balance != 20.0 {
		t.Errorf("Alice should still have 20, got %.2f", alice.Balance)
	}
	if bob.Balance != 25.0 {
		t.Errorf("Bob should still have 25, got %.2f", bob.Balance)
	}
}

// Transferring zero makes no sense; nobody's money changes.
func TestTransfer_Zero(t *testing.T) {
	alice := newAccount(t, "123456", "Alice", 100.0)
	bob := newAccount(t, "654321", "Bob", 25.0)

	err := transfer(alice, bob, 0)

	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("expected ErrInvalidAmount, got: %v", err)
	}
	if alice.Balance != 100.0 || bob.Balance != 25.0 {
		t.Errorf("balances should not change: Alice %.2f, Bob %.2f", alice.Balance, bob.Balance)
	}
}

// Transferring a negative amount would secretly move money the wrong way. Refuse it.
func TestTransfer_Negative(t *testing.T) {
	alice := newAccount(t, "123456", "Alice", 100.0)
	bob := newAccount(t, "654321", "Bob", 25.0)

	err := transfer(alice, bob, -10)

	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("expected ErrInvalidAmount, got: %v", err)
	}
	if alice.Balance != 100.0 || bob.Balance != 25.0 {
		t.Errorf("balances should not change: Alice %.2f, Bob %.2f", alice.Balance, bob.Balance)
	}
}

// "Not a number" must be refused; nobody's money changes.
func TestTransfer_NaN(t *testing.T) {
	alice := newAccount(t, "123456", "Alice", 100.0)
	bob := newAccount(t, "654321", "Bob", 25.0)

	err := transfer(alice, bob, math.NaN())

	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("expected ErrInvalidAmount, got: %v", err)
	}
	if alice.Balance != 100.0 || bob.Balance != 25.0 {
		t.Errorf("balances should not change: Alice %.2f, Bob %.2f", alice.Balance, bob.Balance)
	}
}

// Infinity must be refused; nobody's money changes.
func TestTransfer_Infinite(t *testing.T) {
	alice := newAccount(t, "123456", "Alice", 100.0)
	bob := newAccount(t, "654321", "Bob", 25.0)

	err := transfer(alice, bob, math.Inf(1))

	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("expected ErrInvalidAmount, got: %v", err)
	}
	if alice.Balance != 100.0 || bob.Balance != 25.0 {
		t.Errorf("balances should not change: Alice %.2f, Bob %.2f", alice.Balance, bob.Balance)
	}
}

// Sending money to yourself is pointless, so it is refused.
func TestTransfer_SameAccount(t *testing.T) {
	alice := newAccount(t, "123456", "Alice", 100.0)

	// ACT: Alice sends money to Alice (same account used twice)
	err := transfer(alice, alice, 25.0)

	if !errors.Is(err, ErrSameAccount) {
		t.Errorf("expected ErrSameAccount, got: %v", err)
	}
	if alice.Balance != 100.0 {
		t.Errorf("balance should stay 100, got %.2f", alice.Balance)
	}
}

// If the SENDER does not exist (nil), we must not crash.
func TestTransfer_NilSender(t *testing.T) {
	bob := newAccount(t, "654321", "Bob", 25.0)

	err := transfer(nil, bob, 10)

	if !errors.Is(err, ErrNilAccount) {
		t.Errorf("expected ErrNilAccount, got: %v", err)
	}
	if bob.Balance != 25.0 {
		t.Errorf("Bob should still have 25, got %.2f", bob.Balance)
	}
}

// If the RECEIVER does not exist (nil), we must not crash.
func TestTransfer_NilReceiver(t *testing.T) {
	alice := newAccount(t, "123456", "Alice", 100.0)

	err := transfer(alice, nil, 10)

	if !errors.Is(err, ErrNilAccount) {
		t.Errorf("expected ErrNilAccount, got: %v", err)
	}
	if alice.Balance != 100.0 {
		t.Errorf("Alice should still have 100, got %.2f", alice.Balance)
	}
}
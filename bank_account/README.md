# Bank Account (Go)

A small bank account library with validated operations, tested by **one table-driven test function** that also prints a summary report.

## Files

| File | Purpose |
|---|---|
| `main.go` | `BankAccount`, `createBankAccount`, `deposit`, `withdraw`, `checkBalance`, `transfer`, plus a demo in `main()` |
| `main_test.go` | `TestBank`: one function, 32 cases, per-group report |
| `go.mod` | Module definition (Go 1.21+) |

## Run

```bash
go run .        # demo
go test -v      # all cases + report
go test -v -run TestBank/deposit    # one group only
```

Expected report (all passing):

```
create      9/9  passed
deposit     6/6  passed
withdraw    8/8  passed
transfer    9/9  passed
TOTAL      32/32 passed
```

## Design (software level)

**API contract.** Every operation returns an `error`. Failures are *sentinel errors* (`ErrInvalidAmount`, `ErrInsufficientFunds`, `ErrSameAccount`, `ErrNilAccount`, `ErrEmptyHolder`, `ErrEmptyAccountNumber`, `ErrInvalidBalance`), so callers branch with `errors.Is` instead of parsing strings.

**Input validation at the boundary.** Amounts must be finite and `> 0` (`NaN` would slip past a plain `<= 0` check). Accounts may not be `nil`. Transfers need two different accounts.

**Failure atomicity.** All checks run before any state changes, so a failed call leaves every balance untouched.

## Testing approach

- **Table-driven test.** Each case is one row: `name`, `run` (the operation), `wantErr`, `want` (expected balances). Adding a case is one line; the runner code never changes.
- **Subtests (`t.Run`).** Each row runs independently with a fresh fixture (Alice 100, Bob 25), so cases cannot affect each other, and a failure names the exact row.
- **Oracle.** Expected error plus expected balances, compared after each run.
- **Invariant.** For every error case, `want` equals the starting balances, so one rule enforces "a failed operation changes nothing."
- **Closures as operation builders.** `dep(x)`, `wd(x)`, `xfer(x)` and `create(...)` turn an amount into an operation, keeping rows to a single line.
- **Coverage by equivalence class.** Per operation: valid input, boundary (zero, exact balance), and each invalid class (negative, `NaN`, infinity, `nil`, insufficient funds).
- **Generalized report.** Results are grouped by the name prefix (`create`, `deposit`, ...) and printed as `passed/total` per group, so the report stays correct as cases are added.

## Add a test

Append one row to `cases` in `main_test.go`:

```go
{"withdraw_tiny", wd(0.01), nil, [2]float64{99.99, 25}},
```

## Limitations

- No locking, so accounts are not safe for concurrent use.
- State is in memory only, with no history.
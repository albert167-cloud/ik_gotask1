# 13. Roman to Integer

Convert a Roman numeral to an integer. Solved in Go in a single pass.

## Problem

Roman numerals use seven symbols:

| Symbol | Value |
|--------|-------|
| I      | 1     |
| V      | 5     |
| X      | 10    |
| L      | 50    |
| C      | 100   |
| D      | 500   |
| M      | 1000  |

Numerals are usually written largest to smallest, left to right, and the values are added. For example, `XII` is 10 + 1 + 1 = 12, and `XXVII` is 10 + 10 + 5 + 1 + 1 = 27.

There is one exception. Four is written `IV`, not `IIII`, because a smaller symbol placed **before** a larger one is subtracted. There are six such pairs:

- `I` before `V` or `X` makes 4 and 9
- `X` before `L` or `C` makes 40 and 90
- `C` before `D` or `M` makes 400 and 900

**Task:** given a valid Roman numeral `s`, convert it to an integer.

### Examples

| Input       | Output | Explanation        |
|-------------|--------|--------------------|
| `"III"`     | 3      | 1 + 1 + 1          |
| `"LVIII"`   | 58     | 50 + 5 + 1 + 1 + 1 |
| `"MCMXCIV"` | 1994   | 1000 + 900 + 90 + 4 |

### Constraints

- `1 <= s.length <= 15`
- `s` contains only `I, V, X, L, C, D, M`
- `s` is guaranteed to be a valid Roman numeral in the range `[1, 3999]`

## Solution

### Idea

You do not need to memorize the six subtraction pairs. One rule covers all of them:

> Look at each symbol and the one after it. If it is **smaller** than the next one, subtract it. Otherwise, add it.

A smaller symbol before a larger one is exactly what a subtraction pair is. In every other case the symbol is simply added.

### Steps

1. Build a lookup table from symbol to value.
2. Start a running total at 0.
3. Go through the string from left to right.
4. Compare each symbol's value with the next symbol's value:
   - smaller than the next: subtract it from the total
   - otherwise: add it to the total
5. The last symbol has no next symbol, so it is always added.
6. Return the total.

### Code

```go
func romanToInt(s string) int {
	val := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}

	total := 0
	for i := 0; i < len(s); i++ {
		cur := val[s[i]]

		// Smaller than the next symbol? Subtract. Otherwise add.
		if i+1 < len(s) && cur < val[s[i+1]] {
			total -= cur
		} else {
			total += cur
		}
	}
	return total
}
```

The check `i+1 < len(s)` makes sure a next symbol exists before it is read. Go stops evaluating `&&` as soon as the left side is false, so the last symbol never reads past the end of the string.

### Walkthrough: `MCMXCIV`

| Symbol | Value | Next     | Smaller than next? | Action   | Total    |
|--------|-------|----------|--------------------|----------|----------|
| M      | 1000  | C (100)  | no                 | add      | 1000     |
| C      | 100   | M (1000) | yes                | subtract | 900      |
| M      | 1000  | X (10)   | no                 | add      | 1900     |
| X      | 10    | C (100)  | yes                | subtract | 1890     |
| C      | 100   | I (1)    | no                 | add      | 1990     |
| I      | 1     | V (5)    | yes                | subtract | 1989     |
| V      | 5     | none     | no                 | add      | **1994** |

## Complexity

- **Time: O(n).** One pass over the string, and n is at most 15.
- **Space: O(1).** The lookup table always has seven entries.

## Run it

The full program, with tests, is in `roman.go`:

```bash
go run roman.go
```

It prints `PASS` or `FAIL` for each test case.

## Summary

> Add each symbol's value, but if a symbol is smaller than the one after it, subtract it instead.

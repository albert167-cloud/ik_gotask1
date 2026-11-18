// LeetCode 13. Roman to Integer
//
// Approach: scan left to right and add each symbol's value, except when a
// symbol is smaller than the one right after it. In that case it is part of a
// subtraction pair (IV, IX, XL, XC, CD, CM), so subtract it instead.
//
// Time:  O(n)
// Space: O(1)
//
// Run:  go run roman.go
package main

import "fmt"

func romanToInt(s string) int {
	val := map[byte]int{
		'I': 1, 
		'V': 5, 
		'X': 10, 
		'L': 50,
		'C': 100, 
		'D': 500, 
		'M': 1000,
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

func main() {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"test 1 (example 1)", "LX", 60},
		{"test 2 (example 2)", "LVIII", 58},
		{"test 3 (example 3)", "MCMXCIV", 1994},
		{"four", "IV", 4},
		{"nine", "IX", 9},
		{"forty", "XL", 40},
		{"ninety", "XC", 90},
		{"four hundred", "CD", 400},
		{"nine hundred", "CM", 900},
		{"largest value", "MMMCMXCIX", 3999},
		{"single symbol", "V", 5},
	}

	allPassed := true

	for _, t := range tests {
		got := romanToInt(t.input)
		pass := got == t.expected

		status := "PASS"
		if !pass {
			status = "FAIL"
			allPassed = false
		}

		fmt.Printf("%s: %s\n", t.name, status)
		fmt.Printf("  Input    s = %q\n", t.input)
		fmt.Printf("  Output   %d\n", got)
		fmt.Printf("  Expected %d\n\n", t.expected)
	}

	if allPassed {
		fmt.Println("All tests passed.")
	} else {
		fmt.Println("Some tests failed.")
	}
}

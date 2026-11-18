// LeetCode 37. Sudoku Solver
//
// Approach: backtracking. Pick the next empty cell, try digits 1-9, keep a
// digit only if its row, column, and 3x3 box don't already use it, then move
// on. If we get stuck, undo the last guess and try the next digit.
// Bitmasks make the "is this digit already used?" check O(1).
//
// Run:  go run sudoku.go
package main

import "fmt"

// solveSudoku fills the empty cells ('.') of board in place.
func solveSudoku(board [][]byte) {
	// Bit d (0..8) is set when digit d+1 is already used.
	var rows, cols, boxes [9]int

	type cell struct{ r, c int }
	var empties []cell

	// Record the starting digits and collect the empty cells.
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				empties = append(empties, cell{r, c})
				continue
			}
			bit := 1 << int(board[r][c]-'1')
			rows[r] |= bit
			cols[c] |= bit
			boxes[(r/3)*3+c/3] |= bit
		}
	}

	var solve func(idx int) bool
	solve = func(idx int) bool {
		if idx == len(empties) {
			return true // every empty cell is filled
		}

		r, c := empties[idx].r, empties[idx].c
		b := (r/3)*3 + c/3

		for d := 0; d < 9; d++ {
			bit := 1 << d

			// Digit already used in this row, column, or box? Skip it.
			if (rows[r]|cols[c]|boxes[b])&bit != 0 {
				continue
			}

			// Place the digit.
			rows[r] |= bit
			cols[c] |= bit
			boxes[b] |= bit
			board[r][c] = byte('1' + d)

			if solve(idx + 1) {
				return true
			}

			// Dead end: undo the guess (backtrack).
			rows[r] &^= bit
			cols[c] &^= bit
			boxes[b] &^= bit
			board[r][c] = '.'
		}

		return false // no digit fits here, so the previous guess was wrong
	}

	solve(0)
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	expected := []string{
		"534678912",
		"672195348",
		"198342567",
		"859761423",
		"426853791",
		"713924856",
		"961537284",
		"287419635",
		"345286179",
	}

	solveSudoku(board)

	pass := true
	for r := 0; r < 9; r++ {
		fmt.Println(string(board[r]))
		if string(board[r]) != expected[r] {
			pass = false
		}
	}

	if pass {
		fmt.Println("\nPASS: matches the expected solution")
	} else {
		fmt.Println("\nFAIL: does not match the expected solution")
	}
}

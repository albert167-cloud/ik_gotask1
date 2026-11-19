// Validate Binary Search Tree (LeetCode 98)
//
// Run with:  go run main.go
package main

import (
	"fmt"
	"math"
)

// TreeNode is the binary tree node definition.
// (LeetCode provides this for you; remove it if you paste only the solution.)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isValidBST reports whether the tree rooted at root is a valid BST.
//
// Idea: every node must lie strictly inside an allowed range (low, high).
// The range is narrowed as we walk down:
//   - going left  -> the current value becomes the new upper bound
//   - going right -> the current value becomes the new lower bound
//
// This guarantees each node respects ALL of its ancestors, not just its parent.
func isValidBST(root *TreeNode) bool {
	var check func(n *TreeNode, low, high int) bool

	check = func(n *TreeNode, low, high int) bool {
		if n == nil {
			return true // an empty spot is always valid
		}
		if n.Val <= low || n.Val >= high {
			return false // outside the allowed range (equal values are invalid too)
		}
		return check(n.Left, low, n.Val) && check(n.Right, n.Val, high)
	}

	// Constraints say values fit in 32 bits, so the int extremes are safe sentinels.
	return check(root, math.MinInt, math.MaxInt)
}

// node is a small helper to build trees more concisely.
func node(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

func main() {
	// Example 1: [2,1,3] -> true
	t1 := node(2, node(1, nil, nil), node(3, nil, nil))
	fmt.Println("Example 1: [2,1,3] -> ", isValidBST(t1)) // true

	// Example 2: [5,1,4,null,null,3,6] -> false
	t2 := node(5,
		node(1, nil, nil),
		node(4, node(3, nil, nil), node(6, nil, nil)))
	fmt.Println("Example 2: [5,1,4,null,null,3,6] -> ", isValidBST(t2)) // false
	
	// Extra: empty tree -> true
	fmt.Println("Empty tree ->", isValidBST(nil)) // true
}
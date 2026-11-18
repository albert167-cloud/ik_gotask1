// LeetCode 297. Serialize and Deserialize Binary Tree
//
// Approach: preorder traversal (node, left, right) with "#" marking empty
// children. Reading the string back in the same order rebuilds the same tree.
//
// Time:  O(n)
// Space: O(n)
//
// Run:  go run main.go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------------------------
// Tree definition (LeetCode already provides this; it is included here so the
// file runs on its own. Remove it if you paste only the Codec into LeetCode.)
// ---------------------------------------------------------------------------

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ---------------------------------------------------------------------------
// Solution
// ---------------------------------------------------------------------------

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// serialize encodes a tree to a single string.
// Walk the tree as "node, left, right" and write "#" for an empty child.
func (c *Codec) serialize(root *TreeNode) string {
	var parts []string

	var walk func(n *TreeNode)
	walk = func(n *TreeNode) {
		if n == nil {
			parts = append(parts, "#") // nobody here
			return
		}
		parts = append(parts, strconv.Itoa(n.Val)) // the node
		walk(n.Left)                               // its left side
		walk(n.Right)                              // its right side
	}

	walk(root)
	return strings.Join(parts, ",")
}

// deserialize decodes the string back into the original tree.
// Read items in the same order they were written.
func (c *Codec) deserialize(data string) *TreeNode {
	parts := strings.Split(data, ",")
	i := 0 // the "finger" shared by every build() call

	var build func() *TreeNode
	build = func() *TreeNode {
		if i >= len(parts) {
			return nil
		}
		item := parts[i]
		i++

		if item == "#" {
			return nil
		}

		val, _ := strconv.Atoi(item)
		node := &TreeNode{Val: val}
		node.Left = build()
		node.Right = build()
		return node
	}

	return build()
}

// ---------------------------------------------------------------------------
// Helpers for testing: convert to and from LeetCode's display format,
// e.g. [1,2,3,null,null,4,5]
// ---------------------------------------------------------------------------

// parseLeetCode builds a tree from LeetCode's level-order format.
func parseLeetCode(s string) *TreeNode {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	if strings.TrimSpace(s) == "" {
		return nil
	}

	vals := strings.Split(s, ",")
	for i := range vals {
		vals[i] = strings.TrimSpace(vals[i])
	}
	if vals[0] == "null" {
		return nil
	}

	v, _ := strconv.Atoi(vals[0])
	root := &TreeNode{Val: v}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		if i < len(vals) && vals[i] != "null" {
			lv, _ := strconv.Atoi(vals[i])
			node.Left = &TreeNode{Val: lv}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(vals) && vals[i] != "null" {
			rv, _ := strconv.Atoi(vals[i])
			node.Right = &TreeNode{Val: rv}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// toLeetCode prints a tree in LeetCode's level-order format.
func toLeetCode(root *TreeNode) string {
	if root == nil {
		return "[]"
	}

	var out []string
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if n == nil {
			out = append(out, "null")
			continue
		}
		out = append(out, strconv.Itoa(n.Val))
		queue = append(queue, n.Left, n.Right)
	}

	// Trailing nulls are not part of LeetCode's format.
	for len(out) > 0 && out[len(out)-1] == "null" {
		out = out[:len(out)-1]
	}

	return "[" + strings.Join(out, ",") + "]"
}

// ---------------------------------------------------------------------------
// Tests
// ---------------------------------------------------------------------------

func main() {
	tests := []struct {
		name  string
		input string
	}{
		{"test 1 (example 1)", "[1,2,3,null,null,4,5]"},
		{"test 2 (example 2, empty tree)", "[]"},
		{"single node", "[7]"},
		{"negative values", "[-1,-2,3,null,-4]"},
		{"left-skewed", "[1,2,null,3,null,4]"},
		{"right-skewed", "[1,null,2,null,3,null,4]"},
	}

	codec := Constructor()
	allPassed := true

	for _, t := range tests {
		root := parseLeetCode(t.input)

		encoded := codec.serialize(root)
		decoded := codec.deserialize(encoded)

		output := toLeetCode(decoded)
		expected := t.input
		pass := output == expected

		if !pass {
			allPassed = false
		}

		status := "PASS"
		if !pass {
			status = "FAIL"
		}

		fmt.Printf("%s: %s\n", t.name, status)
		fmt.Printf("  Input      root = %s\n", t.input)
		fmt.Printf("  Serialized %s\n", encoded)
		fmt.Printf("  Output     %s\n", output)
		fmt.Printf("  Expected   %s\n\n", expected)
	}

	if allPassed {
		fmt.Println("All tests passed.")
	} else {
		fmt.Println("Some tests failed.")
	}
}

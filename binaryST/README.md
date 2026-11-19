# Validate Binary Search Tree (LeetCode 98) in Go

Determine whether a binary tree is a valid binary search tree (BST).

A valid BST means:
- Everything in a node's **left** subtree is **strictly less** than the node.
- Everything in a node's **right** subtree is **strictly greater** than the node.
- Both subtrees are themselves valid BSTs.

## Files

| File      | Purpose                                         |
|-----------|-------------------------------------------------|
| `main.go` | Solution (`isValidBST`) plus runnable test cases |
| `go.mod`  | Go module definition                            |
| `README.md` | This file                                     |

## Run it

```bash
go run main.go
```

Expected output:

```
Example 1: true
Example 2: false
Duplicate: false
Single node: true
Min int32: true
Deep violation: false
Empty tree: true
```

## Approach: range tracking

Each node must fall strictly inside an allowed range `(low, high)`.

1. Start at the root with range `(-infinity, +infinity)`.
2. At each node, if `Val <= low` or `Val >= high`, return `false`.
3. Recurse left with range `(low, node.Val)`.
4. Recurse right with range `(node.Val, high)`.
5. A `nil` node is valid.

This checks each node against **all ancestors**, which a simple parent/child
comparison misses (see Example 2: `3` is fine under `4`, but illegal under `5`'s right subtree).

## Complexity

- **Time:** O(n), each node is visited once.
- **Space:** O(h) recursion stack, where h is the tree height (O(n) worst case for a skewed tree).

## Bugs fixed from the original draft

1. `func(nTreeNode, ...)` was a typo; it must be `func(n *TreeNode, ...)`.
2. `isValidBST(root TreeNode)` took a value; it should take `*TreeNode` (LeetCode's signature, and what `main` passes).
3. `math.MinInt64` / `math.MaxInt64` replaced by `math.MinInt` / `math.MaxInt` so the code compiles on any platform where `int` is 32-bit.
4. `main` was unfinished (unclosed braces, empty comment); completed with test cases.
5. Removed a stale comment mentioning a "Codec" that doesn't exist in this file.

## Using it on LeetCode

Paste only the `isValidBST` function. LeetCode already defines `TreeNode`
and imports; add `import "math"` if needed.
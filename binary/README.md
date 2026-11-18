297. Serialize and Deserialize Binary Tree
Turn a binary tree into a string, then rebuild the identical tree from that string. Solved in Go with a preorder traversal.
Problem
Serialization is converting a data structure into a sequence of bits (here, a string) so it can be stored in a file or memory buffer, or sent across a network, and rebuilt later.
Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your algorithm works. You only need to guarantee that a tree can be serialized to a string, and that the string can be deserialized back to the original tree.
LeetCode displays trees in level order, like `[1,2,3,null,null,4,5]`. You do not have to use that format for your own string.
Examples
Input	Output
`[1,2,3,null,null,4,5]`	`[1,2,3,null,null,4,5]`
`[]`	`[]`
Constraints
The number of nodes is in the range `[0, 10^4]`
`-1000 <= Node.val <= 1000`
Solution
Idea
A tree is two-dimensional and a string is one line, so we need a fixed way to walk the tree and write each node down.
Pick one walk and keep to it. We use preorder: visit the node, then its whole left side, then its whole right side.
Values alone are not enough. The list `1,2,3` could be many different trees, because you cannot tell whether 2 is the left or the right child of 1.
Mark every empty spot with `#`. Now the string records the exact shape. A node is followed by its entire left subtree and then its entire right subtree, so only one tree fits the string.
Rebuild by reading in the same order. Read an item. If it is `#`, there is no node. Otherwise make a node, then build its left side, then its right side.
Serialize
Walk the tree as "node, left, right" and write `#` for an empty child. For this tree:
```
      1
     / \
    2   3
       / \
      4   5
```
Step	Where we are	We write
1	node 1	`1`
2	go left, node 2	`2`
3	2's left is empty	`#`
4	2's right is empty	`#`
5	back to 1, go right, node 3	`3`
6	go left, node 4	`4`
7	4's left and right are empty	`#`, `#`
8	back to 3, go right, node 5	`5`
9	5's left and right are empty	`#`, `#`
Result: `1,2,#,#,3,4,#,#,5,#,#`
Deserialize
Split the string on commas and keep a "finger" (an index) on the next item. Each call to `build()`:
Reads the item under the finger and moves the finger forward.
If the item is `#`, returns `nil`.
Otherwise creates a node, sets its left to `build()`, sets its right to `build()`, and returns it.
Because every call shares the same finger, each call picks up exactly where the previous one stopped. The `#` markers tell each call when its branch has ended.
Code
```go
type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// serialize encodes a tree to a single string.
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
```
LeetCode supplies `TreeNode`, so do not redeclare it when submitting.
Edge cases
Empty tree: `serialize` returns `"#"` and `deserialize` returns `nil`, with no special code.
Negative values: `-5` is just text. Only `#` means "nothing", so there is no confusion.
Skewed tree: up to 10^4 nodes in a line means recursion 10^4 levels deep. Go grows its stack as needed, so this is fine.
Commas: the separator keeps `12` from being confused with `1,2`.
Complexity
Time: O(n). Each node is visited once to serialize and once to deserialize.
Space: O(n). The string holds n values plus n + 1 `#` markers, and the recursion can go as deep as the tree is tall.
Run it
The full program, with LeetCode-format tests, is in `main.go`:
```bash
go run .
```
It prints `PASS` or `FAIL` for each test case, along with the serialized string.
Summary
> Walk the tree as node, left, right and write `#` for empty spots. To rebuild, read the list in the same order: a number makes a node with a left and right to build next, and `#` means stop.
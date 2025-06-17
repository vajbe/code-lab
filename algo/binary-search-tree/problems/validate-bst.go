package problems

/*
LeetCode Problem: 98. Validate Binary Search Tree
https://leetcode.com/problems/validate-binary-search-tree/

Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:
- The left subtree of a node contains only nodes with keys less than the node's key.
- The right subtree of a node contains only nodes with keys greater than the node's key.
- Both the left and right subtrees must also be binary search trees.

Example 1:
    2
   / \
  1   3
Input: root = [2,1,3]
Output: true

Example 2:
    5
   / \
  1   4
     / \
    3   6
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.

Constraints:
- The number of nodes in the tree is in the range [1, 10^4].
- -2^31 <= Node.val <= 2^31 - 1
*/

func validate(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}

	if min != nil && root.Val <= *min {
		return false
	}
	if max != nil && root.Val >= *max {
		return false
	}

	return validate(root.Left, min, &root.Val) && validate(root.Right, &root.Val, max)

}

func isValidBST(root *TreeNode) bool {
	return validate(root, nil, nil)
}

func ValidBST() {
	// Example 1: Valid BST
	root1 := &TreeNode{Val: 2}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 3}
	println("Example 1 (should be true):", isValidBST(root1))

	// Example 2: Invalid BST
	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 4}
	root2.Right.Left = &TreeNode{Val: 3}
	root2.Right.Right = &TreeNode{Val: 6}
	println("Example 2 (should be false):", isValidBST(root2))

	// Example 2: Invalid BST
	root3 := &TreeNode{Val: 0}
	root3.Left = &TreeNode{Val: -1}
	println("Example 3 (should be true):", isValidBST(root3))
}

// LeetCode 530. Minimum Absolute Difference in BST
// Given the root of a Binary Search Tree (BST), return the minimum absolute difference between the values of any two different nodes in the tree.
//
// Example 1:
//     4
//    / \
//   2   6
//  / \
// 1   3
// Output: 1
//
// Example 2:
//   1
//    \
//     3
//    /
//   2
// Output: 1
//
// Constraints:
// - The number of nodes in the tree is in the range [2, 10^4].
// - 0 <= Node.val <= 10^5
// - The input is guaranteed to be a valid BST.
//
// https://leetcode.com/problems/minimum-absolute-difference-in-bst/

package problems

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {

	var prev *int
	minDiff := math.MaxInt
	var inorder func(node *TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)

		if prev != nil {
			diff := node.Val - *prev
			if diff < minDiff {
				minDiff = diff
			}
		}

		val := node.Val
		prev = &val
		inorder(node.Right)
	}
	inorder(root)
	return minDiff
}

func MinimumDiff() {
	root1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 6},
	}
	fmt.Print("Min diff is ", getMinimumDifference(root1))
}

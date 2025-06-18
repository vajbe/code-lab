// Problem: Kth Smallest Element in a BST
// Source: https://leetcode.com/problems/kth-smallest-element-in-a-bst/?envType=study-plan-v2&envId=top-interview-150
//
// Description:
// Given the root of a binary search tree and an integer k, return the kth smallest
// value (1-indexed) of all the values of the nodes in the tree.
// The solution uses an iterative inorder traversal (with a stack) to collect node
// values in sorted order.
//
// Example BST:
//      3
//     / \
//    1   4
//     \
//      2
//
// For k = 2, the output is 2.

package problems

func kthSmallest(root *TreeNode, k int) int {
	head := root
	list := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for head != nil || len(stack) > 0 {
		for head != nil {
			stack = append(stack, head)
			head = head.Left
		}
		le := len(stack) - 1
		head = stack[le]
		stack = stack[:le]
		list = append(list, head.Val)
		head = head.Right
	}

	return list[k-1]
}

func KthSmallestElem() {
	// Example BST:
	//      3
	//     / \
	//    1   4
	//     \
	//      2
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	k := 2
	result := kthSmallest(root, k)
	println("The", k, "th smallest element is:", result)
}

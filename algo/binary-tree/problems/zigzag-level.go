// Binary Tree Zigzag Level Order Traversal
// LeetCode: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
// Given the root of a binary tree, return the zigzag level order traversal of its nodes' values.
// (i.e., from left to right, then right to left for the next level and alternate between).
//
// Example:
// Input: [3,9,20,null,null,15,7]
// Output: [[3],[20,9],[15,7]]
//
// The function zigzagLevelOrder traverses the tree level by level (BFS) and alternates the order
// of traversal for each level to achieve the zigzag pattern.

package problems

import "fmt"

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	isLeft := true
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := len(queue)
		list := make([]int, level)
		for i := 0; i < level; i++ {
			node := queue[0]
			queue = queue[1:]

			index := i
			if !isLeft {
				index = level - 1 - i
			}

			list[index] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, list)
		isLeft = !isLeft
	}
	return result
}

func ZigZagLevel() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println(zigzagLevelOrder(root)) // Output: [[3], [20, 9], [15, 7]]
}

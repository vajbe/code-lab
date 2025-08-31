/*
Problem: Spiral Matrix

Description:
Given an m x n matrix, return all elements of the matrix in spiral order.

Example:
Input:
[
 [1, 2, 3],
 [4, 5, 6],
 [7, 8, 9]
]
Output: [1, 2, 3, 6, 9, 8, 7, 4, 5]

Approach (Pseudo code):
1. Initialize four pointers: top, bottom, left, right.
   - top = 0, bottom = m-1
   - left = 0, right = n-1
2. While top <= bottom and left <= right:
   - Traverse from left → right across the top row, then increment top.
   - Traverse from top → bottom along the right column, then decrement right.
   - If top <= bottom: traverse from right → left across the bottom row, then decrement bottom.
   - If left <= right: traverse from bottom → top along the left column, then increment left.
3. Continue until all layers are traversed.

Time Complexity: O(m * n), each element visited once.
Space Complexity: O(m * n) for result storage, O(1) extra.

Patterns:
- Matrix traversal
- Layer-by-layer peeling
- Directional simulation

*/

package problems

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	res := []int{}

	for top <= bottom && left <= right {
		// left → right
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		// top → bottom
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		// right → left
		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		// bottom → top
		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}

	return res
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(spiralOrder(matrix)) // Output: [1 2 3 6 9 8 7 4 5]
}

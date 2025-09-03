package problems

import "fmt"

/*
Problem: 48. Rotate Image

Description:
------------
You are given an n x n 2D matrix representing an image.
Rotate the image by 90 degrees (clockwise).
The rotation must be done in-place (no extra matrix should be used).

Example:
--------
Input:
[
 [1, 2, 3],
 [4, 5, 6],
 [7, 8, 9]
]

Output:
[
 [7, 4, 1],
 [8, 5, 2],
 [9, 6, 3]
]

-------------------------------------------------------------------------------
Approach:
---------
1. Rotation by 90° clockwise can be broken into two steps:
   - Step 1: Transpose the matrix
             (swap matrix[i][j] with matrix[j][i] for all i < j).
   - Step 2: Reverse each row
             (swap elements in row[i] from left to right).

2. This ensures in-place rotation without extra space.

-------------------------------------------------------------------------------
Pseudocode:
-----------
function rotate(matrix):
    n = length(matrix)

    # Step 1: Transpose
    for i in range(0, n):
        for j in range(i+1, n):
            swap(matrix[i][j], matrix[j][i])

    # Step 2: Reverse rows
    for i in range(0, n):
        reverse(matrix[i])

-------------------------------------------------------------------------------
Complexity:
-----------
Time Complexity:  O(n^2)  -> we visit each cell at most twice
Space Complexity: O(1)    -> in-place modification, no extra 2D matrix

-------------------------------------------------------------------------------
Similar Patterns (Matrix Manipulation):
---------------------------------------
- Transpose Matrix (LC 867)
- Spiral Matrix (LC 54, LC 59)
- Set Matrix Zeroes (LC 73)
- Game of Life (LC 289)
- Word Search (LC 79)
- Image Smoother (LC 661)

-------------------------------------------------------------------------------
Commit Message Example:
-----------------------
feat(matrix): implement in-place 90 degree clockwise rotation

- Added transpose step (swap [i][j] with [j][i])
- Added row reversal step
- Time: O(n^2), Space: O(1)
- Verified with sample input/output
*/

func rotate(matrix [][]int) {
	// Transpose matrix
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// reverse rows
	n := len(matrix[0])
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0])/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}

}

func RotateIamge() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rotate(matrix)
	fmt.Println(matrix)
}

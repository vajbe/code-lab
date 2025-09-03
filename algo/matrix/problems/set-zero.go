package problems

import "fmt"

/*
Problem: 73. Set Matrix Zeroes

Description:
------------
Given an m x n integer matrix, if an element is 0, set its entire row and column to 0's.
You must do it in place.

Example 1:
----------
Input:
[
 [1, 1, 1],
 [1, 0, 1],
 [1, 1, 1]
]

Output:
[
 [1, 0, 1],
 [0, 0, 0],
 [1, 0, 1]
]

Example 2:
----------
Input:
[
 [0, 1, 2, 0],
 [3, 4, 5, 2],
 [1, 3, 1, 5]
]

Output:
[
 [0, 0, 0, 0],
 [0, 4, 5, 0],
 [0, 3, 1, 0]
]

-------------------------------------------------------------------------------
Approach:
---------
1. Naive (extra space O(m+n)):
   - Keep two arrays: rows[m], cols[n].
   - Traverse the matrix → mark row[i] = 0 and col[j] = 0 if matrix[i][j] == 0.
   - Traverse again → set cell = 0 if row[i] == 0 or col[j] == 0.

2. Optimal (O(1) extra space):
   - Use the **first row and first column** of the matrix as markers.
   - Two additional booleans track if the **first row** or **first column** should be zeroed.
   - Traverse the matrix (except first row/col):
       if matrix[i][j] == 0 → mark matrix[i][0] = 0, matrix[0][j] = 0
   - Traverse again and zero out based on markers.
   - Finally, handle the first row/col separately.

-------------------------------------------------------------------------------
Pseudocode:
-----------
function setZeroes(matrix):
    m, n = dimensions of matrix
    firstRowZero, firstColZero = false, false

    # Step 1: Check if first row/col should be zero
    if any(matrix[0][j] == 0) → firstRowZero = true
    if any(matrix[i][0] == 0) → firstColZero = true

    # Step 2: Use first row/col as markers
    for i in range(1..m-1):
        for j in range(1..n-1):
            if matrix[i][j] == 0:
                matrix[i][0] = 0
                matrix[0][j] = 0

    # Step 3: Zero cells based on markers
    for i in range(1..m-1):
        for j in range(1..n-1):
            if matrix[i][0] == 0 or matrix[0][j] == 0:
                matrix[i][j] = 0

    # Step 4: Zero first row/col if needed
    if firstRowZero:
        set all elements in row[0] = 0
    if firstColZero:
        set all elements in col[0] = 0

-------------------------------------------------------------------------------
Complexity:
-----------
Time Complexity:  O(m * n) → each cell processed at most twice
Space Complexity: O(1)     → in-place, only a few flags used

-------------------------------------------------------------------------------
Similar Patterns:
-----------------
- Rotate Image (LC 48)
- Game of Life (LC 289)
- Flood Fill (LC 733)
- Surrounded Regions (LC 130)
- Matrix Diagonal Sum (LC 1572)

-------------------------------------------------------------------------------
Commit Message Example:
-----------------------
feat(matrix): implement in-place matrix zeroing

- optimized approach using first row/col as markers
- reduced extra space from O(m+n) → O(1)
- time: O(m*n), space: O(1)
- tested with multiple examples
*/

func setZeroes(matrix [][]int) {

	n, m := len(matrix), len(matrix[0])
	firstColumn, firstRow := false, false

	for i := 0; i < n; i++ {
		if matrix[i][0] == 0 {
			firstColumn = true
		}
	}

	for i := 0; i < m; i++ {
		if matrix[0][i] == 0 {
			firstRow = true
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstColumn {
		for j := 0; j < n; j++ {
			matrix[j][0] = 0
		}
	}

	if firstRow {
		for i := 0; i < m; i++ {
			matrix[0][i] = 0
		}
	}

	fmt.Print(matrix)

}

/* func setZeroes(matrix [][]int) {
	//Brute Force
	rowCount := len(matrix)
	colCount := len((matrix[0]))
	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			if matrix[i][j] == 0 {
				row := i
				column := j
				for c := 0; c < colCount; c++ {
					if matrix[row][c] != 0 {
						matrix[row][c] = math.MaxInt
					}

				}

				for r := 0; r < rowCount; r++ {
					if matrix[r][column] != 0 {
						matrix[r][column] = math.MaxInt
					}

				}
			}
		}
	}

	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			if matrix[i][j] == math.MaxInt {
				matrix[i][j] = 0
			}
		}
	}

	fmt.Print(matrix)
}
*/

func SetZero() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	setZeroes(matrix)
}

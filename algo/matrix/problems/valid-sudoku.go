/*
Problem: Valid Sudoku

Description:
Given a 9x9 Sudoku board, determine if it is valid. Only the filled cells need
to be validated according to these rules:
  1. Each row must not contain duplicate digits.
  2. Each column must not contain duplicate digits.
  3. Each of the nine 3x3 sub-boxes must not contain duplicate digits.

Note:
- The board does not need to be solvable.
- Cells are either '1'-'9' or '.' representing empty.

Approach (Pseudo code):
1. Create 9 sets/maps for rows, 9 for columns, and 9 for 3x3 boxes.
2. Traverse each cell (r, c):
   - Skip if value is '.'.
   - Check if digit already exists in row[r], col[c], or box[(r/3)*3 + c/3].
     → If yes, return false.
   - Otherwise, add digit to these sets/maps.
3. If no duplicates found, return true.

Time Complexity:
- O(9x9) = O(1) (constant work, since board size is fixed).
- More generally, O(N^2) for an N x N board.

Space Complexity:
- O(27N) in worst case (9 rows + 9 cols + 9 boxes storing up to N digits).
- For 9x9, it’s effectively O(1).

Patterns:
- Hashing / Set-based duplicate detection.
- Matrix traversal.
- Constraint validation.

*/

package problems

import "fmt"

func isValidSudoku(board [][]byte) bool {

	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	box := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		box[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			val := board[r][c]

			if val == '.' {
				continue
			}
			row := rows[r]
			if row[val] {
				return false
			}
			col := cols[c]
			if col[val] {
				return false
			}
			pos := (r/3)*3 + (c / 3)
			bx := box[pos]
			if bx[val] {
				return false
			}

			row[val] = true
			col[val] = true
			bx[val] = true
		}
	}

	return true
}
func ValidSudoku() {
	sudoku := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Print("Is valid: ", isValidSudoku(sudoku))
}

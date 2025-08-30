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

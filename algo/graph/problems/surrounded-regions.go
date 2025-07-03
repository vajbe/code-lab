// Surrounded Regions (LeetCode 130)
// ----------------------------------
// Given an m x n matrix board containing 'X' and 'O', capture all regions surrounded by 'X'.
// A region is captured by flipping all 'O's into 'X's in that surrounded region.
// Any 'O' on the border, or connected to an 'O' on the border, is not captured.
//
// Example:
// Input: board = [
//
//	['X','O','X','O','X','O'],
//	['O','X','O','X','O','X'],
//	['X','O','X','O','X','O'],
//	['O','X','O','X','O','X']
//
// ]
// Output: (no change, as all 'O's are on the border or connected to the border)
//
// Approach:
// - Use DFS to mark all 'O's connected to the border as safe (temporarily mark as '#').
// - Flip all remaining 'O's to 'X' (these are surrounded).
// - Convert all '#' back to 'O'.
package problems

import "fmt"

func regionDfs(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	var dfs func(int, int)
	r, c := len(board), len(board[0])
	dfs = func(m, n int) {
		if m < 0 || m >= r || n < 0 || n >= c || board[m][n] != 'O' {
			return
		}
		board[m][n] = '#' //marking safe
		dfs(m+1, n)
		dfs(m-1, n)
		dfs(m, n+1)
		dfs(m, n-1)
	}

	for i := 0; i < r; i++ {
		dfs(i, 0)
		dfs(i, c-1)
	}

	for i := 0; i < c; i++ {
		dfs(0, i)
		dfs(r-1, i)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'

			case '#':
				board[i][j] = 'O'

			}
		}
	}

}

func solve(board [][]byte) {
	regionDfs(board)
}

func SurroundedRegions() {
	board := [][]byte{
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
	}
	solve(board)
	// Print the board to verify the result
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
}

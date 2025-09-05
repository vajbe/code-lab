package problems

import "fmt"

// 1 -> 1 	live neigh == 2,3
// 1 -> 0	live neigh > 3
// 0 -> 1 	live neigh === 3
// 1 -> 0	live neigh < 2
var directions = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func countLives(board [][]int, row, column int) int {
	count := 0

	for _, d := range directions {
		rw, col := row+d[0], column+d[1]

		if rw < 0 || col < 0 || rw >= len(board) || col >= len(board[0]) {
			continue
		}
		count += board[rw][col]
	}
	return count
}

func gameOfLifeExtraSpace(board [][]int) {
	r, c := len(board), len(board[0])
	newBoard := make([][]int, r)

	for i := range newBoard {
		newBoard[i] = make([]int, c)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			lives := countLives(board, i, j)

			if board[i][j] == 1 {
				if lives == 2 || lives == 3 {
					newBoard[i][j] = 1 // survives
				} else {
					newBoard[i][j] = 0 // dies
				}
			} else {
				if lives == 3 {
					newBoard[i][j] = 1 // reproduction
				}
			}
		}
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			board[i][j] = newBoard[i][j]
		}
	}
	fmt.Println(board)
}

func GameLife() {
	gameOfLifeExtraSpace([][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}})
}

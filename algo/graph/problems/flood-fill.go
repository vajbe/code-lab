package problems

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	var dfs func(int, int)
	rows, cols := len(image), len(image[0])
	col := image[sr][sc]
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= rows || j >= cols {
			return
		}
		if image[i][j] != color && image[i][j] == col {
			image[i][j] = color
			dfs(i+1, j)
			dfs(i-1, j)
			dfs(i, j+1)
			dfs(i, j-1)
		}

	}

	dfs(sr, sc)

	return image
}

func FloodFill() {
	// board := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	// floodFill(board, 1, 1, 2)
	board := [][]int{{2, 2, 2}, {2, 2, 2}}
	floodFill(board, 1, 0, 2)

	for _, row := range board {
		for _, cell := range row {
			fmt.Print(" ", cell)
		}
		fmt.Println()
	}

}

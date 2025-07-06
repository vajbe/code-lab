package problems

import "fmt"

func numIslands(grid [][]byte) int {

	var dfs func(int, int)
	rows := len(grid)
	cols := len(grid[0])
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= rows || j >= cols || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	num := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				num++
				dfs(i, j)
			}
		}
	}

	return num
}

func NumIsland() {
	c1 := numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	})
	fmt.Println(c1)

	c2 := numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	})
	fmt.Println(c2)

}

/*
Problem:
We have a rectangular garden represented as a grid of size height x width.
Each cell contains a number of leaves. A series of wind gusts pushes the leaves one cell at a time
in a specific direction:
- 'U' → up
- 'D' → down
- 'L' → left
- 'R' → right

When leaves are pushed out of the grid, they are lost. The task is to simulate the wind gusts
in order and return the number of leaves remaining in the garden.

Inputs:
- width, height: integers (0 < width, height < 20)
- leaves: 2D integer array representing initial number of leaves in each cell
- winds: string containing the gust directions (0 <= len(winds) < 20)

Approach:
1. For each wind direction in the sequence:
  - Create a new grid of the same size initialized with zeros.
  - Move leaves from each cell one step in the wind direction.
  - If the new position is outside the grid, the leaves are lost.
  - Otherwise, add them to the correct cell in the new grid.

2. After processing all wind gusts, sum all the leaves remaining in the grid.

Time Complexity:
O(H * W * G)
- H = height
- W = width
- G = number of gusts
We process each cell once for each gust.

Space Complexity:
O(H * W) for storing the grid state.

Category:
Simulation, Grid processing, Array manipulation.

Similar Questions:
- LeetCode 289: Game of Life (grid state updates)
- Simulation problems involving grid movement
- "Matrix shifting" style puzzles
*/
package main

import (
	"fmt"
)

func remainingLeaves(width, height int, leaves [][]int, winds string) int {
	for _, gust := range winds {
		newGrid := make([][]int, height)
		for i := range newGrid {
			newGrid[i] = make([]int, width)
		}

		for r := 0; r < height; r++ {
			for c := 0; c < width; c++ {
				if leaves[r][c] > 0 {
					nr, nc := r, c
					switch gust {
					case 'U':
						nr--
					case 'D':
						nr++
					case 'L':
						nc--
					case 'R':
						nc++
					}

					// If still inside grid, move leaves
					if nr >= 0 && nr < height && nc >= 0 && nc < width {
						newGrid[nr][nc] = leaves[r][c]
					}
				}
			}
		}

		leaves = newGrid
	}

	// Count remaining leaves
	total := 0
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			total += leaves[r][c]
		}
	}
	return total
}

func FindRemLeaves() {
	width := 4
	height := 4
	leaves := [][]int{
		{1, 1, 0, 1},
		{0, 0, 2, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	winds := "LRDLD"

	fmt.Println(remainingLeaves(width, height, leaves, winds)) // Output: 3
}

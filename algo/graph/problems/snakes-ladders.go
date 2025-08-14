/*
Package main - Snakes and Ladders (LeetCode 909)

Description:
------------
This program solves the "Snakes and Ladders" problem.
Given a 2D board with -1 for empty squares, and numbers for snakes or ladders
(destinations), we must determine the minimum number of dice rolls needed to
reach the last square starting from square 1. Movement proceeds in zig-zag
row order from bottom-left to top-right.

Approach:
---------
- Use Breadth-First Search (BFS) starting from square 1.
- Each BFS node stores the current square number and the count of dice rolls
  taken so far.
- For each dice roll (1 to 6), calculate the next square:
    - Apply snakes/ladders immediately if present.
    - If we reach the final square, return moves + 1 (current move + this roll).
- Maintain a visited array to avoid revisiting squares.
- The moment we reach the final square, BFS guarantees it’s the minimum count.

Time Complexity:
----------------
O(N^2) where N is the board dimension (since there are N^2 squares and each is visited at most once).

Space Complexity:
-----------------
O(N^2) for the visited array and BFS queue.

Example:
--------
Board:
[
  [-1, -1, -1, -1, -1, -1],
  [-1, -1, -1, -1, -1, -1],
  [-1, -1, -1, -1, -1, -1],
  [-1, 35, -1, -1, 13, -1],
  [-1, -1, -1, -1, -1, -1],
  [-1, 15, -1, -1, -1, -1]
]
Output: 4 (minimum dice rolls to reach the last square)
*/

package problems

import "fmt"

func snakesAndLadders(board [][]int) int {
	l := len(board)
	queue := [][]int{{1, 0}}
	visited := make([]bool, l*l+1)
	visited[1] = true

	flatten := func(num int) (int, int) {
		r := (num - 1) / l
		c := (num - 1) % l

		if r%2 == 1 {
			c = l - 1 - c
		}
		return l - 1 - r, c
	}

	for len(queue) > 0 {
		square, moves := queue[0][0], queue[0][1]
		queue = queue[1:]

		for dice := 1; dice <= 6; dice++ {
			next := square + dice
			r, c := flatten(next)

			if board[r][c] != -1 {
				next = board[r][c]
			}

			if next == l*l {
				return moves + 1
			}

			if !visited[next] {
				queue = append(queue, []int{next, moves + 1})
				visited[next] = true
			}
		}
	}

	return -1
}

func SnakeLad() {
	board := [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1},
	}

	result := snakesAndLadders(board)
	fmt.Println("Minimum moves to reach the end:", result)

}

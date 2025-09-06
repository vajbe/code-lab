/*
🧩 Problem: Game of Life (LeetCode 289)

We are given an m x n grid representing Conway's Game of Life, where:
- 1 = live cell
- 0 = dead cell

Each cell interacts with its 8 neighbors (horizontally, vertically, diagonally).
The transition rules are:
1. Any live cell with fewer than 2 live neighbors → dies (underpopulation).
2. Any live cell with 2 or 3 live neighbors → lives (survival).
3. Any live cell with more than 3 live neighbors → dies (overpopulation).
4. Any dead cell with exactly 3 live neighbors → becomes a live cell (reproduction).

We must compute the next state of the board **in-place**.

---

💡 Intuition:
If we update directly, neighbor counts will get corrupted (mix of old/new states).
To avoid this, we encode transitions using temporary markers:
-  1 → was alive, becomes dead → mark as -1
-  0 → was dead, becomes alive → mark as 2

This way:
- abs(cell) gives us the original state
- sign / encoded value gives us the new state

At the end, we normalize:
- > 0 → alive (1)
- <= 0 → dead (0)

---

🔑 Pseudo Code:
function gameOfLife(board):
    for each cell (i,j):
        count = live neighbors
        if cell alive:
            if count < 2 or count > 3: mark -1
        else:
            if count == 3: mark 2
    for each cell:
        normalize: >0 → 1, else → 0

---

📊 Complexity:
- Time:  O(m * n)  (each cell checks 8 neighbors → O(8*m*n) ≈ O(m*n))
- Space: O(1)      (in-place, only constant extra space)

---

🔗 Similar Problems:
- LeetCode 289 – Game of Life
- LeetCode 73 – Set Matrix Zeroes (similar in-place trick)
- Grid simulation problems (zombie infection, fire spreading, etc.)

---
✅ Example:
Input:
[[0,1,0],
 [0,0,1],
 [1,1,1],
 [0,0,0]]

Output:
[[0,0,0],
 [1,0,1],
 [0,1,1],
 [0,1,0]]
*/

// ---------------- Go Solution ----------------
package problems

func gameOfLife(board [][]int) {
	r, c := len(board), len(board[0])

	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// Count live neighbors
	countLives := func(row, col int) int {
		count := 0
		for _, d := range directions {
			rw, cl := row+d[0], col+d[1]
			if rw >= 0 && cl >= 0 && rw < r && cl < c {
				if abs(board[rw][cl]) == 1 { // use abs to read original state
					count++
				}
			}
		}
		return count
	}

	// First pass: apply rules with markers
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			lives := countLives(i, j)
			if board[i][j] == 1 {
				if lives < 2 || lives > 3 {
					board[i][j] = -1 // live → dead
				}
			} else {
				if lives == 3 {
					board[i][j] = 2 // dead → live
				}
			}
		}
	}

	// Second pass: normalize
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if board[i][j] > 0 {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GameLife() {
	gameOfLife([][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}})
}

// ------------------------------------------------
// Given an undirected graph with n nodes and an array edges where each edge[i] = [ui, vi] represents a bidirectional edge between nodes ui and vi,
// and two integers source and destination, return true if there is a valid path from source to destination, or false otherwise.
//
// Example:
// Input: n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
// Output: true
//
// Approach:
// - Build the graph using an adjacency matrix or adjacency list.
// - Use DFS or BFS to check if destination is reachable from source.

package problems

import "fmt"

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := make([][]int, n)
	for i := 0; i < len(edges); i++ {
		from := edges[i][0]
		to := edges[i][1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	stack := make([]int, 0)
	visited := make([]bool, n)

	stack = append(stack, source)

	for len(stack) > 0 {
		el := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !visited[el] {
			visited[el] = true
			if el == destination {
				return true
			}
		}

		for _, neighbor := range graph[el] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}

	return false
}

func ValidPath() {
	fmt.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
	fmt.Println(validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5))
}

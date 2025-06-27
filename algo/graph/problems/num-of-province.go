// Number of Provinces (LeetCode 547)
// -----------------------------------
// There are n cities. Some of them are connected, while some are not.
// A province is a group of directly or indirectly connected cities and no other cities outside of the group.
// You are given an n x n adjacency matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected, and 0 otherwise.
// Return the total number of provinces.
//
// Example:
// Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
// Output: 2
//
// Approach:
// - Use DFS or BFS to count the number of connected components in the

package problems

import "fmt"

var visitedProvince []bool

func dfsProvince(startNode int, graph [][]int) {

	stack := make([]int, 0)
	stack = append(stack, startNode)
	for len(stack) > 0 {
		el := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !visitedProvince[el] {
			visitedProvince[el] = true
			// fmt.Print("\t", el)
		}

		for i := 0; i < len(graph[el]); i++ {
			if graph[el][i] == 1 && !visitedProvince[i] {
				stack = append(stack, i)
			}
		}
	}
}

func findCircleNum(isConnected [][]int) int {
	nodes := len(isConnected)
	numProvince := 0
	n := len(isConnected)
	visitedProvince = make([]bool, n)
	for i := 0; i < nodes; i++ {
		if !visitedProvince[i] {
			numProvince++
			// fmt.Print("\nDFS for ", i)
			dfsProvince(i, isConnected)
		}
	}
	return numProvince
}

func NumProvince() {
	fmt.Print(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
}

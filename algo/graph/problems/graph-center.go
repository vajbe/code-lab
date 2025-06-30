// ---------------------------------------------
// A star graph is a graph where there is one center node and exactly n-1 leaves.
// You are given a 2D integer array edges where each edges[i] = [ui, vi] indicates that there is an edge between the nodes ui and vi.
// Return the center of the given star graph.
//
// Example:
// Input: edges = [[1,2],[2,3],[4,2]]
// Output: 2
//
// Approach:
// - The center node will appear in every edge.
// - Check the first two edges; the common node is the center.
package problems

import "fmt"

func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}
	return edges[0][1]
}

func FindCenterOfGraph() {
	fmt.Println(findCenter([][]int{{1, 2}, {2, 3}, {4, 2}}))
	fmt.Println(findCenter([][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}))
}

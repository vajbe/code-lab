// Evaluate Division Problem
// ---------------------------------------------
// Given a list of equations representing relationships between variables (e.g., a / b = 2.0),
// and a list of queries (e.g., a / c), this algorithm computes the result of each query
// based on the given equations.
//
// Approach:
// 1. Graph Construction:
//    - Each variable is a node in a graph.
//    - Each equation (a / b = k) adds two directed edges: a->b (weight k) and b->a (weight 1/k).
// 2. Query Evaluation:
//    - For each query (start, end), perform DFS from start to end, multiplying edge weights along the path.
//    - If a path exists, return the product; otherwise, return -1.0.
//
// Example:
//   equations = [["a","b"],["b","c"]], values = [2.0,3.0]
//   queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
//   Output: [6.0,0.5,-1.0,1.0,-1.0]
//
// Time Complexity: O(N + Q * V), where N is the number of equations, Q is the number of queries, and V is the number of variables.

package problems

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	var buildGraph func()
	var dfs func(string, string, map[string]bool, float64) float64
	var graph map[string]map[string]float64
	var printG func()
	var results []float64
	buildGraph = func() {
		graph = make(map[string]map[string]float64)

		for i, eq := range equations {
			from, to := eq[0], eq[1]

			if graph[from] == nil {
				graph[from] = make(map[string]float64)
			}

			if graph[to] == nil {
				graph[to] = make(map[string]float64)
			}

			graph[from][to] = values[i]
			graph[to][from] = 1 / values[i]

		}
	}

	dfs = func(startNode, endNode string, visited map[string]bool, product float64) float64 {
		if _, ok := graph[startNode]; !ok {
			return -1.000
		}
		if startNode == endNode {
			return product
		}

		visited[startNode] = true

		for neighbor, val := range graph[startNode] {
			if !visited[neighbor] {
				res := dfs(neighbor, endNode, visited, product*val)
				if res != -1.000 {
					return res
				}
			}
		}

		return -1.00
	}

	printG = func() {
		/* for key, value := range graph {
			fmt.Print("Node :\t", key)
			for k, v := range value {
				fmt.Print("\t", k, ":", v)
			}
			fmt.Print("\n")
		} */
	}

	buildGraph()
	printG()
	for _, eq := range queries {
		visited := make(map[string]bool)
		results = append(results, dfs(eq[0], eq[1], visited, 1.000))
	}
	return results
}

func CalculateEquations() {
	// Sample data
	equations := [][]string{
		{"a", "b"},
		{"b", "c"},
	}
	values := []float64{2.0, 3.0}
	queries := [][]string{
		{"a", "c"},
		{"b", "a"},
		{"a", "e"},
		{"a", "a"},
		{"x", "x"},
	}

	result := calcEquation(equations, values, queries)
	fmt.Println("Results:", result)
}

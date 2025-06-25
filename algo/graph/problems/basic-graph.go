package problems

import "fmt"

func adjacencyMatrixGraph(vertices []int, edges [][]int) {
	n := len((vertices)) + 1
	graph := make([][]int, n)

	for i := 1; i < n; i++ {
		graph[i] = make([]int, n)
	}

	var addNode func(int, int)
	var printGraph func()

	addNode = func(from int, to int) {
		graph[from][to] = 1
		graph[to][from] = 1
	}

	printGraph = func() {
		fmt.Print("Rows:  ")
		for i := 1; i < n; i++ {
			fmt.Print(" ", i, "")
		}
		fmt.Println()
		for i := 1; i < n; i++ {
			fmt.Print("Node ", i, ":")
			for j := 1; j < n; j++ {
				fmt.Print(" ", graph[i][j])
			}
			fmt.Println()
		}
	}

	for i := 0; i < len(edges); i++ {
		addNode(edges[i][0], edges[i][1])
	}
	printGraph()
}

func BaseGraph() {
	edges := [][]int{{1, 2}, {3, 4}, {1, 3}, {2, 4}}
	vertices := []int{1, 2, 3, 4}
	adjacencyMatrixGraph(vertices, edges)
}

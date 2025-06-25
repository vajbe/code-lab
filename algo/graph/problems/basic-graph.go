package problems

import "fmt"

var graph [][]int

func dfs(startNode int) {
	var isVisited func(int) bool
	n := len(graph)
	visited := make([]bool, n)

	isVisited = func(el int) bool {
		return visited[el]
	}

	stack := make([]int, 0)
	stack = append(stack, startNode)
	fmt.Print("DFS: ->")
	for len(stack) > 0 {
		el := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !isVisited(el) {
			visited[el] = true
			fmt.Print("\t", el)
		}

		for i := 1; i < len(graph[el]); i++ {
			if graph[el][i] == 1 && !isVisited(i) {
				stack = append(stack, i)
			}
		}
	}
}

func printGraph() {
	n := len(graph)
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

func adjacencyMatrixGraph(vertices []int, edges [][]int) {
	var addNode func(int, int)
	n := len((vertices)) + 1
	graph = make([][]int, n)

	for i := 1; i < n; i++ {
		graph[i] = make([]int, n)
	}

	addNode = func(from int, to int) {
		graph[from][to] = 1
		graph[to][from] = 1
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
	dfs(2)
}

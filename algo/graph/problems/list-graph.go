package problems

import "fmt"

func traverseListDFS(graph [][]int, startNode int) {
	stack := make([]int, 0)
	visited := make([]bool, len(graph))

	stack = append(stack, startNode)
	fmt.Print("DFS =>")
	for len(stack) > 0 {
		el := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[el] {
			visited[el] = true
			fmt.Print("\t", el)
		}

		for _, neighbour := range graph[el] {
			if !visited[neighbour] {
				stack = append(stack, neighbour)
			}
		}
	}
}

func traverseListDFSRec(graph [][]int, startNode int, visited []bool) {
	if startNode < 0 || startNode >= len(graph) {
		return
	}
	visited[startNode] = true
	fmt.Print("\t", startNode)
	for _, neigh := range graph[startNode] {
		if !visited[neigh] {
			traverseListDFSRec(graph, neigh, visited)
		}
	}
}

func traverseListBFS(graph [][]int, startNode int) {
	queue := make([]int, 0)
	visited := make([]bool, len(graph))

	queue = append(queue, startNode)
	fmt.Print("\nBFS =>")
	for len(queue) > 0 {
		el := queue[0]
		queue = queue[1:]

		if !visited[el] {
			visited[el] = true
			fmt.Print("\t", el)
		}

		for _, neigh := range graph[el] {
			if !visited[neigh] {
				queue = append(queue, neigh)
			}
		}
	}
}

func printListGraph(graph [][]int) {
	for i := 0; i < len(graph); i++ {
		fmt.Print(i, "")
		for j := 0; j < len(graph[i]); j++ {
			fmt.Print("\t->", graph[i][j])
		}
		fmt.Println()
	}
}

func createGraph(n int, edges [][]int) [][]int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < len(edges); i++ {
		from := edges[i][0]
		to := edges[i][1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}
	return graph
}

func AdjacencyListGraph() {
	graph := createGraph(5, [][]int{{0, 1}, {0, 3}, {1, 2}, {3, 4}, {4, 2}})
	printListGraph(graph)
	traverseListDFS(graph, 0)
	traverseListBFS(graph, 0)
	visited := make([]bool, len(graph))
	fmt.Print("\nRDFS")
	traverseListDFSRec(graph, 0, visited)
}

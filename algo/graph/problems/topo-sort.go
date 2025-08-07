package problems

import "fmt"

func TopoSortExample() {

	var buildGraph func([][]string)
	var printGraph func()
	var graph map[string][]string
	var topoSortVisit func(string, *[]string, map[string]bool)
	var topoSort func()
	buildGraph = func(edges [][]string) {
		graph = make(map[string][]string)
		for _, edge := range edges {
			from := edge[0]
			to := edge[1]
			if graph[from] == nil {
				graph[from] = make([]string, 0)
			}
			graph[from] = append(graph[from], to)
		}
	}

	printGraph = func() {
		for key, value := range graph {
			fmt.Print("\n Edge: ", key, " ->")
			for _, v := range value {
				fmt.Print(" ", v)
			}
		}
	}

	topoSortVisit = func(vertex string, stack *[]string, visited map[string]bool) {
		for _, val := range graph[vertex] {
			if !visited[val] {
				topoSortVisit(val, stack, visited)
			}
		}
		visited[vertex] = true
		*stack = append(*stack, vertex)
	}

	topoSort = func() {
		visited := make(map[string]bool)
		stack := []string{}
		for node := range graph {
			if !visited[node] {
				topoSortVisit(node, &stack, visited)
			}
		}

		fmt.Print("\nSorted Topo ->")
		for i := len(stack) - 1; i >= 0; i-- {
			fmt.Print(" ", stack[i])
		}
	}

	buildGraph([][]string{{"A", "C"},
		{"B", "C"}, {"B", "D"},
		{"C", "E"},
		{"D", "F"}, {"E", "H"}, {"E", "F"}, {"F", "G"},
	})
	printGraph()
	topoSort()
}

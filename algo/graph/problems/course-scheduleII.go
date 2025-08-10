/*
Course Schedule II (LeetCode #210)

Description:
You are given the total number of courses `numCourses` you need to take, labeled from 0 to numCourses - 1,
and an array `prerequisites` where prerequisites[i] = [a, b] means you must take course `b` before course `a`.
Return an ordering of courses you should take to finish all courses, or an empty array if it is impossible.

The problem boils down to finding a topological ordering of a directed graph formed by the courses and prerequisites.
If the graph contains a cycle, it is impossible to complete all courses.

Approach (DFS with Stack):
1. Build an adjacency list to represent the directed graph.
2. Use a visited array with three states:
   - 0 = unvisited
   - 1 = visiting (in current recursion stack)
   - 2 = visited (fully processed)
3. Perform DFS on each unvisited course:
   - If a node is marked as visiting and is reached again → cycle detected → return [].
   - After exploring all prerequisites of a course, push it onto a stack.
4. After DFS finishes for all courses, pop from the stack to get the topological order.

Time Complexity: O(V + E) where V = numCourses and E = number of prerequisite pairs.
Space Complexity: O(V + E) for the adjacency list, visited array, and recursion stack.
*/

package problems

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	var buildGraph func()
	var graph [][]int
	var hasCycle func(int) bool
	visited := make([]int, numCourses)
	stack := []int{}
	buildGraph = func() {
		graph = make([][]int, numCourses)
		for _, val := range prerequisites {
			r, c := val[0], val[1]
			if graph[r] == nil {
				graph[r] = []int{}
			}
			graph[r] = append(graph[r], c)
		}
	}

	var printGraph = func() {
		for i := 0; i < len(graph); i++ {
			fmt.Print("\n Node : ", i, " -> ")
			for j := 0; graph[i] != nil && j < len(graph[i]); j++ {
				fmt.Print(" ", graph[i][j])
			}
		}
	}

	hasCycle = func(node int) bool {
		if visited[node] == 1 {
			return true
		}
		if visited[node] == 2 {
			return false
		}

		visited[node] = 1

		for _, neight := range graph[node] {
			if hasCycle(neight) {
				return true
			}
		}

		visited[node] = 2
		stack = append(stack, node)
		return false
	}

	buildGraph()
	printGraph()
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 && hasCycle(i) {
			return []int{}
		}
	}

	return stack
}

func CourseScheduleII() {
	/* 	fmt.Print("\n", findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	   	fmt.Print("\n", findOrder(2, [][]int{{1, 0}}))
	   	fmt.Print("\n", findOrder(1, [][]int{})) */
	fmt.Print("\n", findOrder(3, [][]int{{1, 0}, {1, 2}, {0, 1}}))
}

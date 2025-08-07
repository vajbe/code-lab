/*
Problem: Course Schedule (Leetcode 207)
---------------------------------------

You are given `numCourses` representing the total number of courses labeled from 0 to numCourses - 1,
and a list of prerequisite pairs where prerequisites[i] = [a, b] indicates that
you must take course `b` before course `a`.

Determine if it is possible to finish all courses. This is a classic cycle detection
problem in a directed graph.

Example 1:
Input: numCourses = 2, prerequisites = [[1, 0]]
Output: true
Explanation: Take course 0 → then course 1

Example 2:
Input: numCourses = 2, prerequisites = [[1, 0], [0, 1]]
Output: false
Explanation: A cycle exists: 0 → 1 → 0

---------------------------------------
Approach:
- Represent the courses and their dependencies as a directed graph.
- Use Depth-First Search (DFS) to detect if a cycle exists in the graph.
- During DFS traversal, maintain a `visited` array with 3 states:
    0 = not visited
    1 = visiting (in recursion stack)
    2 = fully visited (done processing)
- If a node is revisited while still in the visiting state, a cycle is detected.
- If no cycle is detected after traversing all nodes, return true.

---------------------------------------
Time Complexity:
- O(V + E), where V = number of courses, E = number of prerequisite pairs.

Space Complexity:
- O(V + E) for graph representation.
- O(V) for the visited array and recursion stack.

---------------------------------------
Function Signature:
func canFinish(numCourses int, prerequisites [][]int) bool

---------------------------------------
*/

package problems

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	var buildGraph func()
	var hasCycle func(int) bool

	var visited []int

	buildGraph = func() {

		for _, val := range prerequisites {
			from := val[0]
			to := val[1]

			if graph[from] == nil {
				graph[from] = []int{}
			}
			graph[from] = append(graph[from], to)
		}
		visited = make([]int, numCourses)
	}

	hasCycle = func(node int) bool {

		if visited[node] == 1 {
			return true
		}

		if visited[node] == 2 {
			return false
		}

		visited[node] = 1

		for _, neigh := range graph[node] {
			if hasCycle(neigh) {
				return true
			}
		}

		visited[node] = 2

		return false
	}
	buildGraph()

	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 && hasCycle(i) {
			return false
		}
	}
	return true
}

func CourseSchedule() {
	fmt.Print("Can finish?: ", canFinish(2, [][]int{{1, 0}}))
	fmt.Print("\nCan finish?: ", canFinish(2, [][]int{{1, 0}, {0, 1}}))
}

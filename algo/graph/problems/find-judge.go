// Find the Town Judge (LeetCode 997)
// -----------------------------------
// In a town, there are n people labeled from 1 to n. There is a rumor that one of these people is the town judge.
// The town judge has the following properties:
//   1. The town judge trusts nobody.
//   2. Everybody (except for the town judge) trusts the town judge.
// You are given an array trust where trust[i] = [a, b] representing that the person labeled a trusts the person labeled b.
// Return the label of the town judge if the town judge exists and can be identified, or return -1 otherwise.
//
// Example:
// Input: n = 3, trust = [[1,3],[2,3]]
// Output: 3
//
// Approach:
// - Use an array to count trust relationships for each person.
// - The judge will be trusted by n-1 people and trust nobody.

package problems

import "fmt"

func findJudge(n int, trust [][]int) int {
	if n == 1 && len(trust) == 0 {
		return 1
	}
	degree := make([]int, n+1)
	numJudge := -1
	for i := 0; i < len(trust); i++ {
		if trust[i][1] != 0 {
			in := trust[i][1]
			out := trust[i][0]
			degree[in]++
			degree[out]--
		}
	}

	for i := 0; i <= n; i++ {
		if degree[i] == n-1 {
			return i
		}
	}

	return numJudge
}

func FindJudgeGraph() {
	fmt.Print("\nJudge: ", findJudge(3, [][]int{{1, 3}, {2, 3}}))
	fmt.Print("\nJudge: ", findJudge(3, [][]int{{1, 2}, {2, 3}}))
	fmt.Print("\nJudge: ", findJudge(2, [][]int{{1, 2}}))
	fmt.Print("\nJudge: ", findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
}

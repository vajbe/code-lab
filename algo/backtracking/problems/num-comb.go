// combine generates all possible combinations of k numbers chosen from the range [1..n].
//
// Problem:
// Given two integers n and k, return all possible combinations of k numbers out of the range [1..n].
//
// Example:
//   n = 4, k = 2
//   Output: [[1 2], [1 3], [1 4], [2 3], [2 4], [3 4]]
//
// Approach:
// - Use backtracking with pruning to build combinations step by step.
// - At each recursive call, we decide whether to include the current number.
// - If the current combination length == k, we take a snapshot and add it to the result.
// - We prune recursion when the remaining numbers are insufficient to reach k.
// - Backtracking ensures we explore all possible subsets while reverting choices to avoid side effects.
//
// Intuition:
// - Think of it as a decision tree where each level represents choosing the next number.
// - We branch out by including the number, recurse, then remove it (backtrack).
// - Pruning prevents exploring "dead branches" where completion is impossible.
//
// Time Complexity:
// - O(C(n, k) * k)
//   - There are C(n, k) valid combinations.
//   - Copying each combination into the result costs O(k).
//
// Space Complexity:
// - O(k) for the recursion stack and temporary combination slice.
// - O(C(n, k) * k) for storing the final results.

package problems

import (
	"fmt"
)

func combine(n int, k int) [][]int {
	var backtrack func(int, []int)
	res := [][]int{}

	backtrack = func(start int, comb []int) {
		if len(comb) == k {
			c := make([]int, len(comb))
			copy(c, comb)
			res = append(res, c)
			return
		}

		for i := start; i <= n; i++ {
			comb = append(comb, i)
			backtrack(i+1, comb)
			comb = comb[:len(comb)-1]
		}
	}

	backtrack(1, []int{})
	return res
}

func Combinations() {
	fmt.Print(combine(4, 2))
}

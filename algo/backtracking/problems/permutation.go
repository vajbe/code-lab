/*
Problem: Permutations
---------------------
Given an array of distinct integers `nums`, return all possible permutations of the array.

Example:
--------
Input:  nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Approach:
---------
Recursive Insertion Method:
1. Base case:
   - If nums has length 1, return [[nums[0]]].
2. Recursive step:
   - Get all permutations of the remaining array (nums[1:]).
   - For each permutation, insert nums[0] into every possible position.
   - Collect all permutations into the result.

Pseudo Code:
------------
function permute(nums):
    if nums is empty:
        return []
    if length(nums) == 1:
        return [[nums[0]]]

    result = []
    perms = permute(nums[1:])

    for each p in perms:
        for i = 0 to len(p):
            new_perm = insert nums[0] at index i in p
            result.append(new_perm)

    return result

Complexity Analysis:
--------------------
- Time Complexity: O(n * n!)
  * There are n! permutations.
  * Each permutation takes O(n) to build (due to insertion/copying).
- Space Complexity: O(n * n!)
  * Storing all n! permutations each of size n.
  * Recursion depth = O(n).
*/

package problems

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}

	res := [][]int{}
	perms := permute(nums[1:])

	for _, v := range perms {
		for i := 0; i <= len(v); i++ {
			// Insert nums[0] at position i
			c := make([]int, len(v)+1)
			copy(c[:i], v[:i])   // copy left part
			c[i] = nums[0]       // insert nums[0]
			copy(c[i+1:], v[i:]) // copy right part
			res = append(res, c)
		}
	}

	return res
}

func Permutations() {
	fmt.Print(permute([]int{1, 2, 3}))
}

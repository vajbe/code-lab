/*
LeetCode 128. Longest Consecutive Sequence
------------------------------------------

Problem:
Given an unsorted array of integers nums, return the length of the longest
consecutive elements sequence. The algorithm must run in O(n) time.

Example 1:
    Input:  nums = [100, 4, 200, 1, 3, 2]
    Output: 4
    Explanation: The longest consecutive sequence is [1,2,3,4].

Example 2:
    Input:  nums = [0,3,7,2,5,8,4,6,0,1]
    Output: 9
    Explanation: The longest consecutive sequence is [0,1,2,3,4,5,6,7,8].

Approach:
1. Insert all numbers into a HashSet for O(1) lookups.
2. For each unique number:
    - Check if it is the start of a sequence (num-1 not in set).
    - If yes, expand forward (num+1, num+2, ...) and count length.
    - Track the maximum length.
3. Return the maximum length.

Pseudo Code:
    function longestConsecutive(nums):
        if nums is empty: return 0
        set = new HashSet(nums)
        maxLen = 0
        for num in set:
            if num-1 not in set:
                curr = num
                length = 0
                while curr in set:
                    curr += 1
                    length += 1
                maxLen = max(maxLen, length)
        return maxLen

Time Complexity:
    - O(n), since each element is processed at most once.
Space Complexity:
    - O(n), for the HashSet.

Similar Problems:
    - LeetCode 300. Longest Increasing Subsequence (DP, O(n log n))
    - LeetCode 674. Longest Continuous Increasing Subsequence
    - LeetCode 164. Maximum Gap (bucket sort / radix sort)
    - LeetCode 287. Find the Duplicate Number (set/cycle detection)
*/

package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)

	for _, num := range nums {
		m[num] = true
	}

	maxSum := 0

	for num := range m {
		currSum := 0
		if !m[num-1] {
			curr := num
			for m[curr] {
				curr++
				currSum++
			}
			if currSum > maxSum {
				maxSum = currSum
			}
		}
	}
	return maxSum
}

func LongestCon() {
	fmt.Println("Longest is : ", longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println("Longest is : ", longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println("Longest is : ", longestConsecutive([]int{0, -1}))
}

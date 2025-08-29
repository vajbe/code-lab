// https://leetcode.com/problems/minimum-size-subarray-sum/submissions/1751928764/?envType=study-plan-v2&envId=top-interview-150// ...existing code...

/*
Problem Description:
Given an array of positive integers nums and a positive integer target,
find the minimal length of a contiguous subarray of which the sum is greater than or equal to target.
If there is no such subarray, return 0.

Pseudocode:
function minSubArrayLen(target, nums):

	n = length of nums
	min_len = n + 1
	curr_sum = 0
	left = 0

	for right from 0 to n-1:
	    curr_sum += nums[right]
	    while curr_sum >= target:
	        min_len = min(min_len, right - left + 1)
	        curr_sum -= nums[left]
	        left += 1

	if min_len > n:
	    return 0
	return min_len

Test Cases:
minSubArrayLen(7, [2, 3, 1, 2, 4, 3])      // Output: 2
minSubArrayLen(4, [1, 4, 4])                // Output: 1
minSubArrayLen(11, [1, 1, 1, 1, 1, 1, 1, 1])// Output: 0
minSubArrayLen(15, [1,2,3,4,5])             // Output: 5
minSubArrayLen(3, [1,1])                    // Output: 0
*/
package problems

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {

	n := len(nums)

	min_len := n + 1

	curr_sum := 0
	left := 0

	for right := 0; right < n; right++ {

		curr_sum += nums[right]

		for curr_sum >= target {
			min_len = min(min_len, right-left+1)
			curr_sum -= nums[left]
			left++
		}

	}

	if min_len > n {
		return 0
	}

	return min_len

}

func MinSub() {
	fmt.Println("\n", minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println("\n", minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println("\n", minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}

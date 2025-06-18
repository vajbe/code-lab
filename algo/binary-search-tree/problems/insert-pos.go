// LeetCode Problem: Search Insert Position
// https://leetcode.com/problems/search-insert-position
//
// Given a sorted array of distinct integers and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.
//
// Example 1:
// Input: nums = [1,3,5,6], target = 5
// Output: 2
//
// Example 2:
// Input: nums = [1,3,5,6], target = 2
// Output: 1
//
// Example 3:
// Input: nums = [1,3,5,6], target = 7
// Output: 4
//
// Constraints:
// - 1 <= nums.length <= 10^4
// - -10^4 <= nums[i] <= 10^4
// - nums contains distinct values sorted in ascending order.
// - -10^4 <= target <= 10^4

package problems

import "fmt"

func findPos(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	var result int
	for start <= end {
		mid := (start + end) / 2
		if target > nums[mid] {
			result = mid + 1
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
			result = mid
		} else {
			result = mid
			return result
		}
	}
	return result
}

func searchInsert(nums []int, target int) int {
	return findPos(nums, target)
}

func SearchInsertPos() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}

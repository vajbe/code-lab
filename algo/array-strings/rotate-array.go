// https://leetcode.com/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150

package main

import "fmt"

func reverseHelper(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotate(nums []int, k int) {
	l := len(nums)
	k %= l
	reverseHelper(nums, 0, l-1)
	reverseHelper(nums, 0, k-1)
	reverseHelper(nums, k, l-1)
	fmt.Print(nums)
}

func RotateArray() {
	// rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	// rotate([]int{-2}, 3)
	rotate([]int{1, 2}, 2)
}

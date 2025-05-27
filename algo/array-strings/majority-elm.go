// https://leetcode.com/problems/majority-element/description/?envType=study-plan-v2&envId=top-interview-150
package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	currNum := nums[0]
	currMax := 1
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if currNum == nums[i] {
			currMax++
		} else {
			currNum = nums[i]
		}
		if currMax > len(nums)/2 {
			return currNum
		}
	}

	return maxNum
}

/*
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
} */

func FindMajority() {
	num := majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	// num := majorityElement([]int{3, 2, 3})
	fmt.Print(num)
}

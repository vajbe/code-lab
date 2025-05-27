// https://leetcode.com/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150
package main

import "fmt"

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

func RemoveDups() {
	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nums := []int{1, 1, 2}
	l := removeDuplicates(nums)
	fmt.Print("Length is ", l, " ", nums)
}

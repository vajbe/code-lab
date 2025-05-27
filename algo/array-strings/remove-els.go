// https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
package main

import "fmt"

func removeElement(nums []int, val int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	fmt.Print(nums)

	return len(nums)
}

func RemoveEls() {
	removeElement([]int{3, 2, 2, 3}, 3)
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description/?envType=study-plan-v2&envId=top-interview-150
package main

import "fmt"

func removeDuplicate(nums []int) int {
	hashMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if _, exists := hashMap[num]; !exists {
			hashMap[num] = 0
		}
		fmt.Print("\nUpcoming length is ", hashMap[num]+1)
		if hashMap[num]+1 > 2 {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			hashMap[num]++
			fmt.Print("\tIncremeaned ", num, " -> ", hashMap[num])
		}
	}
	fmt.Print("\nBefore return \t", nums)
	return len(nums)
}

func RemoveDupsSorted() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	l := removeDuplicate(nums)
	fmt.Print("\n", l, " ", nums)
}

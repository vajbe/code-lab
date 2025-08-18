package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		key := nums[i]
		rem := target - key

		if _, ok := m[key]; !ok {
			m[key] = i
		}

		if _, ok := m[rem]; ok && m[rem] != i {
			if rem+key == target {
				return []int{m[rem], i}
			}
		}

	}
	return []int{}
}

func TwoSumm() {
	fmt.Print("Two sums ", twoSum(([]int{2, 7, 11, 15}), 9))
	fmt.Print("Two sums ", twoSum(([]int{3, 2, 4}), 6))
	fmt.Print("Two sums ", twoSum(([]int{3, 3}), 6))
}

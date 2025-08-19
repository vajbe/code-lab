package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if prevIdx, ok := m[nums[i]]; ok {
			if i-prevIdx <= k {
				return true
			}
		}
		m[nums[i]] = i
	}
	return false
}

func ContainsDup() {
	fmt.Println("Contains dup: ", containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println("Contains dup: ", containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println("Contains dup: ", containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

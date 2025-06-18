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

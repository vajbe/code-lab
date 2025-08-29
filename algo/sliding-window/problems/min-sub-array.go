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

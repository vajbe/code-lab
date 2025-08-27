package problems

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates for i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})

				// Skip duplicates for j
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				// Skip duplicates for k
				for j < k && nums[k] == nums[k-1] {
					k--
				}

				j++
				k--
			} else if sum < 0 {
				j++ // need larger value
			} else {
				k-- // need smaller value
			}
		}
	}

	return res
}

func Three() {
	fmt.Print("\n Three Sum: ", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Print("\n Three Sum: ", threeSum([]int{0, 0, 0}))
	fmt.Print("\n Three Sum: ", threeSum([]int{0, 1, 1}))
	fmt.Print("\n Three Sum: ", threeSum([]int{0, 0, 0, 0}))
}

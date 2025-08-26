package problems

import "fmt"

func twoSum(numbers []int, target int) []int {

	i, j := 0, 1

	for j < len(numbers) {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		} else if numbers[i]+numbers[j] > target || j == len(numbers)-1 {
			i++
			j = i + 1
		} else {
			j++
		}
	}

	return []int{i, j}
}

func TwoSum() {
	fmt.Print(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Print(twoSum([]int{2, 3, 4}, 6))
	fmt.Print(twoSum([]int{-1, 0}, -1))
	fmt.Print(twoSum([]int{5, 25, 75}, 100))
}

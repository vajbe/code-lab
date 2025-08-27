package problems

func maxArea(height []int) int {

	max_area := 0
	left := 0
	right := len(height) - 1

	for left < right {
		width := right - left
		h := min(height[left], height[right])

		area := width * h

		max_area = max(area, max_area)

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return max_area
}

func MaxAreaContainer() {
	maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
}

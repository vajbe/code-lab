package problems

import "fmt"

func lengthOfLongestSubstring(s string) int {

	m := make(map[byte]bool)
	left := 0
	max_len := 0
	curr_max := 0
	for right := 0; right < len(s); right++ {
		ch := s[right]
		if m[ch] {
			m = map[byte]bool{}
			left++
			right = left

			curr_max = 0
		}
		m[s[right]] = true
		curr_max++
		max_len = max(max_len, curr_max)
	}

	return max_len
}

func LongSubstring() {
	fmt.Print("\nLongest substring : ", lengthOfLongestSubstring("abcabcbb"))
	fmt.Print("\nLongest substring : ", lengthOfLongestSubstring("dvdf"))
}

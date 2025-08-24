package problems

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	i := 0
	j := 0
	for j < len(t) {
		// fmt.Println(i, j, len(s))
		if i < len(s) && s[i] == t[j] {
			// fmt.Println(i, len(s))
			i++
			j++
			continue

		}

		j++
	}

	return i == len(s)
}

func Subsequence() {
	fmt.Println("\nSubsequence: ", isSubsequence("abc", "ahbgdc"))
	// fmt.Print("\nSubsequence: ", isSubsequence("aec", "ahbgdc"))
	// fmt.Print("\nSubsequence: ", isSubsequence("", "ahbgdc"))
	fmt.Println("\nSubsequence: ", isSubsequence("bb", "ahbgdc"))
}

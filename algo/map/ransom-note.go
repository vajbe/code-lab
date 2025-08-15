package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {

	m := make(map[rune]int)

	for _, s := range magazine {
		m[s] = m[s] + 1
	}

	for _, s := range ransomNote {
		if m[s] == 0 {
			return false
		} else {
			m[s] = m[s] - 1
		}
	}

	return true
}

func RansomNote() {
	fmt.Print("\n Can be formed: ", canConstruct("a", "b"))
	fmt.Print("\n Can be formed: ", canConstruct("aa", "ab"))
	fmt.Print("\n Can be formed: ", canConstruct("aa", "aab"))
}

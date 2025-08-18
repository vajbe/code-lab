package main

import "fmt"

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	m := make(map[string]int)

	for _, chr := range s {
		str := string(chr)

		if _, ok := m[str]; !ok {
			m[str] = 0
		}

		m[str] = m[str] + 1
	}

	for _, chr := range t {
		str := string(chr)

		m[str] = m[str] - 1
	}

	for _, val := range m {
		if val != 0 {
			return false
		}

	}

	return true
}

func IsAna() {
	fmt.Println("Is anagram: ", isAnagram("anagram", "nagaram"))
	fmt.Println("Is anagram: ", isAnagram("rat", "cat"))
}

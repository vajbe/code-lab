package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	m := make(map[string]string)
	alreadyMapped := make(map[string]bool)
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		str := string(pattern[i])
		word := words[i]

		if _, ok := m[str]; !ok {
			if alreadyMapped[word] {
				return false
			} else {
				m[str] = word
				alreadyMapped[word] = true
			}

		} else if m[str] != word {
			return false
		}
	}
	return true
}

func WrdPattern() {
	fmt.Println("Is pattern: ", wordPattern("abba", "dog cat cat dog"))
	fmt.Println("Is pattern: ", wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println("Is pattern: ", wordPattern("abba", "dog dog dog dog"))
}

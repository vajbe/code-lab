package main

import (
	"fmt"
	"sort"
)

func sortString(s string) string {
	// Convert to rune slice to handle Unicode properly
	runes := []rune(s)

	// Sort runes in place
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// Convert back to string
	return string(runes)
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		s := sortString(str)

		if _, ok := m[s]; !ok {
			m[s] = []string{}
		}
		m[s] = append(m[s], str)
	}
	res := [][]string{}
	for _, value := range m {
		res = append(res, value)
	}

	return res
}

func GroupAna() {
	fmt.Print(groupAnagrams([]string{"eat", "tan"}))
}

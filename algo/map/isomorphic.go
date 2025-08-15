package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	pair := make(map[string]string)
	alreadyMapped := make(map[string]bool)

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		st1 := string(s[i])
		st2 := string(t[i])
		if _, ok := pair[st1]; !ok {
			if !alreadyMapped[st2] {
				pair[st1] = st2
				alreadyMapped[st2] = true
			} else {
				return false
			}

		} else if pair[st1] != st2 {
			return false
		}
	}

	fmt.Println("Map is ", pair)

	return true
}

func Isomorphic() {
	fmt.Println("Is isomorphic ", isIsomorphic("egge", "abba"))
	fmt.Println("Is isomorphic ", isIsomorphic("badc", "baba"))
}

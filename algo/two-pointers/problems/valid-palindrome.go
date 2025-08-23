package problems

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	rs := []rune(s)
	i, j := 0, len(rs)-1

	isAlphaNum := func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsDigit(r)
	}

	for i < j {
		if !isAlphaNum(rs[i]) {
			i++
			continue
		}
		if !isAlphaNum(rs[j]) {
			j--
			continue
		}
		if unicode.ToLower(rs[i]) != unicode.ToLower(rs[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func Palindrome() {
	// fmt.Println("Is palindrome: ", isPalindrome("Vivek"))
	fmt.Println("Is palindrome: ", isPalindrome("0P"))
	// fmt.Println("Is palindrome: ", isPalindrome("A man, a plan, a canal: Panama"))
}

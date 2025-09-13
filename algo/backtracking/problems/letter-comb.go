package problems

import "fmt"

func helperFunction(char string, m map[string][]string, digits string) []string {
	if len(digits) == 0 {
		return []string{char}
	}
	res := []string{}
	letters := m[string(digits[0])]
	for _, c := range letters {
		res = append(res, helperFunction(char+string(c), m, digits[1:])...)
	}
	return res
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	m := make(map[string][]string)
	m["2"] = []string{"a", "b", "c"}
	m["3"] = []string{"d", "e", "f"}
	m["4"] = []string{"g", "h", "i"}
	m["5"] = []string{"j", "k", "l"}
	m["6"] = []string{"m", "n", "o"}
	m["7"] = []string{"p", "q", "r", "s"}
	m["8"] = []string{"t", "u", "v"}
	m["9"] = []string{"w", "x", "y", "z"}
	res := helperFunction("", m, digits)
	return res
}

func LetterComb() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("234"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
}

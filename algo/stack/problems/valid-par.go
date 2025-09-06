package problems

import "fmt"

func isValid(s string) bool {
	stack := []string{}
	m := make(map[string]string)

	m["("] = ")"
	m["{"] = "}"
	m["["] = "]"

	for _, char := range s {
		if string(char) == "(" || string(char) == "{" || string(char) == "[" {
			stack = append(stack, string(char))
		} else {
			pop := "-"
			if len(stack) > 0 {
				pop = m[stack[len(stack)-1]]
				stack = stack[:len(stack)-1]
			}

			if string(char) != pop {
				return false
			}

		}
	}
	return len(stack) == 0
}

func ValidPars() {
	/* 	fmt.Println("Is valid : ", isValid("()"))
	   	fmt.Println("Is valid : ", isValid("([)]"))
	   	fmt.Println("Is valid : ", isValid("()[]{}"))
	   	fmt.Println("Is valid : ", isValid("(]"))
	   	fmt.Println("Is valid : ", isValid("([])")) */
	fmt.Println("Is valid : ", isValid("]"))
}

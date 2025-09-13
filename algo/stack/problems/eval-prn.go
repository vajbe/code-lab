package problems

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, char := range tokens {
		s := string(char)

		if num, err := strconv.Atoi(s); err == nil {
			stack = append(stack, num)
		} else {
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			switch s {
			case "+":
				stack = append(stack, (a + b))
			case "-":
				stack = append(stack, (a - b))
			case "/":
				stack = append(stack, (a / b))
			case "*":
				stack = append(stack, (a * b))
			}
		}
	}

	return stack[0]
}

func EvalPRN() {
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

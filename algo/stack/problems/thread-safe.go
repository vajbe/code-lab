package problems

import "fmt"

type Stack struct {
	elems []int
}

func (stack *Stack) Peek() int {
	if len(stack.elems) == 0 {
		fmt.Print("Stack is empty ")
		return -1
	}
	top := stack.elems[len(stack.elems)-1]
	return top
}

func (stack *Stack) Pop() int {
	if len(stack.elems) == 0 {
		fmt.Print("Stack is empty ")
		return -1
	}
	top := stack.elems[len(stack.elems)-1]
	stack.elems = stack.elems[0 : len(stack.elems)-1]
	return top
}

func (stack *Stack) Push(value int) {
	stack.elems = append(stack.elems, value)
}

func ThreadSafeStack() {

	stack := Stack{}
	stack.Push(1)
	stack.Push(5)
	stack.Push(8)
	stack.Push(2)

	fmt.Printf("%+v\n", stack.elems)
	fmt.Printf("%d\n", stack.Pop())
	fmt.Printf("%d\n", stack.Pop())
	fmt.Printf("%d\n", stack.Pop())
	fmt.Printf("Top Elem %d\n", stack.Peek())
	fmt.Printf("%d\n", stack.Pop())
	fmt.Printf("%d\n", stack.Pop())
	fmt.Printf("%d\n", stack.Pop())
	fmt.Printf("Top Elem %d\n", stack.Peek())
}

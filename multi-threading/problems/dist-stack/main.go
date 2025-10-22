package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	s    []int
	lock sync.Mutex
}

func NewStack() *Stack {
	return &Stack{}
}

func main() {
	stack := NewStack()
	var wg sync.WaitGroup
	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go func(val int) {
			stack.Add(i * 10)
			defer wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Print("\nLength is ", len(stack.s))
}

func (stack *Stack) Add(num int) {
	stack.lock.Lock()
	stack.s = append(stack.s, num)
	stack.lock.Unlock()
}

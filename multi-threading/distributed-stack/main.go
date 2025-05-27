package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	elem []int
	mx   sync.Mutex
}

func (s *Stack) Push(elem int) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.elem = append(s.elem, elem)
}

func (s *Stack) Pop() int {
	if len(s.elem) == 0 {
		return -1
	}
	s.mx.Lock()
	defer s.mx.Unlock()
	num := s.elem[len(s.elem)-1]
	s.elem = s.elem[0 : len(s.elem)-1]
	return num
}

func (s *Stack) Print() {
	fmt.Print("\nElems are: [")
	for x := range s.elem {
		fmt.Printf(" %d", s.elem[x])
	}
	fmt.Print(" ]\n")
}

func main() {
	stack := Stack{}
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			stack.Push(i)
		}(i)
	}

	wg.Wait()

	stack.Print()
}

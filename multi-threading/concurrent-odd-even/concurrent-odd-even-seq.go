package main

import (
	"fmt"
	"sync"
)

var data = make(chan int32)
var wg sync.WaitGroup

func SequenceBasedOddEven() {
	wg.Add(2)
	go printOdd()
	go printEven()
	data <- 1
	wg.Wait()
}

func printOdd() {
	defer wg.Done()
	for d := range data {
		if d%2 != 0 {
			fmt.Print("\nG1-", d)
			if d >= 10 {
				close(data)
				return
			}
			data <- d + 1
		} else {
			data <- d
		}
	}
}

func printEven() {

	defer wg.Done()
	for d := range data {
		if d%2 == 0 {
			fmt.Print("\nG2-", d)
			if d >= 10 {
				close(data)
				return
			}
			data <- d + 1
		} else {
			data <- d
		}
	}
}

package main

import (
	"fmt"
	"time"
)

// Generator function
func GeneratorFun(num int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < num; i++ {
			out <- i
			time.Sleep(1 * time.Second) // Simulate some work
		}
		close(out)
	}()
	return out
}

func main() {
	fmt.Printf("Starting main thread:\n")

	// Run generator loop in a separate goroutine
	go func() {
		for i := range GeneratorFun(5) {
			fmt.Printf("Generated value: %v \n", i)
		}
		fmt.Printf("Generator finished.\n")
	}()

	// Main function continues executing other tasks
	for i := 0; i < 5; i++ {
		fmt.Printf("Main thread doing other work: %v \n", i)
		time.Sleep(500 * time.Millisecond) // Simulate work in main thread
	}

	fmt.Printf("Main thread finished other work.\n")

	// Adding a sleep to ensure the goroutine completes before main exits
	time.Sleep(5 * time.Second)
}

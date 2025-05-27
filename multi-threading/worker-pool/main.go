package main

import (
	"fmt"
	"time"
)

var maxNum = 1000000

func verifyEven(num int) bool {
	fmt.Print(num)
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	startTime := time.Now()
	fmt.Printf("Starting %s\n", startTime)
	for i := 1; i <= maxNum; i++ {
		verifyEven(i)
	}
	fmt.Printf("\nStarted at %s eneded at %s\n", startTime, time.Now())
	fmt.Printf("Elapsed Time %f \n", time.Now().Sub(startTime).Seconds())
}

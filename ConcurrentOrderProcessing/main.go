package main

import (
	"fmt"
	"time"
)

func orderScheduler(orderId int) {
	fmt.Printf("\nProcessing order %d\t", orderId)
	time.Sleep(100 * time.Microsecond)
	fmt.Printf("\nDone with order %d", orderId)
}

func main() {
	startTime := time.Now()

	for i := 0; i < 50; i++ {
		go orderScheduler(i)
	}
	fmt.Printf("\nStarted at %s eneded at %s\n", startTime, time.Now())
	fmt.Printf("\nElapsed Time %s\n", time.Since(startTime))

}

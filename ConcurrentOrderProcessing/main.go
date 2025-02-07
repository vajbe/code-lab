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

func worker(jobId int, jobs <-chan int, result <-chan int) {

}

func main() {
	numJobs := 5
	startTime := time.Now()
	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)
	for i := 0; i < 50; i++ {
		go worker(i, jobs, result)
	}
	fmt.Printf("\nStarted at %s eneded at %s\n", startTime, time.Now())
	fmt.Printf("\nElapsed Time %s\n", time.Since(startTime))

}

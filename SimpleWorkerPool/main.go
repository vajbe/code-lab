package main

import (
	"fmt"
	"time"
)

func performJob(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %v starting job %v \n", id, j)
		time.Sleep(3 * time.Second)
		fmt.Printf("Worker %v finished job %v \n", id, j)
		results <- (j * 2)
	}
}

func main() {
	numJobs := 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 1; i <= 3; i++ {
		go performJob(i, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var totalUnfairPrimeNumber int32 = 0

func checkUnfairPrime(num int) {
	if num&1 == 0 {
		return
	}

	for i := 3; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return
		}
	}
	atomic.AddInt32(&totalUnfairPrimeNumber, 1)
}

func doBatch(start int, end int, wg *sync.WaitGroup, worker int) {
	defer wg.Done()
	startTime := time.Now()
	for i := start; i <= end; i++ {
		checkUnfairPrime(i)
	}
	fmt.Print("\nWorker: ", worker, " \tStart: ", start, " \tEnd:", end, " \tElapsed: ", time.Since(startTime))
}

func UnfairCheck() {
	var wg sync.WaitGroup
	batchSize := int(MAX_INT) / CONCURRENCY
	start := 2
	end := batchSize
	for i := 0; i < CONCURRENCY; i++ {
		wg.Add(1)
		go doBatch(start+1, end, &wg, i)
		start = end
		end = start + batchSize
	}
	wg.Wait()
	fmt.Print("\nTotal Prime Numbers ", totalUnfairPrimeNumber+1)
}

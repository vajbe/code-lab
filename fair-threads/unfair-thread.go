package main

import (
	"fmt"
	"math"
	"sync"
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
	totalUnfairPrimeNumber++
}

func doBatch(start int64, end int64, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print(start, end)
	checkUnfairPrime(int(start))
}

func UnfairCheck() {
	var wg sync.WaitGroup
	startTime := time.Now()
	batchSize := int(MAX_INT) / CONCURRENCY
	start := 2
	end := batchSize
	for i := 0; i < CONCURRENCY; i++ {
		fmt.Print("\nStart ", start+1, " End ", end, " i ", i)
		go doBatch(int64(start+1), int64(end), &wg)
		start = end
		end = start + batchSize
	}

	fmt.Print("Elapsed:", time.Since(startTime))

	wg.Wait()
}

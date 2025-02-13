package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var totalFairPrimeNumbers int32 = 0
var currentNum int32 = 0

func checkFairPrime(num int) {
	if num&1 == 0 {
		return
	}

	for i := 3; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return
		}
	}

	atomic.AddInt32(&totalFairPrimeNumbers, 1)
}

func doWork(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	startTime := time.Now()
	for {
		x := atomic.AddInt32(&currentNum, 1)
		if x > int32(MAX_INT) {
			break
		}
		checkFairPrime(int(x))
	}

	fmt.Print("\n Thread: ", id, " Elapsed: ", time.Since(startTime))
}

func FairCheck() {
	fmt.Print("\nStarting At: ", time.Now())
	var wg sync.WaitGroup

	for i := 0; i <= CONCURRENCY; i++ {
		wg.Add(1)
		go doWork(&wg, i)
	}

	wg.Wait()

	fmt.Print("\nEnding At: ", time.Now())

}

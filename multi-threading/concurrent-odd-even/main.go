package main

import (
	"sync"
	"sync/atomic"
)

const MAX_NUM int64 = 89237849009

var CurrentNum int64
var EvenCount int64 // Changed to int64 for consistency with MAX_NUM

func IsEven(num int64) bool {
	if num&1 == 0 {
		return true
	}
	return false
}

func Task(start int64, step int, wg *sync.WaitGroup) {
	defer wg.Done()
	localCount := 0
	for x := start; x < MAX_NUM; x += int64(step) {
		if IsEven(x) {
			localCount += 1
		}
	}
	atomic.AddInt64(&EvenCount, int64(localCount))
}

func main() {
	/* startTime := time.Now()

	numWorkers := 2
	wg := sync.WaitGroup{}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Task(int64(i), numWorkers, &wg)
	}

	wg.Wait()

	fmt.Printf("\nTime Elapsed: %fs", time.Since(startTime).Seconds())
	fmt.Printf("\nTotal Even: %d, Total Odd: %d\n", EvenCount, MAX_NUM-EvenCount) */
	HandleOddEven()
}

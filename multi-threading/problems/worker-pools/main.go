package main

import (
	"fmt"
	"runtime"
)

const MAX_NUM = 5000000

func main() {
	workers := runtime.NumCPU()
	chn := make(chan int)
	total := 0
	batchSize := MAX_NUM / workers
	start := 1
	for i := 0; i < workers; i++ {
		end := start + batchSize - 1
		go isOdd(start, end, chn)
		start = end + 1
	}
	done := 0
	for c := range chn {
		total += c
		done++
		if done == workers {
			// close(chn)
			break
		}
	}

	fmt.Print("\n Total ", total)
	close(chn)
	/* go func() {
		close(chn)
	}() */
}

func isOdd(start int, end int, chn chan<- int) {
	count := 0
	for i := start; i <= end; i++ {
		if i%2 == 0 {
		} else {
			count++
		}
	}
	chn <- count
}

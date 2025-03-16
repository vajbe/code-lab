package main

import (
	"sync"
	"time"
)

const MAX_NUM = 89237849009

var EvenCount int

func IsEven(num int) bool {
	return num%2 == 0
}

func Task(start int, end int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	totalEven := 0
	for i := start; i <= end; i++ {
		if IsEven(i) {
			totalEven += 1
		}
	}
	ch <- totalEven
}

func main() {
	/*
		PART 1
		ch := make(chan int)
		wg := sync.WaitGroup{}
		timeNow := time.Now()
		wg.Add(1)
		go Task(1, MAX_NUM/2, &wg, ch)
		wg.Add(1)
		go Task(MAX_NUM/2+1, MAX_NUM, &wg, ch)
		b, c := <-ch, <-ch
		wg.Wait()
		fmt.Printf("\nTime Elapsed: %fs", time.Now().Sub(timeNow).Seconds())
		fmt.Printf("\nTotal Even %d, Total Odd %d", b+c, MAX_NUM-(b+c)) */

	ch := make(chan int)
	wg := sync.WaitGroup{}
	timeN := time.Now()

}

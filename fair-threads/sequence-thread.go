package main

import (
	"fmt"
	"math"
	"time"
)

var MAX_INT int64 = 100000000
var totalPrimeNumber int32 = 0

func checkPrime(num int) {
	if num&1 == 0 {
		return
	}

	for i := 3; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return
		}
	}
	totalPrimeNumber++
}

func SequenceCheck() {
	fmt.Print("\nStarting...")
	start := time.Now()

	for i := 3; i <= int(MAX_INT); i++ {
		checkPrime(i)
	}

	fmt.Print("\nFinding for ", MAX_INT, " Started around ", start, " Found ", totalPrimeNumber+1, " Elapsed ", time.Since(start))

}

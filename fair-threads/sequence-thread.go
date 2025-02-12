package main

import (
	"fmt"
	"math"
	"time"
)

var totalSeqPrimeNumber int32 = 0

func checkSeqPrime(num int) {
	if num&1 == 0 {
		return
	}

	for i := 3; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return
		}
	}
	totalSeqPrimeNumber++
}

func SequenceCheck() {
	fmt.Print("\nStarting...")
	start := time.Now()

	for i := 3; i <= int(MAX_INT); i++ {
		checkSeqPrime(i)
	}

	fmt.Print("\nFinding for ", MAX_INT, " Started around ", start, " Found ", totalSeqPrimeNumber+1, " Elapsed ", time.Since(start))

}

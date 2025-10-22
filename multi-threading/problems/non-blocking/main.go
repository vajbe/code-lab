package main

import (
	"fmt"
	"sync"
)

func main() {
	chn := make(chan string, 1)
	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(1)
	go waitForData(chn, &wg)
	wg.Wait()
}

func waitForData(chn chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(chn)

	chn <- "\nBooka"
	fmt.Print("\nContinue")
}

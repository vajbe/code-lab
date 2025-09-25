package main

import "fmt"

const MAX_SEQ = 18

var result chan bool

func printOdd_(oddChan <-chan int, evenChan chan<- int) {
	for {
		num := <-oddChan

		fmt.Println("Odd num: ", num)
		num++
		if num > MAX_SEQ {
			result <- true
			return
		}
		evenChan <- num
	}
}

func printEven_(oddChan chan<- int, evenChan <-chan int) {
	for {
		num := <-evenChan

		fmt.Println("Even num: ", num)
		num++
		if num > MAX_SEQ {
			result <- true
			return
		}
		oddChan <- num
	}
}

func HandleOddEven() {
	result = make(chan bool, 1)
	oddChan := make(chan int)
	evenChan := make(chan int)
	go printOdd_(oddChan, evenChan)
	go printEven_(oddChan, evenChan)

	oddChan <- 1

	<-result

	fmt.Println("Done with processing..!!!")
	close(result)
	close(oddChan)
	close(evenChan)

}

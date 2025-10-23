package main

import (
	"fmt"
	"runtime"
	"time"
)

var (
	workers  int
	queue    chan string
	doneChan chan bool
)

func main() {
	workers = runtime.NumCPU()
	queue = make(chan string)
	doneChan = make(chan bool)

	go Producer()
	go Consumer()
	go Vitals()

	<-doneChan
	close(doneChan)
}

func Producer() {

	tick := time.NewTicker(2 * time.Second)
	timeout := time.After(10 * time.Second)

	for {
		select {
		case <-tick.C:
			queue <- "Random string: " + time.Now().String()
		case <-timeout:
			close(queue)
			doneChan <- true
			return
		}

	}

}

func Consumer() {

	for str := range queue {
		fmt.Print("\n", str)
	}

}

func Vitals() {

	for {
		time.Sleep(5 * time.Second)
		fmt.Print("\t ", runtime.NumGoroutine())
	}
}

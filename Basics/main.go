package main

import (
	"fmt"
	"time"
)

func greet(msg string, doneChan chan bool) {
	fmt.Println("Hello ", msg)
	doneChan <- true
}

func slowGreet(msg string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hi I am slower, ", msg)
	doneChan <- true
}

func main() {
	done := make(chan bool)
	go greet("Vivek", done)
	go slowGreet("Agastya", done)
	go slowGreet("Shambhavi", done)
	<-done
	<-done
	<-done
}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	Id     int
	Item   string
	Amount int
}

const (
	numWorkers   = 10  // Number of concurrent workers
	numOrders    = 50  // Total orders to process
	initialStock = 100 // Initial stock quantity
)

var lock sync.Mutex
var stockLevels = map[string]int{
	"Laptop": initialStock,
}

var wg sync.WaitGroup

func ProcessOrder(order Order) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	lock.Lock()
	if stockLevels[order.Item] >= order.Amount {
		stockLevels[order.Item] = stockLevels[order.Item] - order.Amount
		fmt.Printf("\n Processed order: %d Stocks required:%d Stocks left: %d", order.Id, order.Amount, stockLevels[order.Item])
	} else {
		fmt.Printf("\n Failed to process order %d no stock available", order.Id)
	}
	lock.Unlock()
}

func worker(orderQueue <-chan Order) {
	for order := range orderQueue {
		ProcessOrder(order)
	}
}

func main() {
	numJobs := 5
	orderQueue := make(chan Order, numJobs)

	// Start worker pool
	for i := 0; i < numWorkers; i++ {
		go worker(orderQueue)
	}

	//Generate and Queue order

	for j := 0; j <= numOrders; j++ {
		order := Order{
			Id:     j,
			Item:   "Laptop",
			Amount: rand.Intn(5) + 1,
		}
		wg.Add(1)
		orderQueue <- order
	}
	close(orderQueue)
	wg.Wait()
	fmt.Println("\n🏁 All orders processed!")

}

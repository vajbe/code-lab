package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type ConcurrentQueue struct {
	queue []int32
	muE   sync.Mutex
}

var wg sync.WaitGroup

func (q *ConcurrentQueue) Enqueue(item int32) {
	q.muE.Lock()
	defer q.muE.Unlock()
	q.queue = append(q.queue, item)
}

func (q *ConcurrentQueue) Dequeue() int32 {
	if len(q.queue) == 0 {
		panic("Oops queue is empty")
	}
	q.muE.Lock()
	defer q.muE.Unlock()
	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func (q *ConcurrentQueue) Size() int {
	q.muE.Lock()
	defer q.muE.Unlock()
	return len(q.queue)
}

func main() {
	q := ConcurrentQueue{
		queue: make([]int32, 0),
	}

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			q.Enqueue(rand.Int31())
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Print("\n Size: ", q.Size())
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			q.Dequeue()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Print("\n Size: ", q.Size())
}

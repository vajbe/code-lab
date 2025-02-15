package main

import (
	"fmt"
	"math/rand"
)

type ConcurrentQueue struct {
	queue []int32
}

func (q *ConcurrentQueue) Enqueue(item int32) {
	q.queue = append(q.queue, item)
}

func (q *ConcurrentQueue) Dequeue() int32 {
	if len(q.queue) == 0 {
		panic("Oops queue is empty")
	}
	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func (q *ConcurrentQueue) Size() int {
	return len(q.queue)
}

func main() {
	q := ConcurrentQueue{
		queue: make([]int32, 0),
	}

	for i := 0; i < 1000000; i++ {
		q.Enqueue(rand.Int31())
	}

	fmt.Print("\n Size: ", q.Size())

}

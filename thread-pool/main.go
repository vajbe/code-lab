package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

type Worker struct {
	id        int
	taskQueue chan Task
	wg        *sync.WaitGroup
}

func StartWorker(id int, taskQueue chan Task, wg *sync.WaitGroup) {
	for task := range taskQueue {
		task()
		wg.Done()
	}
}

type Pool struct {
	taskQueue chan Task
	wg        sync.WaitGroup
}

func NewPool(numWorkers int, queueSize int) *Pool {
	pool := &Pool{
		taskQueue: make(chan Task),
	}

	for i := 0; i < numWorkers; i++ {
		go StartWorker(i, pool.taskQueue, &pool.wg)
	}
	return pool
}

func (p *Pool) Submit(task Task) {
	p.wg.Add(1)
	p.taskQueue <- task
}
func (p *Pool) Shutdown() {
	p.wg.Wait()
	close(p.taskQueue)
}

func main() {
	pool := NewPool(5, 10)
	for i := 1; i <= 20; i++ {
		jobId := i
		pool.Submit(func() {
			fmt.Print("\nTaskId: ", jobId)
			time.Sleep(1 * time.Second)
		})
	}

	pool.Shutdown()
	fmt.Print("\nAll Tasks Completed")

}

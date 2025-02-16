// Mimick basic deadlock condition
package main

import (
	"fmt"
	"sync"
	"time"
)

type Resource struct {
	id int
	mu sync.Mutex
}

func processOrder(id int, resource1 *Resource, resource2 *Resource, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Print("\nProcess: ", id, " For Resource: ", resource1.id)
	resource1.mu.Lock()
	fmt.Print("\nProcess: ", id, " Acquired Lock For Resource: ", resource1.id)

	time.Sleep(1 * time.Second)

	fmt.Print("\nProcess: ", id, " For Resource: ", resource2.id)
	resource2.mu.Lock()
	fmt.Print("\nProcess: ", id, " Acquired Lock For Resource: ", resource2.id)

	/* time.Sleep(1 * time.Second) */

	resource2.mu.Unlock()
	resource1.mu.Unlock()

}

func main() {

	var wg sync.WaitGroup

	R1 := &Resource{id: 1}
	R2 := &Resource{id: 2}

	wg.Add(2)
	go processOrder(1, R1, R2, &wg)
	go processOrder(2, R2, R1, &wg) // Process 2: R2 → R1 (Conflicting order!)

	wg.Wait()

	fmt.Print("\nFinished processing")
}

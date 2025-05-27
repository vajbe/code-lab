package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Resource struct {
	id int
	mu sync.Mutex
}

var MAX_TRANS int = 100

func processOrder(txnId int, res1 *Resource, res2 *Resource, wg *sync.WaitGroup) {
	defer wg.Done()

	res1.mu.Lock()
	fmt.Print("\nAcquired lock on ", res1.id)
	res2.mu.Lock()
	fmt.Print("\nAcquired lock on ", res2.id)

	time.Sleep(2 * time.Second)

	res2.mu.Unlock()
	res1.mu.Unlock()
	time.Sleep(2 * time.Second)
	fmt.Print("\nFinished Processing ", txnId)
}

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()
	res1 := Resource{id: 1}
	res2 := Resource{id: 2}
	res3 := Resource{id: 3}

	resources := []*Resource{&res1, &res2, &res3}

	min, max := 0, 2

	for i := 0; i <= MAX_TRANS; i++ {

		rand1 := rand.Intn(max-min+1) + min
		rand2 := rand.Intn(max-min+1) + min
		r1 := resources[rand1]
		r2 := resources[rand2]
		txn := rand.Int()
		if r1.id == r2.id {
			continue
		}

		/* This part is about making sure to follow the total order processing of shared resources
		in case if r2 is has higher value then make sure to acquire it only after r1
		thus by swapping the resource order
		If removed it will go into deadlock
		/*/
		if r1.id < r2.id {
			temp := r2
			r2 = r1
			r1 = temp
		}

		wg.Add(1)
		fmt.Print("\nProcessing transaction: ", txn)
		go processOrder(txn, r1, r2, &wg)
	}

	/* fmt.Print("All transactions performed...!") */

}

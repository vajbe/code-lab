package main

import (
	"fmt"
	"sync"
)

type SingleSyncOnce struct {
}

var syncOnce sync.Once

var singleSyncOnce *SingleSyncOnce

func getInstanceSyncOnce() *SingleSyncOnce {
	if singleSyncOnce == nil {
		syncOnce.Do(
			func() {
				fmt.Println("Creating sync once instance")
				singleSyncOnce = &SingleSyncOnce{}
			})
	} else {
		fmt.Println("Sync once instance already created.")
	}
	return singleSyncOnce
}

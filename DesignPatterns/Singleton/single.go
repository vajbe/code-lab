package main

import (
	"fmt"
	"sync"
)

type Single struct {
}

var lock = &sync.Mutex{}
var singleInstance *Single

func getInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating a single instance...")
			singleInstance = &Single{}
		} else {
			fmt.Println("Singleton already exists...")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleInstance
}

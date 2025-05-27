package main

import (
	"context"
	"fmt"
	"time"
)

func WithCancellation() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	go func() {
		<-ctx.Done()
		fmt.Println("Context cancelled...", ctx.Err())
	}()
	time.Sleep(2 * time.Second)
	cancel()
	fmt.Scanln()
}

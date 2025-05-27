package main

import (
	"context"
	"fmt"
	"time"
)

func WithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Context timedout")
	}
}

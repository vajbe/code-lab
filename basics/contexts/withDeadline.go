package main

import (
	"context"
	"fmt"
	"time"
)

func WithDeadline() {
	deadline := time.Now().Add(12 * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	select {
	case <-time.After(10 * time.Second):
		fmt.Print("Task completed")
	case <-ctx.Done():
		fmt.Print("Cancelled due to timeout")
	}

}

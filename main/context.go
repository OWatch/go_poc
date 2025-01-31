package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	testContext()
}

func testContext() {
	ctx, cancal := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancal()

	go handle(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

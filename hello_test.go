package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	go func() {
		fmt.Println("I am goroutine")
	}()

	time.Sleep(100)
	fmt.Println("I am main")
}

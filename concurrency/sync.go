package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("I am goroutine")
	}()

	time.Sleep(100)
	fmt.Println("I am main")
}

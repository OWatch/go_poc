package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	chan1 := make(chan int)

	go func() {
		time.Sleep(time.Second * 1)
		chan1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 1)
		chan1 <- 2
	}()

	go func() {
		time.Sleep(time.Second * 1)
		chan1 <- 3
	}()

	fmt.Sprint()
}

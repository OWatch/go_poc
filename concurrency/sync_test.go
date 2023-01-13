package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	mu sync.Mutex
)

func TestMutex(t *testing.T) {
	var cnt = new(int)
	*cnt = 0
	var i = 0
	for ; i < 10; i++ {
		go func(cnt *int) {
			mu.Lock()
			defer mu.Unlock()
			*cnt++
			fmt.Println(*cnt)
			if *cnt == 10 {
				return
			}
		}(cnt)

		time.Sleep(1000 * time.Microsecond)
	}

	fmt.Println("end")

}

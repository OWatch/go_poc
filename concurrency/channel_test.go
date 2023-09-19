package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestChannel(t *testing.T) {
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	ch := make(chan []int, 1)
	ch <- arr2

	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			arr1 = append(arr1, i)

			tmpArr := <-ch
			tmpArr = append(tmpArr, i)
			ch <- tmpArr
		}()
	}
	wg.Wait()

	fmt.Printf("arr1:%d", len(arr1))
	fmt.Println()

	fmt.Printf("arr2:%d", len(<-ch))
	fmt.Println()

}

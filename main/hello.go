package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// ans := calSum(3, []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7})
	// fmt.Println(ans)

	// var foo = func(i int) (int, error) {
	// 	return i, errors.New("this is a err\n")
	// }
	// ans, errs := batchExecuteTask(3, []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7}, foo)
	// fmt.Println(ans)
	// fmt.Println(errs)

	// towGoroutineCrossPrint()

	kGoroutineCrossPrint(3)
}

func calSum(k int, arr []int) int64 {
	L := len(arr)
	var sum int64 // CAS 乐观锁
	var wg sync.WaitGroup
	var index int64
	for i := 0; i < k; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				currentIndex := atomic.LoadInt64(&index)
				if currentIndex >= int64(L) {
					return
				}

				// atomic.AddInt64(&index, 1)
				// atomic.AddInt64(&sum, int64(arr[currentIndex]))
				if atomic.CompareAndSwapInt64(&index, currentIndex, currentIndex+1) {
					atomic.AddInt64(&sum, int64(arr[currentIndex]))
				}
			}
		}(i)
	}
	wg.Wait()

	return sum
}

func batchExecuteTask(k int, taskList []int, fn func(int) (int, error)) ([]int, []error) {
	n := len(taskList)
	ans := make([]int, n)
	errs := make([]error, n)
	var index int32
	var wg sync.WaitGroup
	for i := 0; i < k; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			for {
				currentIndex := atomic.LoadInt32(&index)
				if currentIndex >= int32(n) {
					return
				}
				if atomic.CompareAndSwapInt32(&index, currentIndex, currentIndex+1) {
					result, err := fn(taskList[currentIndex])
					ans[currentIndex] = result
					errs[currentIndex] = err
				}

			}
		}(i)
	}
	wg.Wait()

	return ans, errs
}

func towGoroutineCrossPrint() {
	var ch = make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 1; i < 101; i++ {
			ch <- 1 // 核心是通过chan，同步2个协程，奇数先打印，然后chan阻塞住，等偶数协程消费打印，然后继续奇数线程可打印...
			if i%2 != 0 {
				fmt.Println(i)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 1; i < 101; i++ {
			<-ch
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()

	wg.Wait()
	fmt.Println("end")
}

func kGoroutineCrossPrint(k int) {
	var chs = make([]chan int, k)
	var exitChan = make(chan int, 1)
	for i := range chs {
		chs[i] = make(chan int, 1) // 初始化channel数组
	}

	var j = 0 // 指引协程轮询num
	var num = 1
	for i := 0; i < k; i++ {
		go func(i int) {
			j = i
			for {
				if num > 100 {
					exitChan <- 1 // 退出协程
					break
				}

				<-chs[j]
				fmt.Println(num)
				num++

				j++
				if j == k {
					j = 0
				}
				chs[j] <- 1
			}
		}(i)
	}
	chs[0] <- 1

	select {
	case <-exitChan:
		fmt.Println("end")
		// 关闭所有通道
		for _, ch := range chs {
			close(ch)
		}
		close(exitChan)
	}
}

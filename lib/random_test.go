package lib

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	var uin uint64 = 931055517
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	shuffleTask(arr1, uin)
	fmt.Printf("%+v\n", arr1)
}

func shuffleTask(tasks []int, uin uint64) {
	dateInt, _ := strconv.ParseUint(time.Now().Format("20060102"), 10, 64)

	seed := dateInt + uin
	rand.Seed(int64(seed))
	rand.Shuffle(
		len(tasks), func(i, j int) {
			tasks[i], tasks[j] = tasks[j], tasks[i]
		},
	)
}

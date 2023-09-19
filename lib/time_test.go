package lib

import (
	"fmt"
	"testing"
	"time"
)

func TestStringTrim(t *testing.T) {
	targetTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 21, 0, 0, 0, time.Local)
	currentTime := time.Now()

	if currentTime.After(targetTime) {
		fmt.Println("It's later than 9:00 PM now")
	} else {
		fmt.Println("It's earlier than 9:00 PM now")
	}

	time1 := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
	if time1.After(targetTime) {
		fmt.Println("It's later than 9:00 PM now")
	} else {
		fmt.Println("It's earlier than 9:00 PM now")
	}
}

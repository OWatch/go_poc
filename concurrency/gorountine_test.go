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

	go func() {
		for i := range chan1 {
			fmt.Println(i)
		}
	}()
	go func() {
		time.Sleep(time.Second * 1)
		chan1 <- 4
	}()
	go func() {
		time.Sleep(time.Second * 1)
		chan1 <- 5
	}()
	time.Sleep(time.Second * 5)

}

type Person struct {
	name string
}

func TestLoop(t *testing.T) {

	persons := []*Person{
		{
			name: "bob",
		},
		{
			name: "alice",
		},
		{
			name: "tom",
		},
	}

	for _, person := range persons {
		time.Sleep(time.Second * 1)
		go justPrint(person)
	}

}

func justPrint(p *Person) {
	fmt.Printf("pointer:%p", p)
	fmt.Printf("name:%s", p.name)
	fmt.Println("")
}

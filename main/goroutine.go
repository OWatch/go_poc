package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
}

func main() {

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

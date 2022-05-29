package main

import "fmt"

func main() {
	fmt.Println("hello world!")

	slice := make([]int, 6)
	slice = append(slice, 1)

	fmt.Println(slice)
}

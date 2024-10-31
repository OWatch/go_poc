package main

import (
	"fmt"
	"strings"
)

func main() {
	testArray()
}

func testArray() {
	// init
	arr1 := []string{"bob"}

	// read
	fmt.Println(arr1[0])

	// write
	// arr1 = append()
	// fmt.Println(arr1[1])
}

func testSlice() {
	s := make([]int, 0, 3)
	fmt.Println(s, len(s))
	s = append(s, 101)
	fmt.Println(s, len(s))
	fmt.Println("===========")

	s2 := make([]int, 3)
	fmt.Println(s, len(s2))
	s2 = append(s2, 101)
	s2[1] = 1
	fmt.Println(s2, len(s2))
}

func testMap() {
	// init1
	hash1 := map[string]string{
		"name": "bob",
	}
	// init2
	hash2 := make(map[string]string, 8)

	// write
	hash1["address"] = "beijing"
	hash2["name"] = "alice"

	// read
	fmt.Println(hash1["name"])
	fmt.Println(hash2["name"])

	// scan
	for k, v := range hash2 {
		fmt.Println(k, v)
	}

	// delete
	fmt.Println(hash1)
	delete(hash1, "name")
	fmt.Println(hash1)

	var dp [1][2]int
	fmt.Println(dp)
}

func testErr() {
	arr := strings.Split("", ",")

	for _, str := range arr {
		if str == "" {
			fmt.Println("kongzifuchguan" + str)
		}
	}

	fmt.Println("end")
}

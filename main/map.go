package main

import (
	"fmt"
	"testing"
)

func TestMapDeclaration(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic:%v", r)
			fmt.Println()
		}
	}()

	var map2 = make(map[string]string)
	fmt.Printf("len map2:%d", len(map2))
	fmt.Println()

	var map1 map[string]string
	map1["name"] = "alice"
	fmt.Printf("map1:%v", map1)
	fmt.Println()
}

func TestSliceDeclaration(t *testing.T) {
	var slice1 []string
	fmt.Printf("len slice1:%d", len(slice1))
	fmt.Println()

	slice1 = append(slice1, "alice")
	fmt.Printf("slice1:%v", slice1)
	fmt.Println()

}

func TestInterfaceDeclaration(t *testing.T) {
	type Person struct {
		Name string
	}

	var person1 Person
	fmt.Printf("person1:%v", person1)
	fmt.Println()
}

func TestMapValueNone(t *testing.T) {
	var map1 = make(map[string]string)
	map1["name"] = "alice"
	fmt.Printf("map1:%v", map1)
	fmt.Println()
	if map1["age"] == "" {
		fmt.Printf("map1-age:%s", map1["age"])
	}
	fmt.Println()
}

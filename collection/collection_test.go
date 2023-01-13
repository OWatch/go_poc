package collection

import (
	"fmt"
	"strings"
	"testing"
)

func TestArray(t *testing.T) {
	// init
	arr1 := []string{"bob"}

	// read
	fmt.Println(arr1[0])

	// write
	//arr1 = append()
	//fmt.Println(arr1[1])
}

func TestMap(t *testing.T) {
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

}

func TestErr(t *testing.T) {
	arr := strings.Split("", ",")

	for _, str := range arr {
		if str == "" {
			fmt.Println("kongzifuchguan" + str)
		}
	}

	fmt.Println("end")
}

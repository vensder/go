package main

import (
	"fmt"
)

func main() {
	counts_map := make(map[string]int)
	fmt.Printf("%v\n", counts_map)
	counts_map["key1"] = 1
	fmt.Printf("%v\n", counts_map)
	counts_map["key2"] = 3
	fmt.Printf("%v\n", counts_map)
	counts_map["key2"] = 2
	fmt.Printf("%v\n", counts_map)
	counts_map["key3"] = 3
	fmt.Printf("%v\n", counts_map)
	counts_map["key3"] = 300
	fmt.Printf("%v\n", counts_map)
	fmt.Printf("%v\n", counts_map["key3"])
	counts_map["key3"] *= counts_map["key3"]
	fmt.Printf("%v\n", counts_map["key3"])
	fmt.Printf("%v\n", counts_map)

	for key, value := range counts_map {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}
}

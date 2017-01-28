package main

import (
	"fmt"
)

func main() {

	arr := [10]int64{}
	i := 0

	for range arr {
		fmt.Printf("%v: %v\n", i, arr)
		i++

	}
}

package main

import (
	"fmt"
	"time"
)

func main() {

	arr := [10]int64{}
	i := 0

	fmt.Println("Welcome to the playground!",
		"\nThe time is %s", time.Now())
	fmt.Println("\"time.Now()\"")

	for range arr {
		fmt.Printf("%v: %v\n", i, arr)
		i++

	}
}

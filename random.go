package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX int = 10

var rand_arr [10]int

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i, _ := range rand_arr {
		rand_arr[i] = rand.Intn(MAX)
	}
	fmt.Println(rand_arr)
}

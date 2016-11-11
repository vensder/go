package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX int = 100

var rand_arr [10][10]int

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range rand_arr {
		for j := range rand_arr[i] {
			rand_arr[i][j] = rand.Intn(MAX)
		}
	}
	for i := range rand_arr {
		fmt.Println(rand_arr[i])
	}
}

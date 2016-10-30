package main

import (
	"math/rand"
)

func main() {
	for i := 1; i <= 100; i++ {
		rand.Seed(int64(i))
		println("Rand: ", rand.Intn((100)))
	}
}

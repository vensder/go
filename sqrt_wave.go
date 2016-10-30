package main

import (
	"math"
)

func sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	start := 0.000001111111111
	for i := 1; i <= 100; i++ {
		start = 1 / sqrt(start)
		println(start)
	}
}

package main

import (
	"math"
)

func main() {
	pi := math.Pi
	sqr := func(x float64) float64 {
		return math.Sqrt(x)
	}
	print(sqr(pi))
}

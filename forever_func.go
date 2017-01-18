package main

import (
	"fmt"
	"math"
	"strings"
)

var max_l int = 80

func line(n int) {
	symbol := "-"
	fmt.Println(strings.Repeat(symbol, n))
}

func sin_int(n int) int {
	return int(math.Sin(float64(n/max_l) * (math.Pi / 2.0)) * float64(max_l))
}

func main() {
	for {
		for i := 0; i < max_l; i++ {
			line(sin_int(i))
		}
		for i := max_l; i > 0; i-- {
			line(sin_int(i))
		}
	}
}

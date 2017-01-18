package main

import (
	"fmt"
	"math"
	"strings"
)

var max_l int = 160

func line(n int) {
	symbol := "#"
	fmt.Println(strings.Repeat(symbol, n))
}

// max_l * sin( (pi/2) * n/max_l)
func sin_int(n int) int {
	return int(math.Ceil(float64(max_l) * math.Sin((math.Pi/2.0) * float64(n)/float64(max_l) )))
}

func main() {
	for {
		for i := 1; i < max_l; i++ {
			line(sin_int(i))
		}
		for i := max_l; i > 0; i-- {
			line(sin_int(i))
		}
	}
}

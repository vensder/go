package main

import (
	"fmt"
	"math"
)

func main() {
	var delta float64 = 0.05
	for i := 0; i <= 10; i++ {
		fmt.Printf("%2.2f: %2.4f\n", float64(i)*delta, math.Sin(delta*float64(i)*math.Pi))
	}
}

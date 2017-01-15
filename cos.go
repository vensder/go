package main

import (
	"fmt"
	"math"
)

func main() {
	var delta float64 = 0.05
	for i := 0; i <= 10; i++ {
		fmt.Printf("%v: %v\n", float64(i)*delta, math.Cos(delta*float64(i)*math.Pi))
	}
}

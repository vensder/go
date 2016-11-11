package main

import (
	"fmt"
	"math"
)

var (
	max_repeats int     = 555
	x           float64 = 555
	prec        float64 = 0.00000000000000001
)

func mod(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Sqrt(x float64) float64 {
	z := 2.0
	for i := 1; i <= max_repeats; i++ {
		z1 := 0.5 * (z + x/z)
		delta := mod(z1 - z)
		if delta > prec {
			fmt.Printf("Iteration: %v. Delta: %v\n", i, z1-z)
			z = z1
		} else {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(x))
	fmt.Println(math.Sqrt(x))
}

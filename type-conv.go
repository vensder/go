package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var sqr32 float32
	var sqr64 float64
	var f float32 = float32(float64(float32(math.Sqrt(float64(x*x + y*y)))))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	sqr64 = (math.Sqrt(3))
	fmt.Printf("%v\n", sqr64)
	for i := 0; i < 10; i++ {
		sqr32 = float32(sqr64)
		fmt.Printf("%v\n", sqr32)
		sqr64 = float64(sqr32)
		fmt.Printf("%v\n", sqr64)
	}

}

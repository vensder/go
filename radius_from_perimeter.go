package main

import "math"

func R(P float64) float64 {
	return (P / (2 * math.Pi))
}

//Perimeter (P) = 2 · π · R
//R = P/(2π)

func main() {
	println(R(10))
}

package main

import "math"
import "fmt"

func Radius(Perimeter float64) float64 {
	return (Perimeter / (2 * math.Pi))
}

func Perimeter(Radius float64) float64 {
	return (2 * math.Pi * Radius)
}

//Perimeter (P) = 2 · π · R
//R = P/(2π)

func main() {

	fmt.Printf("If Perimeter = 10, than Padius = %v\n", Radius(10))
	fmt.Printf("If Radius = 1, than Perimeter = %v\n", Perimeter(1))
}

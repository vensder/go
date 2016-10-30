package main

import (
	"math"
)

func main() {
	p := func(s string) { print("\n", s, "\n") }
	f := func(f float64) { print("\n", f, "\n") }
	p("Now you have")
	f(math.Sqrt(7))
	p("problems.")
	p("test")
}

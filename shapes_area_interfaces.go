package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type triangle struct {
	a float64
	b float64
	c float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t triangle) area() float64 {
	p := (t.a + t.b + t.c) / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

type shape interface {
	area() float64
}

func printArea(sh shape) {
	fmt.Printf("Area of %T is %.2f\n", sh, sh.area())
}

func main() {
	s1 := square{10}
	c1 := circle{5}
	t1 := triangle{10, 10, 10}

	printArea(s1)
	printArea(c1)
	printArea(t1)
}

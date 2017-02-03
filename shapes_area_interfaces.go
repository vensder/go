package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perim() float64
}

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

func (s square) perim() float64 {
	return s.side * 4
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t triangle) perim() float64 {
	return t.a + t.b + t.c
}

func (t triangle) area() float64 {
	p := t.perim() / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func printArea(sh shape) {
	fmt.Printf("Area of %T is %.2f, perimeter is %.2f\n", sh, sh.area(), sh.perim())
}

func main() {
	s1 := square{10}
	c1 := circle{5}
	t1 := triangle{10, 10, 10}

	fmt.Println(s1.area())
	fmt.Println(c1.area())
	fmt.Println(t1.area())

	printArea(s1)
	printArea(c1)
	printArea(t1)
}

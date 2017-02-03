package main

import "fmt"

type square struct {
    side float64
}

func (s square) area() float64 {
    s.side * s.side

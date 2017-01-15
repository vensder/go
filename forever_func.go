package main

import (
	"fmt"
	"strings"
    "math"
)

func line(n int) {
	symbol := "-"
	fmt.Println(strings.Repeat(symbol, n))
}

func cos_int(n int) int {
    math.Cos(float64(n))
}

func main() {
	for {
		for i := 0; i < 60; i++ {
			line(i)
		}
        for i := 60; i > 0; i--{
            line(i)
        }
	}
}

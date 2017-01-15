package main

import (
	"fmt"
	"strings"
)

func line(n int) {
	symbol := "-"
	fmt.Println(strings.Repeat(symbol, n))
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

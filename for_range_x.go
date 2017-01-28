package main

import (
	"fmt"
)

func main() {
	x := "hello"
	y := "12345"
    i := 0

    for range x {
		fmt.Printf("%v\n", x)
	}

    for range y {
		fmt.Printf("%c\n", y[i])
        i++
	}
}

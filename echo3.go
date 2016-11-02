package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}
	for a, b := range os.Args[1:] {
        fmt.Printf("%v: %v\n", a, b)
	}
	fmt.Printf("Total args: %v\n", len(os.Args[1:]))
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    fmt.Printf("Executable file name: %v\n", os.Args[0])
	if len(os.Args[1:]) > 0 {
        println("Args:")
		fmt.Println(strings.Join(os.Args[1:], " "))
		fmt.Printf("Total args: %v\n", len(os.Args[1:]))
        fmt.Printf("Array of args: %v\n",os.Args[1:])
	} else {
		println("NO args")
	}
}

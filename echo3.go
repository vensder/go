package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}
	for index, arg := range os.Args[1:] {
		fmt.Printf("%v: %v\n", index+1, arg)
	}
	fmt.Printf("Total args: %v\n", len(os.Args[1:]))
	fmt.Printf("Args: %v\n", string(os.Args[1:]))
}

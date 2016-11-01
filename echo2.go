package main

import (
	"fmt"
	"os"
)

func main() {
	sum_string, separator := "", ""
	for _, arg := range os.Args[1:] {
		sum_string += separator + arg
		separator = " "
	}
	fmt.Println(sum_string)
    fmt.Println(range os.Args[1:])
}

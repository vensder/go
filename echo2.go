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
    println("Executable path:", os.Args[0])
    println("Quantity of the arguments:", len(os.Args[1:]))
	fmt.Println("Arguments:",sum_string)
}

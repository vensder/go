package main

import (
	"fmt"
	"os"
)

func main() {
	var count int
	var sum_string, separator string = "", ""
	for i := 1; i < len(os.Args); i++ {
		sum_string += separator + os.Args[i]
		separator = " "
		count = i
	}
	println("Executable path:", os.Args[0])
	println("Quantity of the arguments:", count)
	fmt.Println(sum_string)
}

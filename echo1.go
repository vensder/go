package main

import (
	"fmt"
	"os"
)

func main() {
	var sum_string string = ""
	for i := 1; i < len(os.Args); i++ {
		sum_string += os.Args[i] + " "
		//sum_string += " "
	}
	println("Executable path:", os.Args[0])
	println("Quantity of the arguments:", len(os.Args[1:]))
	fmt.Println(sum_string)
}

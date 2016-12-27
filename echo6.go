// echo6.go buffer

package main

import (
	"bytes"
	"fmt"
	"os"
)

var buffer bytes.Buffer

func main() {
	separator := ""
	for _, arg := range os.Args[1:] {
		buffer.WriteString(separator + arg)
		separator = " "
	}
	println("Executable path:", os.Args[0])
	println("Quantity of the arguments:", len(os.Args[1:]))
	fmt.Println("Arguments:", buffer.String())
}

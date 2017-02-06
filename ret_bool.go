package main

import "fmt"

func isNull(a int) bool {
	return (a == 0)
}

func main() {
	b := isNull(0)
	fmt.Printf("%T: %v\n", b, b)
}

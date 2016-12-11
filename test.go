package main

import "fmt"

func divid(x, y float32) int {
	return (int(x) / int(y))
}

func main() {
	fmt.Printf("%v\n", 1*4/9)
	fmt.Printf("%v\n", 7*4/9)
	fmt.Printf("%v\n", 3/2)
	fmt.Printf("%v\n", 3/2.0)

	fmt.Printf("%v\n", divid(3, 2))

	fmt.Printf("%b %b %b \n", 0 << 2, 1 << 2, 2 << 2)
}

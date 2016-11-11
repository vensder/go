package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(max(300, 300))
}

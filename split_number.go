package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("Sum of %d is: ", i)
		fmt.Println(split(i))
	}
}

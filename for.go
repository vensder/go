package main

import "fmt"

func main() {
	sum, i := 0, 0
	for ; ; i++ {
		sum += i
		if sum > 1000 {
			break
		}
		fmt.Println(sum)
	}

	sum = 1
	for sum < 1000 { // like while
		sum += sum
		fmt.Println(sum)
	}
}

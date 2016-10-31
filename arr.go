package main

import "fmt"

func main() {

	const range_arr = 10

	arr := [range_arr]int{}

	for i := range arr {
		arr[i] = range_arr - i - 1
	}

	for i := range arr {
		println(arr[i])
	}

	fmt.Println(arr)

	i := 0
	for range arr {
		println("hello, mr.robot!", i)
		i++
	}
	println("Array len:", len(arr))
}

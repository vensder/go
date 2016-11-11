package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	const i = 64
	const j = 1<<31 - 1
	const f = 5.55
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Printf("const i type is %T: %v\n", i, i)
	fmt.Printf("const j type is %T: %v\n", j, j)
	fmt.Printf("const f type is %T: %v\n", f, f)
}

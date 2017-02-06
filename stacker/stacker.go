package main

import (
	"fmt"
	"math"

	"./stack"
)

func main() {
	fmt.Println("hello from main")
	stack.Hello()

	var mystack stack.Stack
	mystack.Push("Hello")
	mystack.Push("banana")
	mystack.Push(-333)
	mystack.Push([]string{"one", "two", "three"})
	mystack.Push(math.Pi)
	for {
		item, err := mystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)

	}

}

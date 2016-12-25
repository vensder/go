package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	a, b, c = 1, 2.3, "f"
)

func main() {
	println("hi")
	t := time.Now().UTC().UnixNano()
	fmt.Printf("%v\n", t)
	s := strconv.FormatInt(t, 10)
	fmt.Printf("%v, %v\n", s, len(s))
	rand.Seed(t)

	fmt.Println("My favorite number is", rand.Intn(10))
	println("\na:", a, "\nb:", b, "\nc:", c)
	print("%v", c)
}

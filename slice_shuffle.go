package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	slice := []string{"one", "two", "three", "four", "five"}
	fmt.Printf("%v\n", slice)
	rand.Seed(time.Now().UnixNano())
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	fmt.Printf("%v\n", slice)
	output_string := strings.Join(slice, "\n")
	fmt.Printf("%v\n", output_string)
}

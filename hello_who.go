package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "Santa Klaus"
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Printf("Hello, %v\n", who)
}

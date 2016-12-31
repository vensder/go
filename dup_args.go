package main

import (
	//"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	//input := bufio.NewScanner(os.Stdin)
	for i := 1; i < len(os.Args); i++ {
		counts[os.Args[i]]++
	}
	// Potential errors from input.Err() ignored
	for arg, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t\"%s\"\n", n, arg)
		}
	}
}

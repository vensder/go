package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(`
        Hello! Input any multiple strings
        and symbols whith Enter at the end of each,
        with random repeats.
        When you'll finish, enter EOF (Ctrl+D):`)
	counts_map := make(map[string]int)
	fmt.Printf("%v\n", counts_map)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// line := input.Text()
		// counts_map[line] = counts_map[line] + 1
		counts_map[input.Text()]++
		fmt.Printf("input.Text: %v\n", input.Text())
		fmt.Printf("%v\n", counts_map)
	}
	// Potential errors from input.Err() ignored
	for line, n := range counts_map {
		if n > 1 {
			fmt.Printf("%d\t\"%s\"\n", n, line)
		}
	}
	fmt.Printf("%v\n", counts_map)
	fmt.Printf("%v\n", counts_map)
}

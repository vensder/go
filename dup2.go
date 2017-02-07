package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) {
	fmt.Printf("file: %v\n", f.Name())
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
    fmt.Printf("counts: %v\n", counts)
}

func main() {
	fmt.Println(`
        Hello! Input any multiple strings
        and symbols whith Enter at the end of each,
        with random repeats.
        When you'll finish, enter EOF (Ctrl+D).
        Also you can pass files names as args
        for input text from files(needed restart prog with args)`)
	counts_map := make(map[string]int)
//	files_map := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts_map)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts_map)
			f.Close()
		}
	}

	// Potential errors from input.Err() ignored
	for line, n := range counts_map {
		if n > 1 {
			fmt.Printf("%d\t\"%s\"\n", n, line)
		}
	}
	fmt.Printf("%v\n", counts_map)
}

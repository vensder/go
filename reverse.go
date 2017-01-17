package main

import (
	"fmt"
)

func reverse(s string) string {
	r := []rune(s)
	last_i := len(r) - 1

	for i := 0; ; i++ {
		r[i], r[last_i-i] = r[last_i-i], r[i]
		if i >= (last_i - i) {
			break
		}
	}

	return string(r)
}

func main() {
	fmt.Printf("%v\n", reverse("Chaotic"))
}

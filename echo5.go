// for testing:
//time go run echo5.go a b c d e

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const repeats = 10000000
const test = 3 // 1 or 2 or 3

var sum_string, separator string

func main() {
	t0 := time.Now()
	if test == 1 {
		for k := 0; k < repeats; k++ {
			sum_string = ""
			for i := 1; i < len(os.Args); i++ {
				sum_string += separator + os.Args[i]
				separator = " "
			}
		}
	} else if test == 2 {
		for k := 0; k < repeats; k++ {
			sum_string = ""
			for _, arg := range os.Args[1:] {
				sum_string += separator + arg
				separator = " "
			}
		}
	} else if test == 3 {
		for k := 0; k < repeats; k++ {
			sum_string = strings.Join(os.Args[1:], " ")
		}
	}
	fmt.Println(sum_string)
	t1 := time.Now()
	fmt.Printf("The function took %v to run in test #%v.\n", t1.Sub(t0), test)
}

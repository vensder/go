// for testing:
//time go run echo5.go a b c d e

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const repeats = 100000000

var sum_string, separator string
var t0, t1 time.Time

func main() {
	for test := 1; test <= 3; test++ {

		if test == 1 {
			t0 = time.Now()
			for k := 0; k < repeats; k++ {
				sum_string = ""
				for i := 1; i < len(os.Args); i++ {
					sum_string += separator + os.Args[i]
					separator = " "
				}
			}
			t1 = time.Now()
			fmt.Printf("The function took %v to run in test #%v.\n", t1.Sub(t0), test)

		} else if test == 2 {
			t0 = time.Now()
			for k := 0; k < repeats; k++ {
				sum_string = ""
				for _, arg := range os.Args[1:] {
					sum_string += separator + arg
					separator = " "
				}
			}
			t1 = time.Now()
			fmt.Printf("The function took %v to run in test #%v.\n", t1.Sub(t0), test)

		} else if test == 3 {
			t0 = time.Now()
			for k := 0; k < repeats; k++ {
				sum_string = strings.Join(os.Args[1:], " ")
			}
			t1 = time.Now()
			fmt.Printf("The function took %v to run in test #%v.\n", t1.Sub(t0), test)
		}

		fmt.Println(sum_string)

	}
}

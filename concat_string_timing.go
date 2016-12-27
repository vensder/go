// for testing:
//time go run echo5.go a b c d e

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"
)

const repeats = 100000

var sum_string, separator string
var t0, t1 time.Time
var buffer bytes.Buffer
var k, i int

func main() {
	for test := 1; test <= 4; test++ {

		if test == 1 {
			t0 = time.Now()
			for k = 0; k < repeats; k++ {
				for i = 1; i < len(os.Args); i++ {
					sum_string += os.Args[i] + " "
					//separator = " "
				}
			}
			t1 = time.Now()
			fmt.Printf("The function took %v to run in test #%v.\n", t1.Sub(t0), test)

		} else if test == 2 {
			t0 = time.Now()
			for k = 0; k < repeats; k++ {
				for _, arg := range os.Args[1:] {
					sum_string += arg + " "
					//separator = " "
				}
			}
			t1 = time.Now()
			fmt.Printf("The function took %v to run in test #%v.\n", t1.Sub(t0), test)

		} else if test == 3 {
			t0 = time.Now()
			for k = 0; k < repeats; k++ {
				sum_string = strings.Join(os.Args[1:], " ")
			}
			t1 = time.Now()
			fmt.Printf("The function took %v to run in test #%v.\n", t1.Sub(t0), test)

		} else if test == 4 {
			t0 = time.Now()
			for k = 0; k < repeats; k++ {
				//buffer.Reset()
				for _, arg := range os.Args[1:] {
					buffer.WriteString(arg)
					buffer.WriteString(" ")
					//separator = " "
				}
			}
			t1 = time.Now()
			fmt.Printf("The function took %v to run in test #%v.\n", t1.Sub(t0), test)
		}

	}
}

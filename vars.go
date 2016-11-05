package main

import "fmt"

var i, j, s = 1, 2.1, "hello!"

func main() {
	var c, python, java = true, !false, "no!"
	fmt.Println(i, j, c, python, java, s)
}

package main

var i int64
func pow(x, y int64) int64 {
	for i = 1; i < y; i++ {
		x += x
	}
	return x
}

func main() {
	for i = 1; i <= 64; i++ {
		println(2, "^", i, ":", pow(2, i))
	}
}

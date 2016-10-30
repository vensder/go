package main

func pow(x, y int) int {
	for i := 1; i < y; i++ {
		x += x
	}
	return x
}

func main() {
	for i := 1; i <= 20; i++ {
		println(2, "^", i, ":", pow(2, i))
	}
}

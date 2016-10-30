package main

func swap(x, y string) (_, _ string) {
	return y, x
}

func main() {
	a := "a"
	b := "b"
	a, b = swap(a, b)
	println(a, b)
}

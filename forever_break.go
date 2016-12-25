package main

func inf(count int) {
	i := 0
	for {
		i++
		if i >= count {
			break
		}
	}

}

func main() {
	println("Hello")

	inf(3000)
}

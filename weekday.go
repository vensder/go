package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()
	fmt.Printf("Todays weekday is: %v, type: %T\n", int(today), today)
	fmt.Printf("Saturdays day is: %v, type: %T\n", int(time.Saturday), time.Saturday)
	fmt.Println("When's Saturday?")

	delta := time.Saturday - today
	if delta == -1 {
		println("Yesterday")
	}
	if delta == 0 {
		println("Today")
	} else if delta == 1 {
		println("Tomorrow")
	} else if delta == 2 {
		println("In two days")
	} else {
		println("Too far away")
	}

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case 1 + today:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {
	// method 1:
	nowTime := time.Now()
	switch nowTime.Weekday() {
	case time.Saturday:
		fmt.Println("take a rest")
	case time.Sunday:
		fmt.Println("On vacation")
	default:
		fmt.Println("You need to work")
	}
	// method 2:
	switch {
	case nowTime.Weekday() >= time.Monday &&nowTime.Weekday()<=time.Friday:
		fmt.Println("Method 2 : You need to work")
	default:
		fmt.Println("Enjoy your vacation")
	}
}

package main

import "fmt"

func main() {
	// defer FILO, exec before return
	a, b := 1, 2
	defer fmt.Println("First defer output:", add(a, b))
	a = 3
	b = 4
	defer fmt.Println("Second defer output:", add(a, b))
	a = 5
	b = 6
	fmt.Println("Normal defer output:", add(a, b))
}
func add(a, b int) int {
	return a + b
}

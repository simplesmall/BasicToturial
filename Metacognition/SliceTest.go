package main

import "fmt"

func main() {
	// 数组 var name [size]T
	var numlist1 [3]int
	numlist1[0] = 11
	numlist1[1] = 22
	numlist1[2] = 33

	numList2 := [3]int{0, 1, 2}
	fmt.Println(numlist1, numList2)
	SliceTest()
}

// slice不包含end   slcie := source[begin:end]
// slice := make([],size,cap)
func SliceTest() {
	arr1 := [...]int{0, 1, 2, 3}
	slice1 := arr1[0:4]
	slice2 := make([]int, 4, 4)
	for i := 0; i < 4; i++ {
		slice2[i] = i
	}
	slice3 := append(slice1, 5)
	slice4 := []int{0, 1, 2, 3}
	fmt.Println(slice1, slice2, slice3, slice4)
}

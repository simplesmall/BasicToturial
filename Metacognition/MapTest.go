package main

import "fmt"

func main() {
	// var name map[keyType] valueType
	// name = make(map[keyType] valueType)
	map1 := map[int]string{}
	map1[1] = "01"
	map1[2] = "02"
	map1[3] = "03"

	map2 := map[int]string{
		1: "011",
		2: "022",
		3: "033",
	}
	fmt.Println(map1[0])

	// value,ok := map[key] 如果键存在于map中 ok将返回true 否则返回对应位置数据类型的初值 比如int 将返回0
	value, ok := map1[1]
	if ok {
		fmt.Println("1 is", value)
	} else {
		fmt.Println("1 is not existed !")
	}
	fmt.Println(map2)
}

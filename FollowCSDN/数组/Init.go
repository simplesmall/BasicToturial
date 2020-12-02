package main

import "fmt"

func main() {
	var str = []float32{12.3,22.0}
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	// 不指定元素个数时为切片
	var str1 = [...]string{"Hello","W","Word"}
	fmt.Println("数组测试",str1[0])

	// 更改元素值
	str[1] = 900
	fmt.Println(len(str))
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	var a = [3][4]int {
		{0, 1, 2, 3} ,    /* 第一行索引为 0 */
		{4, 5, 6, 7} ,    /* 第二行索引为 1 */
		{8, 9, 10, 11},   /* 第三行索引为 2 */
	}
	fmt.Println(a[1][2])

	// 向函数传递数组
}

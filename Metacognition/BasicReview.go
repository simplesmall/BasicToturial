package main

import "fmt"

func main() {
	fmt.Println("Testing main pak and main method")
	PointTest()
}

// 方法名首字母大写包外可访问
func Dev(source, target int) (handleSource, handleTarget int) {
	handleSource = source + 2
	handleTarget = 18989
	return
}

// 变量声明+初始化
var tempType int = 100

// 可省略类型自动推导
var temp = 100

// 短变量声明赋值,无法在函数外使用   name:= expression

func TestShortVar() {
	// 匿名函数接收error,后续不会用到
	value, _ := fmt.Println()
	// 声明的变量必须都使用到,否则编译报错
	fmt.Println(value)
}

//指针声明 var name *T
func PointTest() {
	name := "aoho"
	// get pointer value
	p := &name
	fmt.Println("Name is : ", *p)
	// set pointer value
	*p = "Jack zhou"
	fmt.Println("Name is : ", *p)
}


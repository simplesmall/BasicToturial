package main

import "fmt"

func main() {
	s0:=Student{}
	s1:=Student{
		StudentID: 0,
		Name: "First",
		birth: "2001-12-23",
	}
	s2:=Student{
		1,
		"Second",
		"1998-10-9",
	}
	s3:=&Student{
		StudentID: 3,
		Name: "Third",
		birth: "1997-09-08",
	}
	fmt.Println(s0,s1,s2,s3)
	fmt.Println(s0.Name,s1.Name,s2.Name,s3.Name)
}
/*
 变量名 + 类型
type structName struct{
	val1 T1
	val2 T2
}
*/

// 希望包外可以访问则首字母大写 + 包内不能重名
type Student struct {
	StudentID int64
	Name      string
	// 首字母小写仅包内可访问
	birth string
}
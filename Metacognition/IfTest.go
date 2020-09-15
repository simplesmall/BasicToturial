package main

import "fmt"

func main() {
	choosenMenu := 3
	ifTest(choosenMenu)
}

/*
	if con1{
		bran1
	}else if con2{
		bran2
	}else{
		bran3
	}
*/
func ifTest(menu int)  {
	if menu ==1 {
		fmt.Println("选中了111")
	}else if menu==2 {
		fmt.Println("选中了222")
	}else {
		fmt.Println("啥也没选中")
	}
}
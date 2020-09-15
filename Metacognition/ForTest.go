package main

import "fmt"

func main() {
	target:=4
	fmt.Println("Sum after add 0 ~",target," is ",sum(target))
	whileTest(target)
	forever(target)
}
/*
GO 中不存在while语句,直接用for语句缺省init 和 end
	for init;condition;end{
		circle body
	}
*/
func sum(times int)(sum int) {
	for i := 0; i <= times; i++ {
		sum += i
	}
	return sum
}

func whileTest(endPoint int)  {
	for endPoint<=10{
		fmt.Println("The value of ",endPoint,"  is  ",endPoint)
		endPoint++
	}
}
func forever(begin int)  {
	for{
		if begin <100{
			begin*=2
			fmt.Println("Forever item ",begin," is ",begin)
		}else {
			break
		}
	}
}


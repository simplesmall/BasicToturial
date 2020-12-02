package main

import "fmt"

func Show(strs ...interface{}){
	for _,str := range strs{
		switch str.(type) {
		case int:
			//for i := 0; i < len(str); i++ {
			//
			//	fmt.Println(str[i])
			//}
		}
	}
}
func StrSender(arr[] int)  {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

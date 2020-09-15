package main

func main() {

}

/*
新建channel
var name chan T // 双向channel
var name chan <-T // 只能发送消息的 channel
var name T <- chan // 只能接收消息的 channel
*/

/*
	channel <- val      //发送消息
	val:=  <- channel  // 接受消息
	val,ok := <- channel // 非阻塞接受消息
*/

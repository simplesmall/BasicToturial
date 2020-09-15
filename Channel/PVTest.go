package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan int)
	defer close(queue)
	for i := 0; i < 3; i++ {
		// 多个生产者
		go Producer(i*5, (i+1)*5, queue)
	}
	// 单个消费者
	go Consumer(queue)
	// 避免主 goroutine 结束程序
	time.Sleep(time.Second)
}

// 生产者
func Producer(begin, end int, queue chan<- int) {
	for i := begin; i < end; i++ {
		fmt.Println("Prodcuer : ", i)
		queue <- i
	}
}
func Consumer(queue <-chan int) {
	// 当前消费者循环消费
	for val := range queue {
		fmt.Println("consume : ", val)
	}
}

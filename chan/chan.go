package main

import "fmt"

func main() {
	// 管道
	chan1 := make(chan int, 2)
	fmt.Println("chan1 length: ", len(chan1), "chan1 cap: ", cap(chan1))
	chan1 <- 2
	close(chan1)
	fmt.Println("chan1 length: ", len(chan1), "chan1 cap: ", cap(chan1))

	// 双向管道
	chan2 := make(chan int, 1)
	chan2 <- 3
	close(chan2)
	fmt.Println("chan2 data: ", <-chan2)

	// 单向管道，只读管道
	chan3 := make(chan int, 1)
	type chanReader <-chan int
	var reader chanReader = chan3
	chan3 <- 3 // chan3用于写入
	close(chan3)
	fmt.Println("reader: ", <-reader) // 从reader只读管道中读出

	// 单向管道，只写管道
	chan4 := make(chan int, 1)
	type chanWriter chan<- int
	var writer chanWriter = chan4
	writer <- 4 // 写入到writer只写管道
	close(writer)
	fmt.Println("writer: ", <-chan4) // 从chan4中读出

	// 遍历管道数据
	chan5 := make(chan int, 5)
	chan5 <- 1
	chan5 <- 2
	chan5 <- 3
	chan5 <- 4
	chan5 <- 5
	close(chan5) // 需要关闭管道，否则遍历会阻塞
	for data := range chan5 {
		fmt.Println(data)
	}
}

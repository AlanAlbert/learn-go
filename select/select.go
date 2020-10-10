package main

import "fmt"

func main() {
	chan1 := make(chan string, 1)
	chan2 := make(chan string, 1)

	chan1 <- "Alan"
	chan2 <- "Jane"

	// 两个都有，随机一个
	select {
	case msg := <-chan1:
		fmt.Println("chan1", msg)
	case msg := <-chan2:
		fmt.Println("chan2", msg)
	default:
	}
}

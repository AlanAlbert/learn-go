package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("before:", runtime.NumGoroutine())
	num := 10
	ret := make(chan int, num)

	for i := 0; i < num; i++ {
		go func(i int) {
			time.Sleep(time.Millisecond * 10)
			ret <- i
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("after:", runtime.NumGoroutine())

	fmt.Println(<-ret)
}

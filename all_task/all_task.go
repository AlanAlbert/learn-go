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
	fmt.Println("after:", runtime.NumGoroutine())

	result := ""
	for i := 0; i < num; i++ {
		go func() {
			fmt.Println(<-ret)
		}()
	}
	fmt.Println(result)
	time.Sleep(time.Second)
}

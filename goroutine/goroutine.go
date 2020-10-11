package main

import (
	"fmt"
	"sync"
)

func goroutineTest() {
	fmt.Println("goroutine test")
}

func waitGroupTest(wg *sync.WaitGroup) {
	// wg.Add(1)
	defer wg.Done()
	fmt.Println("waitGroup test")
}

func main() {
	fmt.Println("main")
	go goroutineTest()

	// waitGroup
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go waitGroupTest(wg)
	wg.Wait()

	// 没有这行sleep，协程不会执行完成
	// time.Sleep(time.Second)
}

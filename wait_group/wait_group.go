package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker:", i)
}

func main() {
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(3)
	go worker(1, waitGroup)
	go worker(2, waitGroup)
	go worker(3, waitGroup)
	waitGroup.Wait()
	fmt.Println("main over")
}

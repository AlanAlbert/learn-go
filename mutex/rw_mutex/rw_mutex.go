package main

import (
	"fmt"
	"sync"
)

func incrTime(time *int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	defer wg.Done()

	// 写锁阻塞读写
	for i := 0; i < 100; i++ {
		mutex.Lock()
		*time++
		mutex.Unlock()
	}
}

func printTime(time *int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	defer wg.Done()

	for i := 0; i < 200; i++ {
		mutex.RLock()
		fmt.Println(*time)
		mutex.RUnlock()
	}
}

func main() {
	rwMutex := new(sync.RWMutex)
	time := new(int)
	wg := new(sync.WaitGroup)

	wg.Add(4)
	go incrTime(time, wg, rwMutex)
	go incrTime(time, wg, rwMutex)
	go printTime(time, wg, rwMutex)
	go printTime(time, wg, rwMutex)
	wg.Wait()
}

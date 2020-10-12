package main

import (
	"fmt"
	"sync"
)

// 普通自增
func incr(count *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		*count++
	}
}

// 加锁自增
func incrByMutex(count *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	for i := 0; i < 1000; i++ {
		*count++
	}
	mutex.Unlock()
}

func main() {
	count := new(int)
	wg := new(sync.WaitGroup)

	wg.Add(3)
	go incr(count, wg)
	go incr(count, wg)
	go incr(count, wg)
	wg.Wait()

	fmt.Println(*count) // 结果不一定是3000，因为*count++不是原子操作

	*count = 0
	mutex := new(sync.Mutex)
	wg.Add(3)
	go incrByMutex(count, wg, mutex)
	go incrByMutex(count, wg, mutex)
	go incrByMutex(count, wg, mutex)
	wg.Wait()

	fmt.Println(*count)
}

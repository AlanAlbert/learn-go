package main

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

type Singleton struct {
}

var singleton *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("create object")
		singleton = new(Singleton)
	})
	return singleton
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		fmt.Println(runtime.NumGoroutine())
		go func() {
			singleton := GetSingletonObj()
			fmt.Println(unsafe.Pointer(singleton))
			wg.Done()
		}()
	}
	wg.Wait()
}

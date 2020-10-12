package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)

	pool := New(4)
	work := func(args ...interface{}) {
		defer wg.Done()
		fmt.Println("No", args[0], "working...")
		time.Sleep(1 * time.Second)
		fmt.Println("No", args[0], "over.")
	}

	task1 := Task{
		function: work,
		args:     []interface{}{1},
	}

	task2 := Task{
		function: work,
		args:     []interface{}{2},
	}

	task3 := Task{
		function: work,
		args:     []interface{}{3},
	}

	wg.Add(3)
	pool.AddTask(task1)
	pool.AddTask(task2)
	pool.AddTask(task3)

	wg.Wait()
}

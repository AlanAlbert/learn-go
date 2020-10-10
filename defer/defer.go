package main

import "fmt"

func main() {
	// defer用于释放资源

	defer deferTest1()
	fmt.Println("after defer")
}

func deferTest1() {
	defer deferTest2()
	fmt.Println("defer test1")
}

func deferTest2() {
	fmt.Println("defer test2")
}

package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	// new 任意类型，初始化为零值，返回指针
	num := new(int)
	fmt.Println(*num)

	arr := new([3]int)
	fmt.Println(*arr)

	slice := new([]int)
	fmt.Println(*slice, len(*slice), cap(*slice))

	stu := new(student)
	fmt.Println(*stu)

	// make 用于slice、map、chan，返回引用
	slice2 := make([]int, 2, 4)
	fmt.Println(slice2)

	map2 := make(map[string]string)
	fmt.Println(map2)
	map2["name"] = "Alan"
	fmt.Println(map2)

	chan2 := make(chan string, 1)
	chan2 <- "Alan"
	fmt.Println(chan2, <-chan2)

}

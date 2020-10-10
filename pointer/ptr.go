package main

import "fmt"

func main() {
	// 获取指针1
	var1 := 18
	ptr1 := &var1
	*ptr1 = 19
	fmt.Println(var1, *ptr1)

	// 获取指针2
	ptr2 := new(int)
	fmt.Println(*ptr2)
	*ptr2 = 100
	fmt.Println(*ptr2)
}

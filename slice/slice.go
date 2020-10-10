package main

import "fmt"

func main() {
	// 1. array to slice
	arr1 := [...]int{1, 2, 3, 4}
	slice1 := arr1[2:4]
	fmt.Println(slice1, len(slice1), cap(slice1))

	// 2. []int{}
	slice2 := []int{1, 2}
	slice2 = append(slice2, 1, 2, 3)
	fmt.Println(slice2, len(slice2), cap(slice2))

	// 3. make
	slice3 := make([]int, 2, 2)
	fmt.Println(slice3, len(slice3), cap(slice3))

	// add data
	slice3 = append(slice3, 1, 2, 3, 4)
	fmt.Println(slice3, len(slice3), cap(slice3))

	slice3 = append(slice3, slice3...)
	fmt.Println(slice3, len(slice3), cap(slice3))

	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6:8]
	fmt.Printf("myslice为 %d, 其长度为: %d\n", myslice, cap(myslice))

	myslice = myslice[:4]
	fmt.Printf("myslice的第四个元素为: %d", myslice[3])
}

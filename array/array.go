package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1)

	arr2 := [...]int{1, 2, 3, 4}
	fmt.Println(arr2)

	arr3 := [4]int{2: 4}
	fmt.Println(arr3)

	arr4 := [...]int{2: 2}
	fmt.Println(arr4)
}

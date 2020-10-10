package main

import "fmt"

func main() {
	// for1
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// for2
	for i2 := 6; i2 <= 10; i2++ {
		fmt.Println(i2)
	}

	// for3
	arr := []int{1, 2, 3, 4, 5, 6}
	for item := range arr {
		fmt.Println(item)
	}
}

package main

import "fmt"

func main() {
	// recover必须放在defer中
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("catched!: ", err)
		}
	}()

	panic("thrown!")

}

package main

import "fmt"

func main() {
	var i interface{} = 100
	if num, ok := i.(int); ok {
		fmt.Println("断言成功: int", num)
	} else {
		fmt.Println("断言失败")
	}

	if str, ok := i.(string); ok {
		fmt.Println("断言成功: string", str)
	} else {
		fmt.Println("断言失败")
	}

	// type switch
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
}

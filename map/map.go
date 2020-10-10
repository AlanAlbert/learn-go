package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 声明1
	map1 := map[string]int{"Alan": 18, "Jane": 17}
	fmt.Println(map1)
	fmt.Println()

	// 声明2
	map2 := make(map[string]int)
	map2["Alan"] = 18
	map2["Jane"] = 17
	fmt.Println(map2)
	fmt.Println()

	// 删除某个元素
	delete(map2, "Alan")
	fmt.Println(map2)
	fmt.Println()

	// 获取不存在的key
	map3 := map[int]string{18: "Alan", 20: "Jane"}
	fmt.Println(map3, reflect.TypeOf(map3[19]))

	name, ok := map3[19]
	fmt.Println(ok)
	if ok {
		fmt.Println(name)
	} else {
		fmt.Println("map3[19] not exists")
	}
	fmt.Println()

	// 遍历
	for name, age := range map1 {
		fmt.Println(name, age)
	}
	fmt.Println()
	for name := range map1 {
		fmt.Println(name)
	}
	fmt.Println()
	for _, age := range map1 {
		fmt.Println(age)
	}
}

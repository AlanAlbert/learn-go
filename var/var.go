package main

import "fmt"

func main() {
	// var <name> <type>
	var age uint8 = 18
	var name = "Alan"
	var from = "China"
	fmt.Println(name, age, from)

	// 多变量声明
	// var ()
	var (
		varStr string
		varInt int
	)
	varStr = "string"
	varInt = 1
	fmt.Println(varStr, varInt)

	// <name> := <value>
	varStr2 := "string2"
	varInt2 := 2
	fmt.Println(varStr2, varInt2)

	// <name1>, <name2> := <value1>, <value2>
	varStr3, varInt3 := "string3", 3
	fmt.Println(varStr3, varInt3)

	// new ptr
	ptrStr := new(string)
	*ptrStr = "string4"
	fmt.Println(*ptrStr)
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// byte <=> uint8
	var varByte byte = 65
	var varInt8 uint8 = 66
	fmt.Printf("%c, %c, %s\n", varByte, varInt8, reflect.TypeOf(varByte))

	// rune <=> int32
	var varRune rune = '测'
	fmt.Println(string(varRune), reflect.TypeOf(varRune))
}

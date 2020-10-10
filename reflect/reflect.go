package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [...]int{1, 2, 3}
	map1 := map[string]int{
		"Alan": 19,
		"Jane": 18,
	}

	fmt.Println(reflect.TypeOf(arr), reflect.TypeOf(map1))

	typeOf, valueOf := reflect.TypeOf(arr), reflect.ValueOf(arr)
	fmt.Println(typeOf, valueOf)

	value := reflect.Value(valueOf)
	types := reflect.Type(typeOf)
	fmt.Println(value, types)

	v := valueOf.Interface()
	fmt.Println(v)
}

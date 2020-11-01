package main

import (
	"fmt"
	"reflect"
)

type testReflect struct {
	name string	`testTag:"this is tag" anotherTag:"this is another tag"`
	age int
}

func (t *testReflect) GetName() string {
	return t.name
}

func (t *testReflect) SetName(name string) {
	t.name = name
}

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

	intV := 89
	str := "csam,dnsa"
	testType(intV)
	testType(str)

	person := &testReflect{
		name: "Alan",
		age: 12,
	}

	method := reflect.ValueOf(person).MethodByName("SetName")
	method.Call([]reflect.Value{reflect.ValueOf("Jane")})
	fmt.Println(person.GetName())

	if field, ok := reflect.TypeOf(*person).FieldByName("name"); !ok {
		fmt.Println("can not get field(`name`)")
	} else {
		fmt.Println("test tag:", field.Tag.Get("testTag"), "\nanother tag:", field.Tag.Get("anotherTag"))
	}
}

func testType(v interface{}) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int,reflect.Int32,reflect.Int64:
		fmt.Printf("The type of %v is %T\n", v, v)
	case reflect.String:
		fmt.Printf("The type of %v is string\n", v)
	}
}

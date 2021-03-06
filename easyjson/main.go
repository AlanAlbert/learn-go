package main

import (
	"fmt"
	easyjson_example "go/easyjson_example/example"
)

func main() {
	stu := easyjson_example.Student{
		Person: easyjson_example.Person{
			Name: "Alan",
			Age: 18,
		},
		Class: "class 1",
		Score: 100,
	}

	teacher := easyjson_example.Teacher{
		Person: easyjson_example.Person{
			Name: "Jane",
			Age:  25,
		},
		Class: "class 1",
		Title: "No 1",
	}

	jsonStu, err := stu.MarshalJSON()
	if err != nil {
		fmt.Println("json encode failed!")
		return
	}
	fmt.Println("stu: ", string(jsonStu))

	jsonTeacher, err := teacher.MarshalJSON()
	if err != nil {
		fmt.Println("json encode failed!")
		return
	}
	fmt.Println("teacher: ", string(jsonTeacher))


	json := "{\"class\":\"class 3\",\"score\":90,\"name\":\"Anhoder\",\"age\":16}"
	stu2 := easyjson_example.Student{}
	if err := stu2.UnmarshalJSON([]byte(json)); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stu2)
}

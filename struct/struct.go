package main

import "fmt"

// Person person
type Person struct {
	name string
	age  uint8
}

// SetName set name
func (p *Person) SetName(name string) {
	p.name = name
}

// SetAge set age
func (p *Person) SetAge(age uint8) {
	p.age = age
}

// GetName get name
func (p Person) GetName() string {
	return p.name
}

// GetAge get age
func (p Person) GetAge() uint8 {
	return p.age
}

// Student student
type Student struct {
	Person
	className string
}

// SetClassName set class name
func (s *Student) SetClassName(className string) {
	s.className = className
}

// GetClassName get class name
func (s Student) GetClassName() string {
	return s.className
}

func main() {
	person := Person{name: "Alan", age: 18}
	fmt.Println(person)
	test(person)

	stu := Student{Person: Person{name: "Jane", age: 17}, className: "#1"}
	fmt.Println(stu)
	// test(stu) 错误
	fmt.Println(stu.GetName(), stu.GetAge(), stu.GetClassName())
}

func test(p Person) {
	fmt.Println(p.GetName(), p.GetAge())
}

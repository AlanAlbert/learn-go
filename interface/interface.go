package main

import "fmt"

// Person interface
type Person interface {
	say(msg string)
	getName() string
	getAge() uint8
}

// Student class
type Student struct {
	name string
	age  uint8
}

func (s Student) say(msg string) {
	fmt.Println("I am a Student,", msg)
}

func (s Student) getName() string {
	return s.name
}

func (s Student) getAge() uint8 {
	return s.age
}

func testPerson(p Person) {
	fmt.Println(p.getName(), p.getAge())
	p.say("Hello")
}

func main1() {
	stu := Student{"Alan", 18}
	testPerson(stu)
}

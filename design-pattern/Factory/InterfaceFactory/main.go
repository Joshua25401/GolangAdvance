package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

// Implement Say Hello
func (p *person) SayHello() {
	fmt.Println("Hello my name", p.name, "and my age is", p.age, "years old!")
}

func (tp *tiredPerson) SayHello() {
	fmt.Println("Hello i'm to tired to talk. Right Now!")
}

// Implement factory function
func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{
			name: name,
			age:  age,
		}
	}
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	person := NewPerson("Joshua", 20)
	person.SayHello()

	tiredPerson := NewPerson("Chesya", 103)
	tiredPerson.SayHello()
}

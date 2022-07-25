package main

import "fmt"

// Here in this section we learn about factory function
// First, we create a Person struct
// That has Name, Status, Age, and EyeCount
type Person struct {
	Name          string
	Status        string
	Age, EyeCount int
}

// Here we create a factory function
func NewPerson(name string, age int) *Person {
	// Here we can specify or validate some attribute
	// Let's say we validate if age > 17 ; Status become Adults
	// And below 17 ; Status become Teenager
	// Else ; Status become a Children
	status := "Children"
	if age >= 17 {
		status = "Adults"
	}
	if age < 17 {
		status = "Teenager"
	}
	return &Person{
		Name:     name,
		Age:      age,
		Status:   status,
		EyeCount: 2, // Here we specify the default value of EyeCount
	}
}

func main() {
	joshua := NewPerson("Joshua", 21)
	fmt.Println(joshua)
}

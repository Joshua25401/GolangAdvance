package main

import "fmt"

// Here the class that we want to build from scratch
type Person struct {
	name, position string
}

// Here we specify the function that can modify the class
type personModification func(person *Person)

// Here we made builder class in order to build a Person
type PersonBuilder struct {
	actions []personModification // Array of personModification function
}

// Here we specify function to build name
func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(person *Person) {
		person.name = name
	})
	return b
}

// So, we can extends the function and modify another attributes of person
// Just make a new function!
func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	b.actions = append(b.actions, func(person *Person) {
		person.position = position
	})
	return b
}

// Here we specify the Build() function that use to executed all the actions
// or all the modification
func (b *PersonBuilder) Build() *Person {
	person := Person{} // Declare a new person

	for _, action := range b.actions {
		action(&person) // Execute all the function modification
	}

	return &person // return the pointer to new person
}

func main() {
	personBuilder := PersonBuilder{}

	person := personBuilder.Called("Joshua").
		WorksAsA("Programmer in Virgo!").Build()

	fmt.Println("We got a new person here! Name", person.name,
		"And he'll work as a", person.position)
}

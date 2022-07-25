package main

import "fmt"

// Here Class Person
// We want to build this class using builder pattern
type Person struct {
	// Person must have address, postcode, city
	StreetAddress, Postcode, City string

	// Person must have a job
	CompanyName, Position string
	AnnualIncome          int
}

// Here our builder class
type PersonBuilder struct {
	person *Person
}

// This function is to start the builder to build
// Empty or nil Person{}
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}

// Connect the two additional builder with the function
// Syntax :
// func(builder *BuilderClass) Method_Name() *Connected_Class
func (builder *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*builder}
}

func (builder *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*builder}
}

func (builder *PersonBuilder) Build() *Person {
	return builder.person
}

// Add additional builder to build the Address
// and the Job. the additional function aggregates to PersonBuilder
type PersonAddressBuilder struct {
	PersonBuilder
}

// Make PersonAddressBuilder job method here
func (it *PersonAddressBuilder) At(
	streetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = streetAddress
	return it
}

func (it *PersonAddressBuilder) In(
	city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

func (it *PersonAddressBuilder) WithPostCode(
	postCode string) *PersonAddressBuilder {
	it.person.Postcode = postCode
	return it
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (it *PersonJobBuilder) At(
	companyName string) *PersonJobBuilder {
	it.person.CompanyName = companyName
	return it
}

func (it *PersonJobBuilder) AsA(
	position string) *PersonJobBuilder {
	it.person.Position = position
	return it
}

func (it *PersonJobBuilder) WithIncome(
	income int) *PersonJobBuilder {
	it.person.AnnualIncome = income
	return it
}

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("Jl.Pdt Justin Sihombing").
		In("Pematangsiantar").
		WithPostCode("21136").
		Works().
		At("Virgo").
		AsA("Backend Principal Engineer").
		WithIncome(7_000_000)
	person := pb.Build()
	fmt.Println(person)
}

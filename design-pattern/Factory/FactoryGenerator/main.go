package main

import "fmt"

type Employee struct {
	name         string
	annualIncome int
	position     string
}

// Functional Approach
func NewFunctionalFactory(position string,
	annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		// Here we fully initialize the Employee
		return &Employee{
			name:         name,
			annualIncome: annualIncome,
			position:     position,
		}
	}
}

// Structural Approach
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

// Make Create function to fully initialize the object
func (ef *EmployeeFactory) Create(name string) *Employee {
	return &Employee{
		name:         name,
		annualIncome: ef.AnnualIncome,
		position:     ef.Position,
	}
}

func NewStructuralFactory(position string,
	annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{
		Position:     position,
		AnnualIncome: annualIncome,
	}
}

func main() {
	// Implement the Functional Approach of Factory
	developerFactory := NewFunctionalFactory("Developer", 1_000_000)
	joshua := developerFactory("Joshua")
	fmt.Println("Here we've developer name ", joshua)

	// Implement the Structural Approach of Factory
	webDesignFactory := NewStructuralFactory("Web Designer", 1_000_000)
	chesya := webDesignFactory.Create("Chesya")
	fmt.Println("Here we've web designer name ", chesya)
}

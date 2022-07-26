package main

import "fmt"

type Employee interface {
	Introduce()
	Create(name string) *employee
}

type employee struct {
	Name, Position string
	AnnualIncome   int
}

func (e *employee) Introduce() {
	fmt.Println("Hello my name is", e)
}

func (e *employee) Create(name string) *employee {
	e.Name = name
	return e
}

const (
	BackEndEngineer = iota
	WebDesigner
)

func NewEmployeeFactory(role int) Employee {
	switch role {
	case BackEndEngineer:
		return &employee{
			Name:         "",
			Position:     "Back-End Engineer",
			AnnualIncome: 7_000_000,
		}

	case WebDesigner:
		return &employee{
			Name:         "",
			Position:     "Web Designer",
			AnnualIncome: 7_000_000,
		}
	}
	return nil
}

func main() {
	backendFactory := NewEmployeeFactory(BackEndEngineer)
	WebDesignerFactory := NewEmployeeFactory(WebDesigner)

	joshua := backendFactory.Create("Joshua")
	joshua.Introduce()

	chesya := WebDesignerFactory.Create("Chesya")
	chesya.Introduce()
}

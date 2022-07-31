package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office *Address
}

func (p *Employee) DeepCopy() *Employee {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	_ = encoder.Encode(&p)

	//fmt.Println(string(buffer.Bytes()))

	decoder := gob.NewDecoder(&buffer)
	result := Employee{}
	_ = decoder.Decode(&result)
	return &result
}

// Pre-defined Address to Main Office and other Aux Office
// this pre-defined called prototype
var mainOffice = &Employee{
	"",
	&Address{
		0,
		"Jl.Pdt Justin Sihombing",
		"Pematangsiantar",
	},
}

var auxOffice = &Employee{
	"",
	&Address{
		0,
		"Jl.Pdt Justin Sihombing",
		"Balige",
	},
}

// Prototype Factory
func newEmployee(prototype *Employee, name string, suite int) *Employee {
	result := prototype.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func main() {
	joshua := newEmployee(mainOffice, "Joshua", 40)
	chesya := newEmployee(auxOffice, "Chesya", 30)
	fmt.Println(joshua, joshua.Office)
	fmt.Println(chesya, chesya.Office)
}

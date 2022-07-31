package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Learning Deep Copy
type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		StreetAddress: a.StreetAddress,
		City:          a.City,
		Country:       a.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	_ = encoder.Encode(&p)

	//fmt.Println(string(buffer.Bytes()))

	decoder := gob.NewDecoder(&buffer)
	result := Person{}
	_ = decoder.Decode(&result)
	return &result
}

func main() {
	joshua := &Person{
		Address: &Address{
			"Jl.Pdt Justin Sihombing",
			"Pematangsiantar",
			"Indonesia",
		},
		Name:    "Joshua",
		Friends: []string{"Irwan", "Wilson"},
	}

	chesya := joshua.DeepCopy()
	chesya.Name = "Chesya"
	chesya.Friends = append(chesya.Friends, "Griselda")
	chesya.Address.City = "Balige"

	fmt.Println(joshua, joshua.Address)
	fmt.Println(chesya, chesya.Address)
}

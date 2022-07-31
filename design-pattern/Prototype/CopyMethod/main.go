package main

import "fmt"

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
	newPerson := *p
	newPerson.Address = p.Address.DeepCopy()
	copy(newPerson.Friends, p.Friends)
	return &newPerson
}

func main() {
	joshua := Person{
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

	fmt.Println(joshua, joshua.Address)
	fmt.Println(chesya, chesya.Address)
}

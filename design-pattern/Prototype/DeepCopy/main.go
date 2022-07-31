package main

import "fmt"

// Learning Deep Copy
type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	joshua := Person{
		Address: &Address{
			"Jl.Pdt Justin Sihombing",
			"Pematangsiantar",
			"Indonesia",
		},
		Name: "Joshua",
	}

	// It'll not deep copy the pointer address
	// Because if we want to copy pointer
	// We must do a Deep Copy Method that means
	// We Create all the new pointer
	chesya := joshua
	chesya.Name = "Chesya"
	//chesya.Address.City = "Balige"

	// Here we create all the address pointer and we change the city
	chesya.Address = &Address{
		StreetAddress: joshua.Address.StreetAddress,
		City:          joshua.Address.City,
		Country:       joshua.Address.Country,
	}
	chesya.Address.City = "Balige"

	fmt.Println(joshua, joshua.Address)
	fmt.Println(chesya, chesya.Address)
}

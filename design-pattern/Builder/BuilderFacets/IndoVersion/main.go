package main

import (
	"fmt"
)

type Person struct {
	Name string

	// Person must have address, postcode, city
	StreetAddress, Postcode, City string

	// Person must have a job
	CompanyName, Position string
}

/* PersonBuilder */
type PersonBuilder struct {
	person *Person
}

func (builder *PersonBuilder) YangPunya() *PersonBiographyBuilder {
	return &PersonBiographyBuilder{*builder}
}

func (builder *PersonBuilder) DanTinggalDi() *PersonAddressBuilder {
	return &PersonAddressBuilder{*builder}
}

func (builder *PersonBuilder) Ya() *Person {
	return builder.person
}

func TolongCreateOrang() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}

/* EOF */

type PersonBiographyBuilder struct {
	PersonBuilder
}

func (it *PersonBiographyBuilder) Nama(
	name string) *PersonBiographyBuilder {
	it.person.Name = name
	return it
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (it *PersonAddressBuilder) Jalan(
	streetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = streetAddress
	return it
}

func (it *PersonAddressBuilder) Kota(
	city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

func main() {
	josh := TolongCreateOrang().
		YangPunya().
		Nama("Joshua Pangaribuan").
		DanTinggalDi().
		Jalan("Jl.Pdt Justin Sihombing").
		Kota("Pematangsiantar").
		Ya()
	fmt.Printf("Hai udah di create nih!\n%v\n", josh)
}

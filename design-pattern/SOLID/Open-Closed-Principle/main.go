package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

// Filter
type Filter struct {
}

func (f *Filter) FilterByColor(
	products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(
	products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec *ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec *SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

type AndSpecification struct {
	first, second Specification
}

func (spec *AndSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) && spec.second.IsSatisfied(p)
}

type BetterFilter struct {
}

func (b *BetterFilter) Filter(
	products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{
		name:  "Apple",
		color: green,
		size:  small,
	}

	tree := Product{
		name:  "Tree",
		color: green,
		size:  large,
	}

	house := Product{
		name:  "House",
		color: blue,
		size:  large,
	}

	products := []Product{apple, tree, house}
	fmt.Println("Green Products (Old):")
	filter := Filter{}
	for _, v := range filter.FilterByColor(products, green) {
		fmt.Printf("- %s is green!\n", v.name)
	}

	fmt.Println("Large Products (Old):")
	for _, v := range filter.FilterBySize(products, large) {
		fmt.Printf("- %s is large!\n", v.name)
	}
	fmt.Println("Green Products (New):")
	greenSpec := ColorSpecification{color: green}
	betterFilter := new(BetterFilter)
	for _, v := range betterFilter.Filter(products, &greenSpec) {
		fmt.Printf("- %s is green!\n", v.name)
	}

	fmt.Println("Green and large Product (New):")
	largeSpec := SizeSpecification{large}
	for _, v := range betterFilter.Filter(products, &AndSpecification{
		first:  &greenSpec,
		second: &largeSpec,
	}) {
		fmt.Printf("- %s is green and large!\n", v.name)
	}
}

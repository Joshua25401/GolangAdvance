package main

import "fmt"

// Liskov Substitution Principle
// - If you have API with base clase it should work with derived class

// API (Application Programming Interface)
type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

// Base class
type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

// Derived class
type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := new(Square)
	sq.width = size
	sq.height = size
	return sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

// Solution
type Square2 struct {
	size int
}

func NewSquare2(size int) *Rectangle {
	rect := Rectangle{size, size}
	return &rect
}

// This method should can be used to both class
func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea, ", but got ", actualArea, "\n")
}

func main() {
	rect := &Rectangle{2, 3}
	UseIt(rect)

	square := NewSquare(5)
	UseIt(square)

	// So, this is the solution about violation to
	// Liskov Substitution Principle
	square2 := NewSquare2(5)
	UseIt(square2)
}

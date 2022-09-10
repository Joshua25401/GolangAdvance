package main

import "fmt"

/*
	Bridge Design-Pattern :
	Purpose :
		-> Prevent 'cartesian product' complexity explosion
		-> A Mechanism to decouples an interface (Hierarchy) from an implementation (Hierarchy)

	Use-Case Example :
		-> Common type ThreadScheduler
		-> Can be preemtive or cooperative
		-> Can run both on Windows and Unix
		-> This End up with 2x2 Scenario
			-> WindowsPTS
			-> WindowsCTS
			-> UnixPTS
			-> UnitCTS

	Solution : Use the bridge design pattern !
*/

// Renderer Example :
// We need to render a circle and a square
// Using vector and raster renderer
// We can do this with bridge design pattern by creating a interface called renderer
type Renderer interface {
	RenderCircle(radius float32)
	RenderSquare(sides int)
}

// Then we create the structure both for vector and raster renderer
// Both of them implements the method in Renderer interface
type VectorRenderer struct{}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle {Vector Renderer} with radius :", radius)
}
func (v *VectorRenderer) RenderSquare(sides int) {
	fmt.Println("Drawing a square {Vector Renderer} with sides :", sides)
}

type RasterRenderer struct{ Dpi int }

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle {RasterRenderer} with radius :", radius, "and DPI:", r.Dpi)
}
func (r *RasterRenderer) RenderSquare(sides int) {
	fmt.Println("Drawing a square {RasterRenderer} with sides :", sides, "and DPI:", r.Dpi)
}

// Then create the Circle and Square struct
type Circle struct {
	renderer Renderer // Depedency Injection
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

type Square struct {
	renderer Renderer // Depedency Injection
	sides    int
}

func NewSquare(renderer Renderer, sides int) *Square {
	return &Square{renderer: renderer, sides: sides}
}

func (sq *Square) Draw() {
	sq.renderer.RenderSquare(sq.sides)
}

func main() {
	vector := &VectorRenderer{}
	//raster := &RasterRenderer{Dpi: 100}

	circle := NewCircle(vector, 10)
	circle.Draw()

	square := NewSquare(vector, 10)
	square.Draw()
}

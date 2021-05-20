package interfaces

import "math"

// Rectangle has the dimensions of a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle represents a circle.
type Circle struct {
	Radius float64
}

// Triangle has the dimensions of a triangle.
type Triangle struct {
	Base   float64
	Height float64
}

// Shape is an interface that shows the Area.
type Shape interface {
	Area() float64
}

// Area returns the area of the Circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area returns the area of the Rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of a rectangle.
func Perimeter(r Rectangle) float64 {
	return (r.Width + r.Height) * 2
}

// Area returns the area of the Triangle.
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

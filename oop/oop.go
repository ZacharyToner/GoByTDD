//Package oop is for exploring structs, methods, and interfaces
package oop

import "math"

//Shape is our base interface for all shapes
type Shape interface {
	Area() float64
}

//Rectangle is our custom type to represent rectangles
type Rectangle struct {
	Width  float64
	Height float64
}

//Area returns the area of our Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Circle is our custom type to represent circles
type Circle struct {
	Radius float64
}

//Area returns the area of our circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Triangle is our custom type to represent triangles
type Triangle struct {
	Base   float64
	Height float64
}

//Area returns the area of our triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

//Perimeter returns the perimeter of a Rectangle
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

package main

import (
	"fmt"
	"math"
)

// Shape interfeysi
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle strukturasi
type Rectangle struct {
	width, height float64
}

// Circle strukturasi
type Circle struct {
	radius float64
}

// Rectangle uchun Area metodi
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Rectangle uchun Perimeter metodi
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle uchun Area metodi
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Circle uchun Perimeter metodi
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	var shapes []Shape
	shapes = append(shapes, Rectangle{width: 5, height: 3})
	shapes = append(shapes, Circle{radius: 4})

	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
		fmt.Printf("Perimeter: %.2f\n", shape.Perimeter())
	}
}

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

// Rectangle struckturasi
type Rectangle struct {
	width, height float64
}

// Circle struktursi
type Circle struct {
	radius float64
}

// Rectangle uchun Area metodi
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Rectangle uchun Perimeter metodi
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Circle uchun Perimetr metodi
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// PrintShapeDetails funckisyasi
func PrintfShapeDetails(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter:%.2f\n", s.Perimeter())
}

func main() {

	shapes := []Shape{
		Rectangle{width: 5, height: 3},
		Circle{radius: 5},
	}

	for _, shape := range shapes {
		PrintfShapeDetails(shape)
	}
}

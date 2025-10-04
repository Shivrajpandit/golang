package main

import (
	"fmt"
	"math"
)

// A `struct` is a collection of fields. It's useful for grouping
// data together to form records.
type Circle struct {
	x, y   float64
	radius float64
}

// A method is a function with a special receiver argument.
// Here, the `area` method has a receiver of type `Circle`.
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func structsExample() {
	// Create an instance of the Circle struct.
	c := Circle{x: 0, y: 0, radius: 5}

	// Access struct fields using dot notation.
	fmt.Printf("Circle radius: %.2f\n", c.radius)

	// Call the method associated with the struct.
	fmt.Printf("Circle area: %.2f\n", c.area())
}

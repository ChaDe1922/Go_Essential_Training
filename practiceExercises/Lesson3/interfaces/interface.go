package main

import (
	"fmt"
	"math"
)

//struct, Square, represents a square
//Square has a Length as a float64
type Square struct {
	Length float64
}

//struct, Circle, represents a Circle
//Circle has a Radius as a float 64
type Circle struct {
	Radius float64
}

//METHOD, Area, points to a Square object as reference
//returns the area as a float 64
func (s *Square) Area() float64 {
	return s.Length * s.Length
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//function, sumAreas, returns the sum of all areas in the slice
//pay attention to []Shape
func sumAreas(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

//Shape is a shape interface that utilizes METHODS that different structs have in common
type Shape interface {
	Area() float64
}

func main() {
	s := &Square{20}
	fmt.Println(s.Area())

	c := &Circle{10}
	fmt.Println(c.Area())

	shapes := []Shape{s, c}
	sa := sumAreas(shapes)
	fmt.Println(sa)

}

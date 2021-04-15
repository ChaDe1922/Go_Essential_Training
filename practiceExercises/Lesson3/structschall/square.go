//Struct Challenge: Square
//Define a Square struct, which has two fields: center of type point, and length of type int
//Add two methods
//Move(dx int, dy int) to move the square
//and, Area() int, calculates the area of the square
//also write NewSquare(x int, y int, length int) (*Square, error)
//x & y as the center of the square
//returns a pointer to a square, and a possible error
package main

import (
	"fmt"
	"log"
	"os"
)

//struct, Point, is a 2D point
type Point struct {
	X int
	Y int
}

//Define a Square struct.
//has two fields: center of type point, and length of type int
type Square struct {
	Center Point
	Length int
}

//write NewSquare(x int, y int, length int) (*Square, error)
func NewSquare(x int, y int, length int) (*Square, error) {
	//validate positive x
	if x < 0 {
		return nil, fmt.Errorf("x must be a positive number (was %d)", x)
	}

	//validate positive y
	if y < 0 {
		return nil, fmt.Errorf("y must be a positive number (was %d)", y)
	}

	//validate positive length
	if length <= 0 {
		return nil, fmt.Errorf("length must be a positive number (was %d)", length)
	}

	//create a variable, sq1, containing a pointer to a Square object
	//x & y as the center of the square
	sq1 := &Square{
		Center: Point{x, y},
		Length: length,
	}

	//returns a pointer to a square, and a possible error
	return sq1, nil
}

//METHOD 1, Move, moves the point, p
func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

//Method Moves the Square
func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}

// METHOD 2 Area() int, calculates the area of the square
func (sq *Square) Area() int {
	area := sq.Length * sq.Length
	return area
}

func main() {
	sq2, err := NewSquare(4, 7, 3)

	if err != nil {
		log.Fatalf("Error: can't create square - %s\n", err)
		os.Exit(1)
	}

	sq2.Move(4, 8)
	fmt.Printf("%+v\n", sq2)
	fmt.Println(sq2.Area())
}

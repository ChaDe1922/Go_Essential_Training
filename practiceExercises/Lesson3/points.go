//receiver example
package main

import (
	"fmt"
)

// struct, Point, is a 2d point
type Point struct {
	X int
	Y int
}

//METHOD, Move, moves the 2d point

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	p := Point{1, 2}
	p.Move(4, 9)
	fmt.Printf("%+v\n", p)
}

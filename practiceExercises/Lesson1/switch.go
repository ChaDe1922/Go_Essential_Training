package main

import (
	"fmt"
)

func main() {
	x := 2

	//This is an explicit switch statement
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("many")
	}

	//This is a naked switch statement where the cases will contain conditions
	switch {
	case x > 100:
		fmt.Println("x is very large")
	case x > 10:
		fmt.Println("x is pretty big")
	default:
		fmt.Println("x is smol")
	}
}

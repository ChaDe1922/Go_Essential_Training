//this program calculates the mean of two numbers
package main

import (
	"fmt"
)

func main() {

	//instantiates variables x and y as integers
	//var x float64
	//var y float64

	// assigns numbers 1 and 2 to variables x and y, respectively.
	//If you don't assign a number, Go will assign it as "0"
	// := allows Go to infer a data type for a variable
	//x := 1.0
	//y := 2.0

	//you can declare 2 variables on one line
	//if you create a variable and don't use it, Go will recognize this as an error
	x, y := 1.0, 2.0

	//The printf function receives a template to print, and then a value to fill this template
	//%v verb will print a Go object.
	//%t verb will pring its type.
	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	//This creates an integer variable called "mean"
	//var mean float64

	//sets mean equal to (x+y)/2
	//integer division returns an integer
	mean := (x + y) / 2.0
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}

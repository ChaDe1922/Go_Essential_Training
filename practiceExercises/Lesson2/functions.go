// basic function definitions in Go
package main

import (
	"fmt"
)

// func defines a function called add that adds a to b.
// Takes (integer arguements) and returns an integer
func add(a int, b int) int {

	//use a return statement to return the values
	return a + b
}

//function divmod takes integers a & b as arguements
// returns two integers, a quotient and remainder
//need parenthesis if you return more than 1 value
func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

//this is the function that is ran when the program runs
func main() {

	//variable val that stores the result of the function add
	val := add(1, 2)
	fmt.Println(val)

	//variables div and mod
	div, mod := divmod(7, 2)
	fmt.Printf("div=%d, mod=%d\n", div, mod)

}

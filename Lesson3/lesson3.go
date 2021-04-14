//conditionals. Examples of "if" statements
package main

import (
	"fmt"
)

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is bigger than 5")
	}

	//it looks like the else statement has to be jammed against the preceeding if statement
	if x > 100 {
		fmt.Println("x is larger than 100")
	} else {
		fmt.Println("x is less than 5. Issa tiny x!")
	}

	//&& gives us the logical "AND" statement in Go
	if x > 5 && x < 15 {
		fmt.Println("x is regular sized just like Rudy!")
	}

	//|| gives us the logical "OR" in Go
	if x < 20 || x > 30 {
		fmt.Println("x is out of range")
	}

	//setting assumed float variables a & b
	a := 11.0
	b := 20.0

	//If statements can have an optional initialization statement
	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}

}

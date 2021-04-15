//fmt.Sprint example

package main

import (
	"fmt"
)

func main() {

	//we set a variable n equal to 42
	n := 42

	//we create a variable, s, and use Sprintf to convert n to a string (%d) and save it to s
	s := fmt.Sprintf("%d", n)

	//Print the value and type for s
	fmt.Printf("s = %s (type %T)\n", s, s)

	//if you rplace %s with %q, it will surround s with quotation marks
	//fmt.Printf("s = %q (type %T)\n", s, s)
}

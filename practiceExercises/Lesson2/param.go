//this is an exercise for passing parameters
package main

import (
	"fmt"
)

//defines a function called doubleAt
//this function receives a slice of integers called values and a single integer i
//the function then takes the integer i to use as an index to double
func doubleAt(values []int, i int) {
	values[i] *= 2
}

//defines a function double that takes an integer n and doubles it
func double(n int) {
	n *= 2
}

//defines a function doublePrt takes a POINTER that points to an integer
func doublePtr(n *int) {

	//*n is dereferencing the pointer and assigning a value to it
	*n *= 2
}

//this is the main function
func main() {

	//initializes a slice "values" of integers 1, 2, 3, 4
	values := []int{1, 2, 3, 4}

	//calls the function doubleAt that takes "values" and the index 2
	//then prints the new values
	doubleAt(values, 2)
	fmt.Println(values)

	val := 10
	//in Go, changes to a variable inside a function do not affect the variable outside the function
	//this is not the case for maps and slices
	//to correct this, we must use pointers. You can pass a pointer to an object
	double(val)
	fmt.Println(val)

	//& passes a pointer to store in value variable
	doublePtr(&val)
	fmt.Println(val)
}

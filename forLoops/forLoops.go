//This is an example of a for loop
//The for loop is your only option for iteration in Go

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 9; i++ {
		fmt.Println(i)
	}

	fmt.Println("_________")
	//This for loops says that as long as i is less than 1, to continue on without executing the loop code below
	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("_________")

	//this version of the for loop is quite similar to the while loops
	a := 0
	for a < 3 {
		fmt.Println(a)
		a++ //This is a manual incrementor
	}

	fmt.Println("_________")

	b := 0
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}
}

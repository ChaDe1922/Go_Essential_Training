//go has trash management. We are exploring defer in this exercise
//When you allocate an object and stop using it, Go's garbage collector will clear it up
//make sure resources are closed when you are done with them
//for this, we use defer
package main

import (
	"fmt"
)

//this is a function that will free your resource.
func cleanup(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}

func worker() {

	//defer frees the resource "A" at the end of function worker
	//defers happen in reverse order
	defer cleanup("A")
	defer cleanup("B")

	fmt.Println("worker")
}

func main() {
	worker()
}

//A slice (aka: dictionary) is a sequence of items.
//All items in a slice must be of the same type

package main

import (
	"fmt"
)

func main() {
	//items of a slice must be of the same type
	loony := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loony = %v (type %T)\n", loony, loony)

	//to find the length of a slice/dictionary
	fmt.Println(len(loony)) //The length of slice "loony" should be 3

	//print divider
	fmt.Println("_______")

	//slices are 0 indexed
	fmt.Println(loony[1]) //This returns the 2nd entry, daffy

	//print divider
	fmt.Println("_______")

	//slices (or segments) of slices
	fmt.Println(loony[1:]) //slice from the 2nd entry to the end

	//print divider
	fmt.Println("_______")

	//iterating through a slicetionary using a for loop
	for i := 0; i < len(loony); i++ {
		fmt.Println(loony[i])
	}

	//print divider
	fmt.Println("_______")

	//single value range
	for i := range loony {
		fmt.Println(i)
	}

	//print divider
	fmt.Println("_______")

	//double value range
	for i, name := range loony {
		fmt.Printf("%s at %d\n", name, i)
	}

	//print divider
	fmt.Println("_______")

	//Double value range, ignore the index by using _. This returns JUST the values
	for _, name := range loony {
		fmt.Println(name)
	}

	//print divider
	fmt.Println("_______")

	//append information to the slicetionary
	loony = append(loony, "elmer")
	fmt.Println(loony)
}

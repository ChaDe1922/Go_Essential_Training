//Working with strings in Go
package main

import (
	"fmt"
)

func main() {
	book := "The colour of magic"
	fmt.Println(book)

	//This tells us how many bytes are in the variable book
	fmt.Println(len(book))

	//You can access individual bytes using [], indexing starts at 0
	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	//uint8 is a byte

	//Strings in Go are immutable and unchangeable
	//book[0] = 116 <----- This right here, it don't work!

	//slices take portions of a string
	//slice (start, end), 0 based, 1/2 empty range (aka exclusive)
	fmt.Println(book[4:11])

	//slice with no end
	fmt.Println(book[4:])

	//slice with no start
	fmt.Println(book[:4])

	//Use + for concatenation
	fmt.Println("t" + book[1:])

	//Unicode
	fmt.Println("It was 1/2 price!")

	//Multi-Line with the `` right under the ~

	poem := `
	The road goes ever on
	Down from the door where it began
	And all that other Jaaz
	with a flash in the pan
	...
	`

	fmt.Println(poem)
}

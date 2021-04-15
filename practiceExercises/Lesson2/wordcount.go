// count and print how many times each word appears in a text.
// To split the text words, use the strings.Fields
// use strings.ToLower to convert to lowercase

package main

import (
	"fmt"
	"strings"
)

func main() {

	// variable text with a series of text stored within
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`

	// variable words that separates each word and removes spaces from variable text
	words := strings.Fields(text)

	//counts defines an empty map from strings to integers, mapping keys to values
	counts := map[string]int{} //Empty map

	//for each word in the range of words, we go over each word, ignoring the indices with the _
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	fmt.Println(counts)
}

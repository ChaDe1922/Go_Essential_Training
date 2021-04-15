//This is learning with go
/*pring a friendly greeting*/

//go is organized in packages
package main

import (
	"fmt" //this packages contains functions for formatted print
)

//this is the main function. It will be executed by the Go runtime when the program starts
func main() {
	fmt.Println("Welcome to Go, Gophers?! :'D")
}

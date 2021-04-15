//This is an exploration of structs. (custon data types)

package main

import (
	"fmt"
)

//the struct called Trade represents a trade in stocks
type Trade struct {
	Symbol string  //Stock symbol as a string
	Volume int     //Number of shares as an integer
	Price  float64 //Trade price as a float
	Buy    bool    //True if we buy, false if we sell
}

func main() {

	//variable t1 is a Trade struct object
	//symbol: msft, volume: 10, price: 99.98, buy: true
	t1 := Trade{"MSFT", 10, 99.98, true}

	//print struct t1 data
	fmt.Println(t1)

	//print struct t1 object structure. %+v gives us the structure
	fmt.Printf("%+v\n", t1)

	//use the . to access an individual field
	fmt.Println(t1.Symbol)

	//we can also specify the field names
	t2 := Trade{
		Symbol: "MSFT", //Stock symbol as a string
		Volume: 10,     //Number of shares as an integer
		Price:  99.98,  //Trade price as a float
		Buy:    true,   //True if we buy, false if we sell
	}
	fmt.Printf("%+v\n", t2)

	//you can also omit all of the field names. Go will insert 0s
	t3 := Trade{}
	fmt.Printf("%+v\n", t3)

	//In go, everything that starts with an Uppercase letter is accessable to other packages (public)
	//Everything that starts with a lowercase letter, it is only accessible within the current package (private)
}

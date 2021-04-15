//This is an exercise in methods
//Structs are great at organizing data, but they have more power
//We can define methods on structs

package main

import (
	"fmt"
)

//struct, Trade, represents a trade in stocks
type Trade struct {
	Symbol string  //Stock symbol as a string
	Volume int     //Number of shares as an integer
	Price  float64 //Trade price as a float
	Buy    bool    //True if we buy, false if we sell
}

//the METHOD Value returns the trade value as a float64
//the METHOD has a receiver, (t *Trade), a pointer to a Trade struct object, t.
//you will usually use a pointer receiver so your method doesn't receive a COPY of the value
func (t *Trade) Value() float64 {

	//the value of the trade is set as the volume of t (converted to a float) * the price of t
	value := float64(t.Volume) * t.Price

	//if t.Buy is true
	if t.Buy {

		//the value of the trade is negative, as we've spent money
		value = -value
	}

	//return the value of the trade as a float64
	return value
}

func main() {

	//variable t contains a trade struct
	t := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}

	//print the value of trade, t.
	fmt.Println(t.Value())
}
